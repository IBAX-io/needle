package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/IBAX-io/needle/compile"
	log "github.com/sirupsen/logrus"
)

const (
	ExtendParentContract   = `parent_contract`
	ExtendOriginalContract = `original_contract`
	ExtendThisContract     = `this_contract`
	ExtendTimeLimit        = `time_limit`
	ExtendGenBlock         = `gen_block`
	ExtendTxCost           = `txcost`
	ExtendStack            = `stack`
	ExtendSc               = `sc`
	ExtendRt               = `rt`

	ExtendResult = `result`
)

const (
	TagOptional = "optional"
)

// system variable cannot be changed through the contract
var sysVars = map[string]struct{}{
	ExtendParentContract:   {},
	ExtendOriginalContract: {},
	ExtendThisContract:     {},
	ExtendTimeLimit:        {},
	ExtendGenBlock:         {},
	ExtendTxCost:           {},
	ExtendStack:            {},
	ExtendRt:               {},
	ExtendSc:               {},
}

type extendInfo struct {
	genBlock  bool
	txCost    int64
	timeLimit int64
	rt        *Runtime
}

func (rt *Runtime) setExtendBy(key string, value any) {
	if _, ok := rt.extend[key]; !ok {
		rt.extend[key] = value
	}
}

func (rt *Runtime) loadExtendBy(key string) *extendInfo {
	var e = &extendInfo{}
	extend, ok := rt.extend[key]
	if !ok {
		return e
	}
	switch key {
	case ExtendGenBlock:
		e.genBlock, _ = extend.(bool)
	case ExtendTxCost:
		e.txCost, _ = extend.(int64)
	case ExtendTimeLimit:
		e.timeLimit, _ = extend.(int64)
	case ExtendRt:
		e.rt, _ = extend.(*Runtime)
	}
	return e
}

// ExecContract runs the name contract where txs contains the list of parameters and
// params are the values of parameters
func ExecContract(rt *Runtime, name, txs string, params ...any) (any, error) {
	if err := rt.SubCost(CostContract); err != nil {
		return nil, err
	}

	obj, ok := rt.vm.Objects[name]
	if !ok || obj.Type != compile.ObjContract {
		return nil, fmt.Errorf(eUnknownContract, name)
	}

	//check if there is loop in contract
	if _, ok := rt.used[name]; ok {
		return nil, fmt.Errorf(eContractLoop, name)
	}
	rt.used[name] = struct{}{}
	defer delete(rt.used, name)
	//save previous extend variables of current contract
	prevExtend := make(map[string]any)
	for key, item := range rt.extend {
		if rt.vm.AssertVar(key) {
			continue
		}
		prevExtend[key] = item
		delete(rt.extend, key)
	}
	extVars, err := genExtVars(obj.GetContractInfo(), txs, params)
	if err != nil {
		return nil, err
	}

	// define extend variables to next contract from parameters
	for key, item := range extVars {
		rt.extend[key] = item
	}

	prevthis := rt.extend[ExtendThisContract]
	_, nameContract := ParseName(name)
	rt.extend[ExtendThisContract] = nameContract

	prevparent := rt.extend[ExtendParentContract]
	parent := ``
	for i := len(rt.blocks) - 1; i >= 0; i-- {
		var b = rt.blocks[i].Block
		if b.Type == compile.ObjFunc &&
			b.Parent != nil &&
			b.Parent.Type == compile.ObjContract {
			parent = b.Parent.GetContractInfo().Name
			fid, fname := ParseName(parent)
			cid, _ := ParseName(name)
			if len(fname) > 0 {
				if fid == 0 {
					parent = `@` + fname
				} else if fid == cid {
					parent = fname
				}
			}
			break
		}
	}

	var (
		stack Stacker
	)
	if stack, ok = rt.extend[ExtendSc].(Stacker); ok {
		if err := stack.AppendStack(name); err != nil {
			return nil, err
		}
	}
	for _, method := range []string{`conditions`, `action`} {
		if block, ok := obj.GetCodeBlock().Objects[method]; ok && block.Type == compile.ObjFunc {
			rtemp := NewRuntime(rt.vm, rt.extend, rt.costRemain)
			rt.extend[ExtendParentContract] = parent
			rtemp.used = rt.used
			_, err = rtemp.Run(block.GetCodeBlock())
			rt.costRemain = rtemp.CostRemain()
			if err != nil {
				break
			}
		}
	}
	if stack != nil {
		stack.PopStack(name)
	}
	if err != nil {
		return nil, err
	}
	rt.extend[ExtendParentContract] = prevparent
	rt.extend[ExtendThisContract] = prevthis
	result := rt.extend[ExtendResult]
	for key := range rt.extend {
		if rt.vm.AssertVar(key) {
			continue
		}
		delete(rt.extend, key)
	}

	for key, item := range prevExtend {
		rt.extend[key] = item
	}
	return result, nil
}

// CallContract executes the name contract in the state with specified parameters
func CallContract(rt *Runtime, state uint32, name string, params *compile.Map) (any, error) {
	name = compile.StateName(uint32(state), name)
	_, ok := rt.vm.Objects[name]
	if !ok {
		log.WithFields(log.Fields{"contract_name": name, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf(eUnknownContract, name)
	}
	if params == nil {
		params = compile.NewMap()
	}
	return ExecContract(rt, name, strings.Join(params.Keys(), `,`), params.Values()...)
}

// GetSettings returns the value of the parameter of contract
func GetSettings(rt *Runtime, cntname, name string) (any, error) {
	contract, found := rt.vm.Objects[cntname]
	if !found || contract.GetCodeBlock() == nil {
		log.WithFields(log.Fields{"contract_name": cntname, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf("unknown contract %s", cntname)
	}
	info := contract.GetContractInfo()
	if info != nil {
		if val, ok := info.Settings[name]; ok {
			return val, nil
		}
	}
	return ``, nil
}

func MemoryUsage(rt *Runtime) int64 {
	return rt.mem
}

func genExtVars(contract *compile.ContractInfo, txs string, params []any) (map[string]any, error) {
	pars := strings.Split(txs, `,`)
	param := make(map[string]struct{})
	for _, par := range pars {
		if _, ok := param[par]; ok {
			return nil, fmt.Errorf("duplicate parameter '%s'", par)
		}
		param[par] = struct{}{}
	}
	if len(pars) != len(params) {
		return nil, fmt.Errorf("wrong number of parameters, expected %d, got %d", len(pars), len(params))
	}

	extVars := make(map[string]any)
	txMap := contract.TxMap()

	for i, par := range pars {
		if len(par) == 0 {
			continue
		}
		_, ok := txMap[par]
		if !ok {
			continue
		}
		if len(par) > 0 {
			extVars[par] = params[i]
		}
	}
	for _, tx := range txMap {
		if _, ok := param[tx.Name]; !ok {
			if !strings.Contains(tx.Tags, TagOptional) {
				return nil, fmt.Errorf(eUndefinedParam, tx.Name)
			}
			extVars[tx.Name] = reflect.Zero(tx.Type).Interface()
		}
	}

	return extVars, nil
}
