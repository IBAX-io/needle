package vm

import (
	"fmt"
	"github.com/IBAX-io/needle/compile"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
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

// ExecContract runs the name contract where txs contains the list of parameters and
// params are the values of parameters
func ExecContract(rt *Runtime, name, txs string, params ...any) (any, error) {
	contract, ok := rt.vm.Objects[name]
	if !ok {
		log.WithFields(log.Fields{"contract_name": name, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf(eUnknownContract, name)
	}
	logger := log.WithFields(log.Fields{"contract_name": name, "type": ContractError})
	parnames := make(map[string]bool)
	pars := strings.Split(txs, `,`)
	if len(pars) != len(params) {
		logger.WithFields(log.Fields{"contract_params_len": len(pars), "contract_params_len_needed": len(params), "type": ContractError}).Error("wrong contract parameters pars")
		return ``, errContractPars
	}
	for _, ipar := range pars {
		parnames[ipar] = true
	}
	if _, ok := rt.extend[Extend_loop+name]; ok {
		logger.WithFields(log.Fields{"type": ContractError, "contract_name": name}).Error("there is loop in contract")
		return nil, fmt.Errorf(eContractLoop, name)
	}
	rt.extend[Extend_loop+name] = true
	defer delete(rt.extend, Extend_loop+name)

	prevExtend := make(map[string]any)
	for key, item := range rt.extend {
		if rt.vm.AssertVar(key) {
			continue
		}
		prevExtend[key] = item
		delete(rt.extend, key)
	}
	cblock := contract.GetCodeBlock()
	if cblock.GetContractInfo().Tx != nil {
		for _, tx := range *cblock.GetContractInfo().Tx {
			if !parnames[tx.Name] {
				if !strings.Contains(tx.Tags, TagOptional) {
					logger.WithFields(log.Fields{"transaction_name": tx.Name, "type": ContractError}).Error("transaction not defined")
					return ``, fmt.Errorf(eUndefinedParam, tx.Name)
				}
				rt.extend[tx.Name] = reflect.New(tx.Type).Elem().Interface()
			}
		}
	}
	for i, ipar := range pars {
		if len(ipar) > 0 {
			rt.extend[ipar] = params[i]
		}
	}
	prevthis := rt.extend[Extend_this_contract]
	_, nameContract := ParseName(name)
	rt.extend[Extend_this_contract] = nameContract

	prevparent := rt.extend[Extend_parent]
	parent := ``
	for i := len(rt.blocks) - 1; i >= 0; i-- {
		if rt.blocks[i].Block.Type == compile.ObjectType_Func && rt.blocks[i].Block.Parent != nil &&
			rt.blocks[i].Block.Parent.Type == compile.ObjectType_Contract {
			parent = rt.blocks[i].Block.Parent.GetContractInfo().Name
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
		err   error
	)
	if stack, ok = rt.extend[Extend_sc].(Stacker); ok {
		if err := stack.AppendStack(name); err != nil {
			return nil, err
		}
	}

	for _, method := range []string{`conditions`, `action`} {
		if block, ok := (*cblock).Objects[method]; ok && block.Type == compile.ObjectType_Func {
			rtemp := NewRuntime(rt.vm, rt.cost)
			rt.extend[Extend_parent] = parent
			_, err = rtemp.Run(block.GetCodeBlock(), rt.extend)
			rt.cost = rtemp.cost
			if err != nil {
				logger.WithFields(log.Fields{"error": err, "method_name": method, "type": ContractError}).Error("executing contract method")
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
func CallContract(rt *Runtime, state uint32, name string, params *compile.Map) (any, error) {
	name = StateName(state, name)
	contract, ok := rt.vm.Objects[name]
	if !ok {
		log.WithFields(log.Fields{"contract_name": name, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf(eUnknownContract, name)
	}
	if params == nil {
		params = compile.NewMap()
	}
	logger := log.WithFields(log.Fields{"contract_name": name, "type": ContractError})
	names := make([]string, 0)
	vals := make([]any, 0)
	if contract.GetContractInfo().Tx != nil {
		for _, tx := range *contract.GetContractInfo().Tx {
			val, ok := params.Get(tx.Name)
			if !ok {
				if !strings.Contains(tx.Tags, TagOptional) {
					logger.WithFields(log.Fields{"transaction_name": tx.Name}).Error("transaction not defined")
					return nil, fmt.Errorf(eUndefinedParam, tx.Name)
				}
				val = reflect.New(tx.Type).Elem().Interface()
			}
			names = append(names, tx.Name)
			vals = append(vals, val)
		}
	}
	if len(vals) == 0 {
		vals = append(vals, ``)
	}
	return ExecContract(rt, name, strings.Join(names, `,`), vals...)
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
