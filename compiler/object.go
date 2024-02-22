package compiler

import (
	"strconv"
)

// ObjectType Types of the compiled objects
type ObjectType int32

const (
	// ObjDefault is an default object.
	ObjDefault ObjectType = iota
	// ObjContract is a contract object.
	ObjContract
	// ObjFunction is a function object.
	ObjFunction
	// ObjExtFunc is an extended build in function object.
	ObjExtFunc
	// ObjVariable is a variable object.
	ObjVariable
	// ObjExtVar is an extended build in variable.
	ObjExtVar
	// ObjOwner is an owner object.
	ObjOwner
)

// ObjectTypeName maps the integer values of ObjectType to their string representations.
var ObjectTypeName = map[int32]string{
	0: "Default",
	1: "Contract",
	2: "Func",
	3: "ExtFunc",
	4: "Var",
	5: "ExtVar",
	6: "Owner",
}

func (x ObjectType) String() string {
	s, ok := ObjectTypeName[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type (
	// Object is the common object type that can be compiled.
	Object struct {
		Type ObjectType
		// Types that are valid to be assigned to Value:
		// *CodeBlock
		// *ExtFuncInfo
		// *ObjInfoVariable
		// *ObjInfoExtendVariable
		Value isObjInfoValue
	}

	// isObjInfoValue is an interface that represents the value of an Object.
	isObjInfoValue interface {
		isObjInfoValue()
	}
)

// NewObject creates a new Object with the given value and determines its ObjectType.
func NewObject(v isObjInfoValue) *Object {
	var t ObjectType
	switch v.(type) {
	case *CodeBlock:
		t = v.(*CodeBlock).Type
	case *ExtFuncInfo:
		t = ObjExtFunc
	case *ObjInfoVariable:
		t = ObjVariable
	case *ObjInfoExtendVariable:
		t = ObjExtVar
	}
	return &Object{Type: t, Value: v}
}

// GetParamsLen returns the number of parameters of the object if it is a function or extended function.
func (obj *Object) GetParamsLen() int {
	if obj.Type == ObjExtFunc {
		return len(obj.GetExtFuncInfo().Params)
	}
	if obj.Type == ObjFunction {
		return len(obj.GetFunctionInfo().Params)
	}
	return 0
}

// GetResultsLen returns the number of results of the object if it is a function or extended function.
func (obj *Object) GetResultsLen() int {
	var retLen int
	if obj.Type == ObjExtFunc {
		for _, result := range obj.GetExtFuncInfo().Results {
			if result.String() != "error" {
				retLen++
			}
		}
	}
	if obj.Type == ObjFunction {
		return len(obj.GetFunctionInfo().Results)
	}
	return retLen
}

// GetName returns the name of the object.
func (obj *Object) GetName() string {
	var name string
	switch obj.Type {
	case ObjContract:
		name = obj.GetContractInfo().Name
	case ObjExtFunc:
		name = obj.GetExtFuncInfo().Name
	case ObjFunction:
		name = obj.GetFunctionInfo().Name
	case ObjVariable:
		name = obj.GetVariable().Name
	case ObjExtVar:
		name = obj.GetExtendVariable().Name
	}
	return name
}

// GetVariadic returns whether the object is a variadic function or extended function.
func (obj *Object) GetVariadic() bool {
	if obj.Type == ObjExtFunc {
		return obj.GetExtFuncInfo().Variadic
	}

	if obj.Type == ObjFunction {
		return obj.GetFunctionInfo().Variadic
	}
	return false
}

// GetValue returns the value of the object.
func (obj *Object) GetValue() isObjInfoValue {
	if obj != nil {
		return obj.Value
	}
	return nil
}

// GetCodeBlock returns the CodeBlock of the object if it exists.
func (obj *Object) GetCodeBlock() *CodeBlock {
	if x, ok := obj.GetValue().(*CodeBlock); ok {
		return x
	}
	return nil
}

// GetContractInfo returns the ContractInfo of the object if it exists.
func (obj *Object) GetContractInfo() *ContractInfo {
	cb := obj.GetCodeBlock()
	if cb != nil {
		if x, ok := cb.Info.(*ContractInfo); ok {
			return x
		}
	}
	return nil
}

// GetFunctionInfo returns the FunctionInfo of the object if it exists.
func (obj *Object) GetFunctionInfo() *FunctionInfo {
	cb := obj.GetCodeBlock()
	if cb != nil {
		if x, ok := cb.Info.(*FunctionInfo); ok {
			return x
		}
	}
	return nil
}

// GetExtFuncInfo returns the ExtFuncInfo of the object if it exists.
func (obj *Object) GetExtFuncInfo() *ExtFuncInfo {
	if x, ok := obj.GetValue().(*ExtFuncInfo); ok {
		return x
	}
	return nil
}

// GetVariable returns the ObjInfoVariable of the object if it exists.
func (obj *Object) GetVariable() *ObjInfoVariable {
	if x, ok := obj.GetValue().(*ObjInfoVariable); ok {
		return x
	}
	return nil
}

// GetExtendVariable returns the ObjInfoExtendVariable of the object if it exists.
func (obj *Object) GetExtendVariable() *ObjInfoExtendVariable {
	if x, ok := obj.GetValue().(*ObjInfoExtendVariable); ok {
		return x
	}
	return nil
}
