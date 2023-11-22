package vm

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/IBAX-io/needle/compile"

	log "github.com/sirupsen/logrus"
)

type GlobalVm struct {
	mu sync.Mutex
	vm *VM
}

var (
	_vm *GlobalVm
)

func init() {
	_vm = &GlobalVm{
		vm: NewVM(),
	}
}

// VM is the main type of the virtual machine
type VM struct {
	*compile.CodeBlock
	ExtCost       func(string) int64 // the cost of executing an extend function
	FuncCallsDB   map[string]struct{}
	Extern        compile.IgnoreLevel // extern mode of compilation
	ShiftContract int64               // id of the first contract
	logger        *log.Entry
}

// NewVM creates a new virtual machine
func NewVM() *VM {
	auto := map[string]string{"*vm.Runtime": "rt"}
	var fn = []compile.ExtendFunc{
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
			return []byte(fmt.Sprintf("%v", data.(interface{})))
		}},
	}
	var v []string
	for p := range sysVars {
		v = append(v, p)
	}
	vm := &VM{
		CodeBlock: compile.NewCodeBlock(&compile.CompConfig{
			Func:   fn,
			PreVar: v,
		}),
		Extern:      compile.IgnoreIdent,
		FuncCallsDB: make(map[string]struct{}),
		logger:      log.WithFields(log.Fields{"type": VMErr, "extern": true}),
	}
	return vm
}

// GetVM is returning smart vm
func GetVM() *VM {
	_vm.mu.Lock()
	defer _vm.mu.Unlock()
	return _vm.vm
}

var smartObjects map[string]*compile.Object
var children uint32

func SavepointSmartVMObjects() {
	smartObjects = make(map[string]*compile.Object)
	for k, v := range GetVM().Objects {
		smartObjects[k] = v
	}
	children = uint32(len(GetVM().Children))
}

// Call executes the name object with the specified params and extended variables and functions
func (vm *VM) Call(name string, extend map[string]any) (ret []any, err error) {
	split := strings.Split(name, ".")
	obj := vm.GetObjByName(split[0])
	if obj == nil {
		return nil, fmt.Errorf(`object %s is empty`, name)
	}
	if extend == nil {
		extend = make(map[string]any)
	}
	var block *compile.CodeBlock
	rt := NewRuntime(vm, extend, extend[ExtendTxCost].(int64))
	switch obj.Type {
	case compile.ObjContract:
		block = obj.GetCodeBlock().GetObjByName(split[1]).GetCodeBlock()
	case compile.ObjFunc:
		block = obj.GetCodeBlock()
	default:
		return nil, fmt.Errorf(`unknown object %s for call`, name)
	}
	ret, err = rt.Run(block)
	extend[ExtendTxCost] = rt.CostRemain()
	fmt.Println("gas:", rt.CostUsed())
	return ret, err
}

func RollbackSmartVMObjects() {
	GetVM().Objects = make(map[string]*compile.Object)
	for k, v := range smartObjects {
		GetVM().Objects[k] = v
	}

	GetVM().Children = GetVM().Children[:children]
	smartObjects = nil
}

func ReleaseSmartVMObjects() {
	smartObjects = nil
	children = 0
}

