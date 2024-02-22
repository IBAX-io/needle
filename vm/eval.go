package vm

import "github.com/IBAX-io/needle/compiler"

// evalCode is a struct that represents a piece of code that can be evaluated.
// It contains the source code as a string and a CodeBlock that represents the compiled code.
type evalCode struct {
	Source string
	Code   *compiler.CodeBlock
}

// evals is a map that associates a checksum of the source code with the corresponding evalCode.
var evals = make(map[uint64]*evalCode)

// CompileEval compiles the source code and stores it in the evals map.
func (vm *VM) CompileEval(input string, state uint32) error {
	source := `func eval bool { return ` + input + `}`
	if input == `1` || input == `0` {
		source = `
		func eval bool { 
			if ` + input + ` == 1 {
				return true
			} else {
				return false
			}
		}`
	}

	block, err := compiler.CompileBlock([]rune(source), vm.MergeCompConfig(&compiler.CompConfig{Owner: &compiler.OwnerInfo{StateId: state}}))
	if err != nil {
		return err
	}
	crc := CalcChecksum([]byte(input))
	evals[crc] = &evalCode{Source: input, Code: block}
	return nil
}

// EvalIf runs the conditional expression. It compiles the source code before that if that's necessary.
func (vm *VM) EvalIf(input string, state uint32, extend map[string]any) (bool, error) {
	if len(input) == 0 {
		return true, nil
	}
	crc := CalcChecksum([]byte(input))
	if eval, ok := evals[crc]; !ok || eval.Source != input {
		if err := vm.CompileEval(input, state); err != nil {
			return false, err
		}
	}
	ret, err := NewRuntime(vm, extend, extend[ExtendTxCost].(int64)).Run(evals[crc].Code.Children[0])
	if err == nil {
		if len(ret) == 0 {
			return false, nil
		}
		return valueToBool(ret[0]), nil
	}
	return false, err
}
