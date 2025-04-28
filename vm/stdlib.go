package vm

import (
	"fmt"
	"strings"

	"github.com/IBAX-io/needle/compiler"
	log "github.com/sirupsen/logrus"
)

// ExecContract runs the name contract where fields contains the list of parameters and
// params are the values of parameters.
func ExecContract(rt *Runtime, name, fields string, params ...any) (any, error) {
	if err := rt.SubCost(CostContract); err != nil {
		return nil, err
	}

	obj, ok := rt.vm.CodeBlock.Objects[name]
	if !ok || !obj.IsCodeBlockContract() {
		return nil, fmt.Errorf(eUnknownContract, name)
	}

	// check if there is loop in contract
	if _, ok := rt.used[name]; ok {
		return nil, fmt.Errorf(eContractLoop, name)
	}
	rt.used[name] = struct{}{}
	defer delete(rt.used, name)
	// save previous extend variables of current contract
	prevExtend := make(map[string]any)
	for key, item := range rt.extend {
		if rt.vm.CodeBlock.AssertVar(key) {
			continue
		}
		prevExtend[key] = item
		delete(rt.extend, key)
	}
	extVars, err := obj.GetContractInfo().MakeParams(fields, params)
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
		b := rt.blocks[i].Block
		if b.Type == compiler.CodeBlockFunction &&
			b.Parent != nil &&
			b.Parent.Type == compiler.CodeBlockContract {
			parent = b.Parent.GetContractInfo().Name
			fid, fname := ParseName(parent)
			cid, _ := ParseName(name)
			if len(fname) > 0 {
				if fid == 0 {
					parent = "@" + fname
				} else if fid == cid {
					parent = fname
				}
			}
			break
		}
	}

	var stack Stacker
	if stack, ok = rt.extend[ExtendSc].(Stacker); ok {
		if err := stack.AppendStack(name); err != nil {
			return nil, err
		}
	}

	for _, method := range []compiler.Token{compiler.CONDITIONS, compiler.ACTION} {
		block, ok := obj.GetCodeBlock().Objects[method.ToString()]
		if !(ok && block.IsCodeBlockFunction()) {
			continue
		}
		rtemp := NewRuntime(rt.vm, rt.extend, rt.costRemain)
		rt.extend[ExtendParentContract] = parent
		rtemp.used = rt.used
		_, err = rtemp.Run(block.GetCodeBlock())
		rt.costRemain = rtemp.CostRemain()
		if err != nil {
			break
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
		if rt.vm.CodeBlock.AssertVar(key) {
			continue
		}
		delete(rt.extend, key)
	}

	for key, item := range prevExtend {
		rt.extend[key] = item
	}
	return result, nil
}

// CallContract executes the name contract in the state with specified parameters.
func CallContract(rt *Runtime, state uint32, name string, params *compiler.Map) (any, error) {
	name = compiler.StateName(state, name)
	_, ok := rt.vm.CodeBlock.Objects[name]
	if !ok {
		log.WithFields(log.Fields{"contract_name": name, "type": ContractError}).Error("unknown contract")
		return nil, fmt.Errorf(eUnknownContract, name)
	}
	if params == nil {
		params = compiler.NewMap()
	}
	return ExecContract(rt, name, strings.Join(params.Keys(), `,`), params.Values()...)
}

// GetSettings returns the value of the setting of the contract.
func GetSettings(rt *Runtime, cntname, name string) (any, error) {
	contract, found := rt.vm.CodeBlock.Objects[cntname]
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

// MemoryUsage returns the memory usage of the runtime.
func MemoryUsage(rt *Runtime) int64 {
	return rt.mem
}