func CompileEval(vm *VM, src string, prefix uint32) error {
	var ok bool
	if len(src) == 0 {
		return nil
	}
	allowed := []string{`0`, `1`, `true`, `false`, `ContractConditions\(\s*\".*\"\s*\)`,
		`ContractAccess\(\s*\".*\"\s*\)`, `RoleAccess\(\s*.*\s*\)`}
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

func GetContractByID(vm *VM, id int32) *compile.ContractInfo {
	var tableID int64
	if id > ShiftContractID {
		tableID = int64(id - ShiftContractID)
		id = int32(tableID + vm.ShiftContract)
	}
	idcont := id
	if len(vm.Children) <= int(idcont) {
		return nil
	}
	if vm.Children[idcont] == nil || vm.Children[idcont].Type != compile.ObjContract {
		return nil
	}
	if tableID > 0 && vm.Children[idcont].GetContractInfo().Owner.TableID != tableID {
		return nil
	}
	return vm.Children[idcont].GetContractInfo()
}

func RunContractById(vm *VM, id int32, methods []string, extend map[string]any) error {
	info := GetContractByID(vm, id)
	if info == nil {
		return fmt.Errorf(`unknown contract id '%d'`, id)
	}
	return RunContractByName(vm, info.Name, methods, extend)
}

func RunContractByName(vm *VM, name string, methods []string, extend map[string]any) error {
	obj, ok := vm.Objects[name]
	if !ok {
		return fmt.Errorf(`unknown object '%s'`, name)
	}

	if obj.Type != compile.ObjContract {
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
		if obj.Type == compile.ObjFunc {
			fn := obj.GetCodeBlock()
			_, err := Run(vm, fn, extend)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Run executes CodeBlock in vm
func Run(vm *VM, block *compile.CodeBlock, extend map[string]any) (ret []any, err error) {
	if block == nil {
		return nil, fmt.Errorf(`code block is nil`)
	}
	rt := NewRuntime(vm, extend, extend[ExtendTxCost].(int64))
	ret, err = rt.Run(block)
	extend[ExtendTxCost] = rt.CostRemain()
	if err != nil {
		vm.logger.WithFields(log.Fields{"type": VMErr, "error": err, "original_contract": extend[ExtendOriginalContract], "this_contract": extend[ExtendThisContract]}).Error("running block in smart vm")
		return nil, err
	}
	return
}

func ObjectExists(vm *VM, name string, state uint32) bool {
	name = StateName(state, name)
	_, ok := vm.Objects[name]
	return ok
}

// SetExtendCost sets the cost of calling extended obj in vm
func (vm *VM) SetExtendCost(ext func(string) int64) *VM {
	vm.ExtCost = ext
	return vm
}

// SetFuncCallsDB Set up functions that can edit the database in vm
func (vm *VM) SetFuncCallsDB(funcCallsDB map[string]struct{}) *VM {
	vm.FuncCallsDB = funcCallsDB
	return vm
}

func (vm *VM) SetPreVar(preVar []string) *VM {
	vm.PreVar = append(vm.PreVar, preVar...)
	return vm
}

// FlushExtern switches off the extern mode of the compilation
func (vm *VM) FlushExtern() {
	vm.Extern = compile.IgnoreNone
	return
}

func (vm *VM) MergeCompConfig(conf *compile.CompConfig) *compile.CompConfig {
	conf.MakeConfig()
	for s, info := range vm.Objects {
		conf.Objects[s] = info
	}
	var m = make(map[string]struct{})

	for _, s := range vm.PreVar {
		m[s] = struct{}{}
	}
	for _, s := range conf.PreVar {
		m[s] = struct{}{}
	}
	var preVar []string
	for s := range m {
		preVar = append(preVar, s)
	}
	conf.PreVar = preVar
	vm.PreVar = preVar
	return conf
}

// Compile compiles a source code and loads the byte-code into the virtual machine,
func (vm *VM) Compile(input []rune, conf *compile.CompConfig) error {
	root, err := compile.CompileBlock(input, vm.MergeCompConfig(conf))
	if err != nil {
		return err
	}
	vm.FlushBlock(root)
	return nil
}

const flushMark = 1 << 20

// FlushBlock loads the compiled CodeBlock into the virtual machine
func (vm *VM) FlushBlock(root *compile.CodeBlock) {
	shift := len(vm.Children)
	for key, item := range root.Objects {
		if cur, ok := vm.Objects[key]; ok {
			switch item.Type {
			case compile.ObjContract:
				item.GetContractInfo().ID = cur.GetContractInfo().ID + flushMark
			case compile.ObjFunc:
				item.GetFuncInfo().ID = cur.GetFuncInfo().ID + flushMark
				vm.Objects[key].Value = item.Value
			}
		}
		vm.Objects[key] = item
	}
	for _, item := range root.Children {
		if item == nil {
			continue
		}
		switch item.Type {
		case compile.ObjContract:
			if item.GetContractInfo().ID > flushMark {
				item.GetContractInfo().ID -= flushMark
				vm.Children[item.GetContractInfo().ID] = item
				shift--
				continue
			}

			item.Parent = vm.CodeBlock
			item.GetContractInfo().ID += uint32(shift)
		case compile.ObjFunc:
			if item.GetFuncInfo().ID > flushMark {
				item.GetFuncInfo().ID -= flushMark
				vm.Children[item.GetFuncInfo().ID] = item
				shift--
				continue
			}
			item.Parent = vm.CodeBlock
			item.GetFuncInfo().ID += uint32(shift)
		}
		vm.Children = append(vm.Children, item)
	}
}
