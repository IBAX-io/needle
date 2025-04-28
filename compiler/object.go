package compiler

import (
	"strconv"
)

// ObjectType Types of the compiled objects
type ObjectType int32

const (
	// ObjDefault is an default object.
	ObjDefault ObjectType = iota
	// ObjCodeBlock is a code block object.
	ObjCodeBlock
	// ObjExtFunc is an extended build in function object.
	ObjExtFunc
	// ObjVariable is a variable object.
	ObjVariable
	// ObjExtVar is an extended build in variable.
	ObjExtVar
)

// ObjectTypeName maps the integer values of ObjectType to their string representations.
var ObjectTypeName = map[int32]string{
	0: "Default",
	1: "CodeBlock",
	2: "ExtFunc",
	3: "Variable",
	4: "ExtVar",
}

func (x ObjectType) String() string {
	s, ok := ObjectTypeName[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Object is the common object type that can be compiled.
type Object struct {
	Type ObjectType
	// Types that are assignable to Value:
	//
	// *CodeBlock
	// *ExtFuncInfo
	// *ObjInfoVariable
	// *ObjInfoExtendVariable
	Value isObjValue
}

// ObjInfoVariable is the structure for the local variable.
type ObjInfoVariable struct {
	// Name is the name of the local variable.
	Name string
	// Index is the position of the variable in the current block.
	Index int
}

// ObjInfoExtendVariable is the structure for the extended variable.
type ObjInfoExtendVariable struct {
	// Name is the name of the extended variable.
	Name string
}

// isObjValue is an interface that represents the value of an Object.
type isObjValue interface {
	isObjInfoValue()
}

func (*CodeBlock) isObjInfoValue()             {}
func (*ExtFuncInfo) isObjInfoValue()           {}
func (*ObjInfoVariable) isObjInfoValue()       {}
func (*ObjInfoExtendVariable) isObjInfoValue() {}

// NewObject creates a new Object with the given value and determines its ObjectType.
func NewObject(v isObjValue) *Object {
	var t ObjectType
	switch v.(type) {
	case *CodeBlock:
		t = ObjCodeBlock
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
	if obj.Type == ObjCodeBlock {
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
	if obj.Type == ObjCodeBlock {
		return len(obj.GetFunctionInfo().Results)
	}
	return retLen
}

// GetName returns the name of the object.
func (obj *Object) GetName() string {
	var name string
	switch obj.Type {
	case ObjCodeBlock:
		name = obj.GetCodeBlock().GetName()
	case ObjExtFunc:
		name = obj.GetExtFuncInfo().Name
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

	if obj.Type == ObjCodeBlock {
		return obj.GetFunctionInfo().Variadic
	}
	return false
}

func (obj *Object) IsCodeBlockContract() bool {
	if obj.Type == ObjCodeBlock && obj.GetCodeBlock().Type == CodeBlockContract {
		return true
	}
	return false
}

func (obj *Object) IsCodeBlockFunction() bool {
	if obj.Type == ObjCodeBlock && obj.GetCodeBlock().Type == CodeBlockFunction {
		return true
	}
	return false
}

// GetValue returns the value of the object.
func (obj *Object) GetValue() isObjValue {
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
	if obj.IsCodeBlockContract() {
		return obj.GetCodeBlock().GetContractInfo()
	}
	return nil
}

// GetFunctionInfo returns the FunctionInfo of the object if it exists.
func (obj *Object) GetFunctionInfo() *FunctionInfo {
	if obj.IsCodeBlockFunction() {
		return obj.GetCodeBlock().GetFunctionInfo()
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
