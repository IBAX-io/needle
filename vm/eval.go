package vm

import (
	"sync"

	"github.com/IBAX-io/needle/compiler"
)

// evalCode is a struct that represents a piece of code that can be evaluated.
// It contains the source code as a string and a CodeBlock that represents the compiled code.
type evalCode struct {
	Source string
	Code   *compiler.CodeBlock
}

// evals is a map that associates a checksum of the source code with the corresponding evalCode.
var evals sync.Map

func Evalsource(input string) string {
	if input == `1` || input == `0` {
		return `
		func eval bool { 
			if ` + input + ` == 1 {
				return true
			} else {
				return false
			}
		}`
	}
	return `func eval bool { return ` + input + `}`
}

// CompileEval compiles the source code and stores it in the evals map.
func (vm *VM) CompileEval(input string, state uint32) error {
	source := Evalsource(input)
	conf := &compiler.Config{Owner: &compiler.OwnerInfo{StateId: state}}
	block, err := compiler.CompileBlock([]rune(source), vm.MergeCompilerConfig(conf))
	if err != nil {
		return err
	}
	crc := CalcChecksum([]byte(input))
	evals.Store(crc, &evalCode{Source: input, Code: block})
	return nil
}

// EvalIf runs the conditional expression. It compiles the source code before that if that's necessary.
func (vm *VM) EvalIf(input string, state uint32, extend map[string]any) (bool, error) {
	if len(input) == 0 {
		return true, nil
	}
	crc := CalcChecksum([]byte(input))
	eval, ok := evals.Load(crc)
	if !ok || eval.(*evalCode).Source != input {
		if err := vm.CompileEval(input, state); err != nil {
			return false, err
		}
		eval, _ = evals.Load(crc)
	}
	code, _ := eval.(*evalCode)
	ret, err := NewRuntime(vm, extend, extend[ExtendTxCost].(int64)).Run(code.Code.Children[0])
	if err == nil {
		if len(ret) == 0 {
			return false, nil
		}
		return valueToBool(ret[0]), nil
	}
	return false, err
}
