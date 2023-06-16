package vm

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/IBAX-io/needle/compile"

	log "github.com/sirupsen/logrus"
)

type GlobalVm struct {
	mu      sync.Mutex
	smartVM *VM
}

var (
	_vm *GlobalVm
)

func init() {
	_vm = &GlobalVm{
		smartVM: NewVM(),
	}
}

const (
	TagFile      = "file"
	TagAddress   = "address"
	TagSignature = "signature"
	TagOptional  = "optional"
)

// VM is the main type of the virtual machine
type VM struct {
	*compile.CodeBlock
	ExtCost       func(string) int64 // the cost of executing an extend function
	FuncCallsDB   map[string]struct{}
	Extern        bool  // extern mode of compilation
	ShiftContract int64 // id of the first contract
	logger        *log.Entry
}

// Stacker represents interface for working with call stack
type Stacker interface {
	AppendStack(fn string) error
	PopStack(fn string)
}

// NewVM creates a new virtual machine
func NewVM() *VM {
	auto := map[string]string{"*vm.Runtime": "rt"}
	var fn = []compile.ExtendFunc{
		{Name: "Settings", Func: GetSettings, AutoPars: auto},
		{Name: "MemoryUsage", Func: MemoryUsage, AutoPars: auto},
		{"ExecContract", ExecContract, true, auto},
		{"CallContract", CallContract, true, auto},
		{Name: "Println", Func: fmt.Println},
		{Name: "Sprintf", Func: fmt.Sprintf},
	}
	var v []string
	for p := range sysVars {
		v = append(v, p)
	}
	vm := &VM{
		CodeBlock:   compile.NewCodeBlock(compile.NewExtendData(nil, fn, v)),
		Extern:      true,
		FuncCallsDB: make(map[string]struct{}),
		logger:      log.WithFields(log.Fields{"type": VMErr, "extern": true}),
	}
	return vm
}

// GetVM is returning smart vm
func GetVM() *VM {
	_vm.mu.Lock()
	defer _vm.mu.Unlock()
	return _vm.smartVM
}

var smartObjects map[string]*compile.ObjInfo
var children uint32

func SavepointSmartVMObjects() {
	smartObjects = make(map[string]*compile.ObjInfo)
	for k, v := range GetVM().Objects {
		smartObjects[k] = v
	}
	children = uint32(len(GetVM().Children))
}

// Call executes the name object with the specified params and extended variables and functions
func (vm *VM) Call(name string, params []any, extend map[string]any) (ret []any, err error) {
	split := strings.Split(name, ".")
	obj := vm.GetObjByName(split[0])
	if obj == nil {
		return nil, fmt.Errorf(`object %s is empty`, name)
	}
	switch obj.Type {
	case compile.ObjContract:
		rt := NewRuntime(vm, extend[Extend_txcost].(int64))
		ret, err = rt.Run(obj.GetCodeBlock().GetObjByName(split[1]).GetCodeBlock(), extend)
		extend[Extend_txcost] = rt.Cost()
	case compile.ObjFunc:
		rt := NewRuntime(vm, extend[Extend_txcost].(int64))
		ret, err = rt.Run(obj.GetCodeBlock(), extend)
		extend[Extend_txcost] = rt.Cost()
	case compile.ObjExtFunc:
		ret = obj.GetExtFuncInfo().Call(params)
	default:
		return nil, fmt.Errorf(`unknown object %s for call`, name)
	}
	return ret, err
}

