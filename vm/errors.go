package vm

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
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

const (
	ContractError     = "Contract"
	JSONMarshallError = "JSONMarshall"
	ConversionError   = "Conversion"
	VMErr             = "VM"
)

var (
	errContractPars       = errors.New(`wrong contract parameters`)
	errWrongCountPars     = errors.New(`wrong count of parameters`)
	errUnsupportedType    = errors.New(`unsupported combination of types in the operator`)
	errMaxArrayIndex      = errors.New(`the index is out of range`)
	errMaxMapCount        = errors.New(`the maxumim length of map`)
	errSelfAssignment     = errors.New(`self assignment`)
	errIncorrectParameter = errors.New(`incorrect parameter of the condition function`)
)

// SetVMError sets error of VM
func SetVMError(eType string, eText any) error {
	errText := fmt.Sprintf(`%v`, eText)
	if len(errText) > MaxErrLen {
		errText = errText[:MaxErrLen] + `...`
	}
	out, err := json.Marshal(&VMError{Type: eType, Error: errText})
	if err != nil {
		log.WithFields(log.Fields{"type": JSONMarshallError, "error": err}).Error("marshalling VMError")
		out = []byte(`{"type": "panic", "error": "marshalling VMError"}`)
	}
	return fmt.Errorf(string(out))
}
