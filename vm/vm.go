package vm

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/IBAX-io/needle/compiler"

	log "github.com/sirupsen/logrus"
)

type GlobalVm struct {
	mu sync.Mutex
	vm *VM
}

var _vm *GlobalVm

func init() {
	_vm = &GlobalVm{
		vm: NewVM(),
	}
}

// VM is the main type of the virtual machine.
type VM struct {
	codeBlockMu sync.RWMutex
	CodeBlock   *compiler.CodeBlock
	// a function returns the cost of executing an external golang function.
	ExtCost func(fnName string) int64
	// if the function is in the list, the first of result must be a int64 that the cost of executing.
	FuncCallsDB map[string]struct{}
	// extern mode of compilation. an inter flag indicating whether a contract is an external contract.
	// It is set to true when a VM is created. Contracts called are not displayed
	// when the code is compiled. In other words, it allows to call the contract code
	// determined in the future;
	Extern compiler.IgnoreLevel
	// id of the first contract in the VM.
	ShiftContract int64
	logger        *log.Entry
}

// NewVM creates a new virtual machine.
func NewVM() *VM {
	auto := map[string]string{"*vm.Runtime": "rt"}
	fn := []compiler.ExtendFunc{
		{Name: "Settings", Func: GetSettings, AutoPars: auto},
		{Name: "MemoryUsage", Func: MemoryUsage, AutoPars: auto},
		{Name: "ExecContract", Func: ExecContract, CanWrite: false, AutoPars: auto},
		{Name: "CallContract", Func: CallContract, CanWrite: true, AutoPars: auto},
		{Name: "Println", Func: func(a ...any) (int64, error) {
			n, err := fmt.Println(a...)
			return int64(n), err
		}},
		{Name: "Sprintf", Func: fmt.Sprintf},
		{Name: "Bytes", Func: func(data any) []byte {
			return []byte(fmt.Sprintf("%v", data))
		}},
	}
	vm := &VM{
		CodeBlock: compiler.NewCodeBlock(&compiler.Config{
			Func:   fn,
			PreVar: GetSysVarsKeys(),
		}),
		Extern:      compiler.IgnoreIdent,
		FuncCallsDB: make(map[string]struct{}),
		logger:      log.WithFields(log.Fields{"type": VMErr, "extern": compiler.IgnoreIdent}),
	}
	return vm
}

// GetVM is returning smart vm.
func GetVM() *VM {
	_vm.mu.Lock()
	defer _vm.mu.Unlock()
	return _vm.vm
}

type SavePoint struct {
	objects       sync.Map
	childrenCount atomic.Uint32
}

// CreateSavePoint creates a snapshot of the VM's current state,
// storing all CodeBlock Objects and Children count.
// This allows restoring to this state later via RestoreSavePoint.
func (vm *VM) CreateSavePoint() *SavePoint {
	sp := &SavePoint{}
	vm.codeBlockMu.RLock()
	for k, v := range vm.CodeBlock.Objects {
		sp.objects.Store(k, v)
	}
	sp.childrenCount.Store(uint32(len(vm.CodeBlock.Children)))
	vm.codeBlockMu.RUnlock()
	return sp
}

// RestoreSavePoint restore the virtual machine to the state of the save point
func (vm *VM) RestoreSavePoint(sp *SavePoint) {
	vm.codeBlockMu.Lock()
	defer vm.codeBlockMu.Unlock()

	vm.CodeBlock.Objects = make(map[string]*compiler.Object)
	sp.objects.Range(func(k, v interface{}) bool {
		vm.CodeBlock.Objects[k.(string)] = v.(*compiler.Object)
		return true
	})

	vm.CodeBlock.Children = vm.CodeBlock.Children[:sp.childrenCount.Load()]
}

