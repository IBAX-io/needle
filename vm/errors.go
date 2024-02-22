package vm

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	eContractLoop        = `there is loop in %s contract`
	eSysVar              = `system variable $%s cannot be changed`
	eTypeParam           = `parameter %d has wrong type`
	eUndefinedParam      = `%s is not defined`
	eUnknownContract     = `unknown contract %s`
	eArrIndex            = `index of array cannot be type %s`
	eMapIndex            = `index of map cannot be type %s`
	eWrongVar            = `wrong var %v`
	eConditionNotAllowed = `condition %s is not allowed`
)

// MaxErrLen is the maximum length of error, over which it will be truncated.
const MaxErrLen = 150

const (
	ContractError     = "Contract"
	JSONMarshallError = "JSONMarshall"
	ConversionError   = "Conversion"
	VMErr             = "VM"
)

var (
	ErrMemoryLimit = errors.New("memory limit exceeded")
	ErrVMTimeLimit = errors.New(`time limit exceeded`)
)

var (
	errWrongCountPars     = errors.New(`wrong count of parameters`)
	errUnsupportedType    = errors.New(`unsupported combination of types in the operator`)
	errMaxArrayIndex      = errors.New(`the index is out of range`)
	errMaxMapCount        = errors.New(`the maximum length of map`)
	errSelfAssignment     = errors.New(`self assignment`)
	errIncorrectParameter = errors.New(`incorrect parameter of the condition function`)
)

// ExtFuncErr represents error of external function
type ExtFuncErr struct {
	Name  string
	Value any
}

func (e ExtFuncErr) Error() string {
	return fmt.Sprint(e.Value)
}

// VMError represents error of VM
type VMError struct {
	Type   string `json:"type"`
	Err    any    `json:"error"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

func (e VMError) Error() string {
	errText := fmt.Sprintf(`%v`, e.Err)
	if len(errText) > MaxErrLen {
		errText = errText[:MaxErrLen] + `...`
	}
	out, err := json.Marshal(&VMError{Type: e.Type, Err: errText, Line: e.Line, Column: e.Column})
	if err != nil {
		out = []byte(`{"type": "panic", "error": "marshalling VMError"}`)
	}
	return string(out)
}
