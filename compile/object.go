package compile

import (
	"reflect"
	"strconv"
	"strings"
)

// ObjectType Types of the compiled objects
type ObjectType int32

const (
	// ObjDefault is an default object.
	ObjDefault ObjectType = iota
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
	// ObjOwner is an owner object.
	ObjOwner
)

var ObjectType_name = map[int32]string{
	0: "Unknown",
	1: "Contract",
	2: "Func",
	3: "ExtFunc",
	4: "Var",
	5: "ExtVar",
	6: "Owner",
}

func (x ObjectType) String() string {
	s, ok := ObjectType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type (
	isObjInfoValue interface {
		isObjInfoValue()
		ObjectType() ObjectType
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
		Tails    map[string]FuncTail
		Variadic bool
		ID       uint32
		CanWrite bool // If the function can update DB
	}

	// FuncTail is storing param of FuncTail
	FuncTail struct {
		Name     string
		Params   []reflect.Type
		Offset   []int
		Variadic bool
		Decl     bool
	}

	// FuncTailCmd for CmdFuncTail
	FuncTailCmd struct {
		Count    int
		FuncTail FuncTail
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
		//object extend variable name or function name
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

func (o *OwnerInfo) ObjectType() ObjectType {
	return ObjOwner
}

func (c *ContractInfo) TxMap() map[string]*FieldInfo {
	if c == nil {
		return nil
	}
	var m = make(map[string]*FieldInfo)
	if c.Tx == nil {
		return m
	}
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
	var retLen int
	if ret.Type == ObjExtFunc {
		for _, result := range ret.GetExtFuncInfo().Results {
			if result.String() != "error" {
				retLen++
			}
		}
	}
	if ret.Type == ObjFunc {
		return len(ret.GetFuncInfo().Results)
	}
	return retLen
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

func (*CodeBlock) isObjInfoValue()                    {}
func (*ExtFuncInfo) isObjInfoValue()                  {}
func (*ObjInfoVariable) isObjInfoValue()              {}
func (*ObjInfoExtendVariable) isObjInfoValue()        {}
func (bc *CodeBlock) ObjectType() ObjectType          { return bc.Type }
func (*ExtFuncInfo) ObjectType() ObjectType           { return ObjExtFunc }
func (*ObjInfoVariable) ObjectType() ObjectType       { return ObjVar }
func (*ObjInfoExtendVariable) ObjectType() ObjectType { return ObjExtVar }

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

// AutoParamsCount returns the number of auto params
func (e *ExtFuncInfo) AutoParamsCount() int {
	count := 0
	for i := 0; i < len(e.Params); i++ {
		if len(e.Auto[i]) > 0 {
			count++
		}
	}
	return count
}

func (e *FuncInfo) ParamsCount() int {
	count := 0
	for i := 0; i < len(e.Params); i++ {
		count++
	}
	return count
}

func (e *FuncInfo) HasTails() bool {
	return e.Tails != nil
}

func (f FuncTail) IsParamEmpty(i int) bool {
	return f.Params[i] == reflect.TypeOf(nil)
}