// Call executes the name object with the specified params and extended variables and functions.
func (vm *VM) Call(name string, extend map[string]any, costLimit int64) (ret []any, err error) {
	split := strings.Split(name, ".")
	obj := vm.CodeBlock.GetObjByName(split[0])
	if obj == nil {
		return nil, fmt.Errorf("object %s is empty", name)
	}
	if extend == nil {
		extend = make(map[string]any)
	}
	var block *compiler.CodeBlock
	rt := NewRuntime(vm, extend, costLimit)
	if obj.IsCodeBlockContract() {
		block = obj.GetCodeBlock().GetObjByName(split[1]).GetCodeBlock()
	} else if obj.IsCodeBlockFunction() {
		block = obj.GetCodeBlock()
	} else {
		return nil, fmt.Errorf("unknown object %s for call", name)
	}

	ret, err = rt.Run(block)
	extend[ExtendTxCost] = rt.CostRemain()
	fmt.Println("gas:", rt.CostUsed())
	return ret, err
}

// CompileEval compiles the source code and loads the byte-code into the virtual machine.
func CompileEval(vm *VM, src string, prefix uint32) error {
	var ok bool
	if len(src) == 0 {
		return nil
	}
	allowed := []string{
		`0`, `1`, `true`, `false`, `ContractConditions\(\s*\".*\"\s*\)`,
		`ContractAccess\(\s*\".*\"\s*\)`, `RoleAccess\(\s*.*\s*\)`,
	}
	for _, v := range allowed {
		re := regexp.MustCompile(`^` + v + `$`)
		if re.Match([]byte(src)) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf(eConditionNotAllowed, src)
	}
	err := vm.CompileEval(src, prefix)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`^@?[\w_]+$`)
	for _, item := range getContractList(src) {
		if len(item) == 0 || !re.Match([]byte(item)) {
			return errIncorrectParameter
		}
	}
	return nil
}

func getContractList(src string) (list []string) {
	for _, funcCond := range []string{`ContractConditions`, `ContractAccess`} {
		if strings.Contains(src, funcCond) {
			if ret := regexp.MustCompile(funcCond +
				`\(\s*(.*)\s*\)`).FindStringSubmatch(src); len(ret) == 2 {
				for _, item := range strings.Split(ret[1], `,`) {
					list = append(list, strings.Trim(item, "\"` "))
				}
			}
		}
	}
	return
}

// GetContractById returns the contract with the specified id.
func GetContractById(vm *VM, id int32) *compiler.ContractInfo {
	var tableId int64
	if id > ShiftContractId {
		tableId = int64(id - ShiftContractId)
		id = int32(tableId + vm.ShiftContract)
	}
	idcont := id
	if len(vm.CodeBlock.Children) <= int(idcont) {
		return nil
	}
	if vm.CodeBlock.Children[idcont] == nil || vm.CodeBlock.Children[idcont].Type != compiler.CodeBlockContract {
		return nil
	}
	if tableId > 0 && vm.CodeBlock.Children[idcont].GetContractInfo().Owner.TableId != tableId {
		return nil
	}
	return vm.CodeBlock.Children[idcont].GetContractInfo()
}

// RunContractById executes the contract with the specified id and methods.
func RunContractById(vm *VM, id int32, methods []string, extend map[string]any) error {
	info := GetContractById(vm, id)
	if info == nil {
		return fmt.Errorf(`unknown contract id '%d'`, id)
	}
	return RunContractByName(vm, info.Name, methods, extend)
}

