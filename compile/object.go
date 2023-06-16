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
	ObjUnknown ObjectType = iota
	// ObjContract is a contract object.
	ObjContract
	// ObjFunc is a function object. myfunc()
	ObjFunc
	// ObjExtFunc is an extended build in function object. $myfunc()
	ObjExtFunc
	// ObjVar is a variable. myvar
	ObjVar
	// ObjExtVar is an extended build in variable. $myvar
	ObjExtVar
)

var ObjectType_name = map[int32]string{
	0: "Unknown",
	1: "Contract",
	2: "Func",
	3: "ExtFunc",
	4: "Var",
	5: "ExtVar",
}

func (x ObjectType) String() string {
	s, ok := ObjectType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

var tailPrefix = "#"

type (
	isObjInfoValue interface {
		isObjInfoValue()
	}

	// OwnerInfo storing info about owner
	OwnerInfo struct {
		StateID  uint32
		Active   bool
		TableID  int64
		WalletID int64
		TokenID  int64
	}

	// ContractInfo contains the contract information
	ContractInfo struct {
		ID       uint32
		Name     string
		Owner    *OwnerInfo
		Used     map[string]bool // Called contracts
		Tx       *[]*FieldInfo
		Settings map[string]any
		CanWrite bool // If the function can update DB
	}

	// FieldInfo describes the field of the data structure
	FieldInfo struct {
		Name     string
		Type     reflect.Type
		Original Token
		Tags     string
	}

	// FuncInfo contains the function information
	FuncInfo struct {
		Name    string
		Params  []reflect.Type
		Results []reflect.Type
		//tail function
		Names    map[string]FuncName
		Variadic bool
		ID       uint32
		CanWrite bool // If the function can update DB
	}

	// FuncName is storing param of FuncName
	FuncName struct {
		Params   []reflect.Type
		Offset   []int
		Variadic bool
	}

	// FuncNameCmd for cmdFuncName
	FuncNameCmd struct {
		Name  string
		Count int
	}

	// ObjInfo is the common object type
	ObjInfo struct {
		Type ObjectType
		// Types that are valid to be assigned to Value:
		// 	*CodeBlock,included: *ContractInfo, *FuncInfo
		//	*ExtFuncInfo
		//	*ObjInfoVariable
		//	*ObjInfoExtendVariable
		Value isObjInfoValue
	}

	ObjInfoVariable struct {
		Name  string
		Index int
	}

	ObjInfoExtendVariable struct {
		//object extend variable name
		Name string
	}
	// ExtFuncInfo is the structure for the extended golang function
	ExtFuncInfo struct {
		Name     string
		Params   []reflect.Type
		Results  []reflect.Type
		Auto     []string
		Variadic bool
		Func     any
		CanWrite bool // If the function can update DB
	}

	// VarInfo contains the variable information.
	// including the location of the variable and the global variable
	VarInfo struct {
		Obj   *ObjInfo
		Owner *CodeBlock // is nil if the variable is global
	}

	// IndexInfo contains the information for SetIndex
	IndexInfo struct {
		VarOffset int
		Owner     *CodeBlock
		Extend    string
	}
)

func (c *ContractInfo) TxMap() map[string]*FieldInfo {
	if c == nil {
		return nil
	}
	var m = make(map[string]*FieldInfo)
	for _, n := range *c.Tx {
		m[n.Name] = n
	}
	return m
}

// ContainsTag returns whether the tag is contained in this field
func (fi *FieldInfo) ContainsTag(tag string) bool {
	return strings.Contains(fi.Tags, tag)
}

func NewObjInfo(t ObjectType, v isObjInfoValue) *ObjInfo {
	return &ObjInfo{Type: t, Value: v}
}

func (ret *ObjInfo) GetParamsLen() int {
	if ret.Type == ObjExtFunc {
		return len(ret.GetExtFuncInfo().Params)
	}
	if ret.Type == ObjFunc {
		return len(ret.GetFuncInfo().Params)
	}
	return 0
}
func (ret *ObjInfo) GetResultsLen() int {
	if ret.Type == ObjExtFunc {
		return len(ret.GetExtFuncInfo().Results)
	}
	if ret.Type == ObjFunc {
		return len(ret.GetFuncInfo().Results)
	}
	return 0
}

func (ret *ObjInfo) GetVariadic() bool {
	if ret.Type == ObjExtFunc {
		return ret.GetExtFuncInfo().Variadic
	}

	if ret.Type == ObjFunc {
		return ret.GetFuncInfo().Variadic
	}
	return false
}

func (*CodeBlock) isObjInfoValue()             {}
func (*ExtFuncInfo) isObjInfoValue()           {}
func (*ObjInfoVariable) isObjInfoValue()       {}
func (*ObjInfoExtendVariable) isObjInfoValue() {}

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

func (m *ObjInfo) GetContractInfo() *ContractInfo {
	if x, ok := m.GetCodeBlock().Info.(*ContractInfo); ok {
		return x
	}
	return nil
}

func (m *ObjInfo) GetFuncInfo() *FuncInfo {
	if x, ok := m.GetCodeBlock().Info.(*FuncInfo); ok {
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

func (m *ObjInfo) GetVariable() *ObjInfoVariable {
	if x, ok := m.GetValue().(*ObjInfoVariable); ok {
		return x
	}
	return nil
}

func (m *ObjInfo) GetExtendVariable() *ObjInfoExtendVariable {
	if x, ok := m.GetValue().(*ObjInfoExtendVariable); ok {
		return x
	}
	return nil
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

// AutoCount returns the number of auto params
func (e *ExtFuncInfo) AutoCount() int {
	count := 0
	for i := 0; i < len(e.Params); i++ {
		if len(e.Auto[i]) > 0 {
			count++
		}
	}
	return count
}