func RollbackSmartVMObjects() {
	GetVM().Objects = make(map[string]*compile.ObjInfo)
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

func VMCompileEval(vm *VM, src string, prefix uint32) error {
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
	re := regexp.MustCompile(`^@?[\d\w_]+$`)
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

func VMGetContractByID(vm *VM, id int32) *compile.ContractInfo {
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
	info := VMGetContractByID(vm, id)
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
	extend[Extend_txcost] = extend[Extend_txcost].(int64) - CostContract - ContractBaseCost(contract)
	if extend[Extend_txcost].(int64) < 0 {
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
			_, err := VMRun(vm, fn, extend)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// VMRun executes CodeBlock in vm
func VMRun(vm *VM, block *compile.CodeBlock, extend map[string]any) (ret []any, err error) {
	if block == nil {
		return nil, fmt.Errorf(`code block is nil`)
	}
	var cost int64
	if ecost, ok := extend[Extend_txcost]; ok {
		cost = ecost.(int64)
	}
	rt := NewRuntime(vm, cost)
	ret, err = rt.Run(block, extend)
	extend[Extend_txcost] = rt.Cost()
	if err != nil {
		vm.logger.WithFields(log.Fields{"type": VMErr, "error": err, "original_contract": extend[Extend_original_contract], "this_contract": extend[Extend_this_contract], "ecosystem_id": extend[Extend_ecosystem_id]}).Error("running block in smart vm")
		return nil, err
	}
	return
}

func VMObjectExists(vm *VM, name string, state uint32) bool {
	name = StateName(state, name)
	_, ok := vm.Objects[name]
	return ok
}

// SetExtendCost sets the cost of calling extended obj in vm
func (vm *VM) SetExtendCost(ext func(string) int64) {
	vm.ExtCost = ext
}

// SetFuncCallsDB Set up functions that can edit the database in vm
func (vm *VM) SetFuncCallsDB(funcCallsDB map[string]struct{}) {
	vm.FuncCallsDB = funcCallsDB
}

// FlushExtern switches off the extern mode of the compilation
func (vm *VM) FlushExtern() {
	vm.Extern = false
	return
}

// Compile compiles a source code and loads the byte-code into the virtual machine,
func (vm *VM) Compile(input []rune, ext *compile.ExtendData) error {
	var d compile.ExtendData
	for s, info := range vm.Objects {
		if info.Type != compile.ObjExtFunc {
			continue
		}
		var fn = compile.ExtendFunc{
			Name:     s,
			Func:     info.GetExtFuncInfo().Func,
			CanWrite: info.GetExtFuncInfo().CanWrite,
			AutoPars: make(map[string]string),
		}
		fobj := reflect.ValueOf(fn.Func).Type()
		for i := 0; i < fobj.NumIn(); i++ {
			if info.GetExtFuncInfo().Auto[i] != "" {
				fn.AutoPars[fobj.In(i).String()] = info.GetExtFuncInfo().Auto[i]
			}
		}
		d.Func = append(d.Func, fn)
	}
	ext.Func = append(d.Func, ext.Func...)
	ext.PreVar = append(vm.PredeclaredVar, ext.PreVar...)
	root, err := compile.CompileBlock(input, ext)
	if err != nil {
		return err
	}
	vm.FlushBlock(root)
	return nil
}

// FlushBlock loads the compiled CodeBlock into the virtual machine
func (vm *VM) FlushBlock(root *compile.CodeBlock) {
	shift := len(vm.Children)
	for key, item := range root.Objects {
		if cur, ok := vm.Objects[key]; ok {
			switch item.Type {
			case compile.ObjContract:
				root.Objects[key].GetContractInfo().ID = cur.GetContractInfo().ID + compile.FlushMark
			case compile.ObjFunc:
				root.Objects[key].GetFuncInfo().ID = cur.GetFuncInfo().ID + compile.FlushMark
				vm.Objects[key].Value = root.Objects[key].Value
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
			if item.GetContractInfo().ID > compile.FlushMark {
				item.GetContractInfo().ID -= compile.FlushMark
				vm.Children[item.GetContractInfo().ID] = item
				shift--
				continue
			}
			item.Parent = vm.CodeBlock
			item.GetContractInfo().ID += uint32(shift)
		case compile.ObjFunc:
			if item.GetFuncInfo().ID > compile.FlushMark {
				item.GetFuncInfo().ID -= compile.FlushMark
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

func LoadSysFuncs(vm *VM, state int) error {
	code := `
func DBFind(table string).Select(query string).Columns(columns string).Where(where map)
.WhereId(id int).Order(order string).Limit(limit int).Offset(offset int).Group(group string).All(all bool) array {
        return DBSelect(table, columns, id, order, offset, limit, where, query, group, all)
}

func One(list array, name string) string {
        if list {
                var row map
                row = list[0]
                if Contains(name, "->") {
                        var colfield array
                        var val string
                        colfield = Split(ToLower(name), "->")
                        val = row[Join(colfield, ".")]
                        if !val && row[colfield[0]] {
                                var fields map
                                var i int
                                fields = JSONDecode(row[colfield[0]])
                                val = fields[colfield[1]]
                                i = 2
                                while i < Len(colfield) {
                                        if GetType(val) == "map[string]interface {}" {
                                                val = val[colfield[i]]
                                                if !val {
                                                        break
                                                }
                                                i = i + 1
                                        } else {
                                                break
                                        }
                                }
                        }
                        if !val {
                                return ""
                        }
                        return val
                }
                return Str(row[name])
        }
        return ""
}

func Row(list array) map {
        var ret map
        if list {
                ret = list[0]
        }
        return ret
}

func DBRow(table string).Columns(columns string).Where(where map)
.WhereId(id int).Order(order string) map {

        var result array
        result = DBFind(table).Columns(columns).Where(where).WhereId(id).Order(order)

        var row map
        if Len(result) > 0 {
                row = result[0]
        }

        return row
}

func ConditionById(table string, validate bool) {
        var row map
        row = DBRow(table).Columns("conditions").WhereId($Id)
        if !row["conditions"] {
                error Sprintf("Item %d has not been found", $Id)
        }

        Eval(row["conditions"])

        if validate {
                ValidateCondition($Conditions,$ecosystem_id)
        }
}

func CurrentKeyFromAccount(account string) int {
        var row map
        row = DBRow("@1keys").Columns("id").Where({"account": account, "deleted": 0})
        if row {
                return row["id"]
        }
        return 0
}
`
	return vm.Compile([]rune(code), compile.NewExtendData(&compile.OwnerInfo{StateID: uint32(state)}, nil, nil))
}