// RunContractByName executes the contract with the specified name and methods.
func RunContractByName(vm *VM, name string, methods []string, extend map[string]any) error {
	obj, ok := vm.CodeBlock.Objects[name]
	if !ok {
		return fmt.Errorf(`unknown object '%s'`, name)
	}

	if !obj.IsCodeBlockContract() {
		return fmt.Errorf(eUnknownContract, name)
	}
	contract := obj.GetCodeBlock()
	extend[ExtendTxCost] = extend[ExtendTxCost].(int64) - CostContract - ContractBaseCost(contract)
	if extend[ExtendTxCost].(int64) < 0 {
		return fmt.Errorf("runtime cost limit overflow")
	}
	for i := 0; i < len(methods); i++ {
		method := methods[i]
		obj, ok := contract.Objects[method]
		if !ok {
			continue
		}
		if obj.IsCodeBlockFunction() {
			fn := obj.GetCodeBlock()
			_, err := Run(vm, fn, extend)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Run executes the code block with the specified extended variables.
func Run(vm *VM, block *compiler.CodeBlock, extend map[string]any) (ret []any, err error) {
	if block == nil {
		return nil, fmt.Errorf(`code block is nil`)
	}
	rt := NewRuntime(vm, extend, extend[ExtendTxCost].(int64))
	ret, err = rt.Run(block)
	extend[ExtendTxCost] = rt.CostRemain()
	if err != nil {
		// vm.logger.WithFields(log.Fields{
		// 	"type":              VMErr,
		// 	"error":             err,
		// 	"original_contract": extend[ExtendOriginalContract],
		// 	"this_contract":     extend[ExtendThisContract],
		// }).Error("running block in smart vm")
		return nil, err
	}
	return
}

// ObjectExists checks if the object with the specified name exists in the virtual machine.
func ObjectExists(vm *VM, name string, state uint32) bool {
	name = compiler.StateName(state, name)
	_, ok := vm.CodeBlock.Objects[name]
	return ok
}

// SetExtendCost sets the cost of calling extended obj in vm.
func (vm *VM) SetExtendCost(ext func(string) int64) *VM {
	vm.ExtCost = ext
	return vm
}

// SetFuncCallsDB Set up functions that can edit the database in vm.
func (vm *VM) SetFuncCallsDB(funcCallsDB map[string]struct{}) *VM {
	vm.FuncCallsDB = funcCallsDB
	return vm
}

// AppendPreVar appends the predeclared variables to the virtual machine.
func (vm *VM) AppendPreVar(preVar []string) *VM {
	vm.CodeBlock.PredeclaredVar = append(vm.CodeBlock.PredeclaredVar, preVar...)
	return vm
}

// FlushExtern switches off the extern mode of the compilation.
func (vm *VM) FlushExtern() {
	vm.Extern = compiler.IgnoreNone
	vm.logger = log.WithFields(log.Fields{"extern": compiler.IgnoreNone})
}

// MergeCompilerConfig merges the virtual machine configuration with the compiler configuration.
func (vm *VM) MergeCompilerConfig(conf *compiler.Config) *compiler.Config {
	conf.EnsureDefault()
	conf.SetObjects(vm.CodeBlock.Objects)
	v := append(vm.CodeBlock.PredeclaredVar, conf.PreVar...)
	slices.Sort(v)
	conf.PreVar = slices.Compact(v)
	return conf
}

// Compile compiles a source code and loads the byte-code into the virtual machine.
func (vm *VM) Compile(input []rune, conf *compiler.Config) error {
	root, err := compiler.CompileBlock(input, vm.MergeCompilerConfig(conf))
	if err != nil {
		return err
	}
	vm.FlushBlock(root)
	return nil
}

const flushMark = 1 << 20

// FlushBlock loads the compiled CodeBlock into the virtual machine.
func (vm *VM) FlushBlock(root *compiler.CodeBlock) {
	vm.codeBlockMu.Lock()
	defer vm.codeBlockMu.Unlock()

	shift := len(vm.CodeBlock.Children)
	for key, item := range root.Objects {
		cur, ok := vm.CodeBlock.Objects[key]
		if ok {
			if cur.IsCodeBlockContract() {
				item.GetContractInfo().Id = cur.GetContractInfo().Id + flushMark
			}
			if cur.IsCodeBlockFunction() {
				item.GetFunctionInfo().Id = cur.GetFunctionInfo().Id + flushMark
			}
		}
		vm.CodeBlock.Objects[key] = item
	}
	for _, item := range root.Children {
		if item == nil {
			continue
		}
		switch item.Type {
		case compiler.CodeBlockContract:
			if item.GetContractInfo().Id > flushMark {
				item.GetContractInfo().Id -= flushMark
				vm.CodeBlock.Children[item.GetContractInfo().Id] = item
				shift--
				continue
			}

			item.Parent = vm.CodeBlock
			item.GetContractInfo().Id += uint32(shift)
		case compiler.CodeBlockFunction:
			if item.GetFunctionInfo().Id > flushMark {
				item.GetFunctionInfo().Id -= flushMark
				vm.CodeBlock.Children[item.GetFunctionInfo().Id] = item
				shift--
				continue
			}
			item.Parent = vm.CodeBlock
			item.GetFunctionInfo().Id += uint32(shift)
		}

		vm.CodeBlock.Children = append(vm.CodeBlock.Children, item)
	}
}
