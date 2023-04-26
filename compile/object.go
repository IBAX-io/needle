package compile

import (
	"reflect"
	"strconv"
	"strings"
)

// ObjectType Types of the compiled objects
type ObjectType int32

const (
	// ObjUnknown is an unknown object.
	ObjectType_Unknown ObjectType = 0
	// ObjectType_Contract is a contract object.
	ObjectType_Contract ObjectType = 1
	// ObjectType_Func is a function object. myfunc()
	ObjectType_Func ObjectType = 2
	// ObjectType_ExtFunc is an extended build in function object. $myfunc()
	ObjectType_ExtFunc ObjectType = 3
	// ObjectType_Var is a variable. myvar
	ObjectType_Var ObjectType = 4
	// ObjectType_ExtVar is an extended build in variable. $myvar
	ObjectType_ExtVar ObjectType = 5
)

var ObjectType_name = map[int32]string{
	0: "Unknown",
	1: "Contract",
	2: "Func",
	3: "ExtFunc",
	4: "Var",
	5: "ExtVar",
}

var ObjectType_value = map[string]int32{
	"Unknown":  0,
	"Contract": 1,
	"Func":     2,
	"ExtFunc":  3,
	"Var":      4,
	"ExtVar":   5,
}

func (x ObjectType) String() string {
	s, ok := ObjectType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type isCodeBlockInfo interface {
	isCodeBlockInfo()
}

func (*OwnerInfo) isCodeBlockInfo()    {}
func (*ContractInfo) isCodeBlockInfo() {}
func (*FuncInfo) isCodeBlockInfo()     {}

func (bc *CodeBlock) GetInfo() isCodeBlockInfo {
	if bc != nil {
		return bc.Info
	}
	return nil
}

func (bc *CodeBlock) GetFuncInfo() *FuncInfo {
	if x, ok := bc.GetInfo().(*FuncInfo); ok {
		return x
	}
	return nil
}

func (bc *CodeBlock) GetContractInfo() *ContractInfo {
	if x, ok := bc.GetInfo().(*ContractInfo); ok {
		return x
	}
	return nil
}

func (bc *CodeBlock) GetOwnerInfo() *OwnerInfo {
	if x, ok := bc.GetInfo().(*OwnerInfo); ok {
		return x
	}
	return nil
}

// OwnerInfo storing info about owner
type OwnerInfo struct {
	StateID  uint32
	Active   bool
	TableID  int64
	WalletID int64
	TokenID  int64
}

// ContractInfo contains the contract information
type ContractInfo struct {
	ID       uint32
	Name     string
	Owner    *OwnerInfo
	Used     map[string]bool // Called contracts
	Tx       *[]*FieldInfo
	Settings map[string]any
	CanWrite bool // If the function can update DB
}

func (c *ContractInfo) TxMap() map[string]*FieldInfo {
	if c == nil {
		return nil
	}
	var m = make(map[string]*FieldInfo)
	for _, n := range *c.Tx {
		m[n.Name] = nil
	}
	return m
}
// FieldInfo describes the field of the data structure
type FieldInfo struct {
	Name     string
	Type     reflect.Type
	Original Token
	Tags     string
}

// ContainsTag returns whether the tag is contained in this field
func (fi *FieldInfo) ContainsTag(tag string) bool {
	return strings.Contains(fi.Tags, tag)
}
// FuncInfo contains the function information
type FuncInfo struct {
	Name    string
	Params  []reflect.Type
	Results []reflect.Type
	//tail function
	Names    *map[string]FuncName
	Variadic bool
	ID       uint32
	CanWrite bool // If the function can update DB
}

// FuncName is storing param of FuncName
type FuncName struct {
	Params   []reflect.Type
	Offset   []int
	Variadic bool
}

// FuncNameCmd for cmdFuncName
type FuncNameCmd struct {
	Name  string
	Count int
}
type isObjInfoValue interface {
	isObjInfoValue()
}
type ObjInfo_Variable struct {
	Name  string
	Index int
}
type ObjInfo_ExtendVariable struct {
	//object extend variable name
	Name string
}

func (*CodeBlock) isObjInfoValue()              {}
func (*ExtFuncInfo) isObjInfoValue()            {}
func (*ObjInfo_Variable) isObjInfoValue()       {}
func (*ObjInfo_ExtendVariable) isObjInfoValue() {}

func (m *ObjInfo) GetValue() isObjInfoValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ObjInfo) GetCodeBlock() *CodeBlock {
	if x, ok := m.GetValue().(*CodeBlock); ok {
		return x
	}
	return nil
}

func (m *ObjInfo) GetExtFuncInfo() *ExtFuncInfo {
	if x, ok := m.GetValue().(*ExtFuncInfo); ok {
		return x
	}
	return nil
}

func (m *ObjInfo) GetVariable() *ObjInfo_Variable {
	if x, ok := m.GetValue().(*ObjInfo_Variable); ok {
		return x
	}
	return nil
}

func (m *ObjInfo) GetExtendVariable() *ObjInfo_ExtendVariable {
	if x, ok := m.GetValue().(*ObjInfo_ExtendVariable); ok {
		return x
	}
	return nil
}

// ExtFuncInfo is the structure for the extended golang function
type ExtFuncInfo struct {
	Name     string
	Params   []reflect.Type
	Results  []reflect.Type
	Auto     []string
	Variadic bool
	Func     any
	CanWrite bool // If the function can update DB
}

func (e *ExtFuncInfo) Call(params []any) (ret []any) {
	foo := reflect.ValueOf(e.Func)
	var result []reflect.Value
	if e.Variadic {
		pars := make([]reflect.Value, len(e.Params))
		pars[len(pars)-1] = reflect.ValueOf(params[len(pars)-1:])
		result = foo.CallSlice(pars)
	} else {
		pars := make([]reflect.Value, len(params))
		for i := 0; i < len(params); i++ {
			pars[i] = reflect.ValueOf(params[i])
		}
		result = foo.Call(pars)
	}
	for _, value := range result {
		ret = append(ret, value.Interface())
	}
	return ret
}

// VarInfo contains the variable information
type VarInfo struct {
	Obj   *ObjInfo
	Owner *CodeBlock
}

// IndexInfo contains the information for SetIndex
type IndexInfo struct {
	VarOffset int
	Owner     *CodeBlock
	Extend    string
}
