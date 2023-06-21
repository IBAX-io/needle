package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/IBAX-io/needle/compile"
	log "github.com/sirupsen/logrus"
)

const (
	Extend_type                = `type`
	Extend_time                = `time`
	Extend_ecosystem_id        = `ecosystem_id`
	Extend_node_position       = `node_position`
	Extend_block               = `block`
	Extend_key_id              = `key_id`
	Extend_account_id          = `account_id`
	Extend_block_key_id        = `block_key_id`
	Extend_parent              = `parent`
	Extend_txcost              = `txcost`
	Extend_txhash              = `txhash`
	Extend_result              = `result`
	Extend_sc                  = `sc`
	Extend_contract            = `contract`
	Extend_block_time          = `block_time`
	Extend_original_contract   = `original_contract`
	Extend_this_contract       = `this_contract`
	Extend_guest_key           = `guest_key`
	Extend_guest_account       = `guest_account`
	Extend_black_hole_key      = `black_hole_key`
	Extend_black_hole_account  = `black_hole_account`
	Extend_white_hole_key      = `white_hole_key`
	Extend_white_hole_account  = `white_hole_account`
	Extend_pre_block_data_hash = `pre_block_data_hash`
	Extend_gen_block           = `gen_block`
	Extend_time_limit          = `time_limit`

	Extend_rt    = `rt`
	Extend_stack = `stack`
	Extend_loop  = `loop_`
)
const (
	//system variable cannot be changed
	sysVars_block               = `block`
	sysVars_block_key_id        = `block_key_id`
	sysVars_block_time          = `block_time`
	sysVars_data                = `data`
	sysVars_ecosystem_id        = `ecosystem_id`
	sysVars_key_id              = `key_id`
	sysVars_account_id          = `account_id`
	sysVars_node_position       = `node_position`
	sysVars_parent              = `parent`
	sysVars_original_contract   = `original_contract`
	sysVars_sc                  = `sc`
	sysVars_contract            = `contract`
	sysVars_stack               = `stack`
	sysVars_this_contract       = `this_contract`
	sysVars_time                = `time`
	sysVars_type                = `type`
	sysVars_txcost              = `txcost`
	sysVars_txhash              = `txhash`
	sysVars_guest_key           = `guest_key`
	sysVars_guest_account       = `guest_account`
	sysVars_black_hole_key      = `black_hole_key`
	sysVars_white_hole_key      = `white_hole_key`
	sysVars_black_hole_account  = `black_hole_account`
	sysVars_white_hole_account  = `white_hole_account`
	sysVars_gen_block           = `gen_block`
	sysVars_time_limit          = `time_limit`
	sysVars_pre_block_data_hash = `pre_block_data_hash`
)

var sysVars = map[string]struct{}{
	sysVars_block:               {},
	sysVars_block_key_id:        {},
	sysVars_block_time:          {},
	sysVars_data:                {},
	sysVars_ecosystem_id:        {},
	sysVars_key_id:              {},
	sysVars_account_id:          {},
	sysVars_node_position:       {},
	sysVars_parent:              {},
	sysVars_original_contract:   {},
	sysVars_sc:                  {},
	sysVars_contract:            {},
	sysVars_stack:               {},
	sysVars_this_contract:       {},
	sysVars_time:                {},
	sysVars_type:                {},
	sysVars_txcost:              {},
	sysVars_txhash:              {},
	sysVars_guest_key:           {},
	sysVars_guest_account:       {},
	sysVars_black_hole_key:      {},
	sysVars_black_hole_account:  {},
	sysVars_white_hole_key:      {},
	sysVars_white_hole_account:  {},
	sysVars_gen_block:           {},
	sysVars_time_limit:          {},
	sysVars_pre_block_data_hash: {},
}

func isSysVar(name string) bool {
	if _, ok := sysVars[name]; ok || strings.HasPrefix(name, Extend_loop) {
		return true
	}
	return false
}

type extendInfo struct {
	genBlock  bool
	txCost    int64
	timeLimit int64
}

func (rt *Runtime) loadExtendBy(key string) *extendInfo {
	var e = &extendInfo{}
	extend, ok := rt.extend[key]
	if !ok {
		return e
	}
	switch key {
	case Extend_gen_block:
		e.genBlock, _ = extend.(bool)
	case Extend_txcost:
		e.txCost, _ = extend.(int64)
	case Extend_time_limit:
		e.timeLimit, _ = extend.(int64)
	}
	return e
}

// ExecContract runs the name contract where txs contains the list of parameters and
// params are the values of parameters
func ExecContract(rt *Runtime, name, txs string, params ...any) (any, error) {
	obj, ok := rt.vm.Objects[name]
	if !ok || obj.Type != compile.ObjContract {
		log.WithFields(log.Fields{"contract_name": name, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf(eUnknownContract, name)
	}
	logger := log.WithFields(log.Fields{"contract_name": name, "type": ContractError})

	//check if there is loop in contract
	if _, ok := rt.extend[Extend_loop+name]; ok {
		logger.WithFields(log.Fields{"type": ContractError, "contract_name": name}).Error("there is loop in contract")
		return nil, fmt.Errorf(eContractLoop, name)
	}
	rt.extend[Extend_loop+name] = true
	defer delete(rt.extend, Extend_loop+name)

	//save previous extend variables of current contract
	prevExtend := make(map[string]any)
	for key, item := range rt.extend {
		if rt.vm.AssertVar(key) {
			continue
		}
		prevExtend[key] = item
		delete(rt.extend, key)
	}
	extVars, err := genExtVars(obj.GetContractInfo(), txs, params...)
	if err != nil {
		return nil, err
	}

	// define extend variables to next contract from parameters
	for key, item := range extVars {
		rt.extend[key] = item
	}

	prevthis := rt.extend[Extend_this_contract]
	_, nameContract := ParseName(name)
	rt.extend[Extend_this_contract] = nameContract

	prevparent := rt.extend[Extend_parent]
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

	if err := rt.SubCost(CostContract); err != nil {
		return nil, err
	}

	var (
		stack Stacker
	)
	if stack, ok = rt.extend[Extend_sc].(Stacker); ok {
		if err := stack.AppendStack(name); err != nil {
			return nil, err
		}
	}

	for _, method := range []string{`conditions`, `action`} {
		if block, ok := obj.GetCodeBlock().Objects[method]; ok && block.Type == compile.ObjFunc {
			rtemp := NewRuntime(rt.vm, rt.cost)
			rt.extend[Extend_parent] = parent
			_, err = rtemp.Run(block.GetCodeBlock(), rt.extend)
			rt.cost = rtemp.cost
			if err != nil {
				//logger.WithFields(log.Fields{"error": err, "method_name": method, "type": ContractError}).Error("executing contract method")
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
	rt.extend[Extend_parent] = prevparent
	rt.extend[Extend_this_contract] = prevthis
	result := rt.extend[Extend_result]
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
func CallContract(rt *Runtime, name string, params *compile.Map) (any, error) {
	name = StateName(rt.vm.GetOwnerInfo().StateID, name)
	_, ok := rt.vm.Objects[name]
	if !ok {
		log.WithFields(log.Fields{"contract_name": name, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf(eUnknownContract, name)
	}
	if params == nil {
		params = compile.NewMap()
	}
	return ExecContract(rt, name, strings.Join(params.Keys(), `,`), params.Values())
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

func genExtVars(contract *compile.ContractInfo, txs string, params ...any) (map[string]any, error) {
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
			return nil, fmt.Errorf("'%s' parameter is not required", par)
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
