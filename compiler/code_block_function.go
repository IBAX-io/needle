package compiler

import "reflect"

// FunctionInfo contains the function information.
type FunctionInfo struct {
	Id      uint32
	Name    string
	Params  []Token
	Results []Token
	// tail function
	Tails    map[string]FuncTail
	Variadic bool
	CanWrite bool // If the function can update DB
}

// ParamsCount returns the number of parameters of the function.
func (e *FunctionInfo) ParamsCount() int {
	count := 0
	for i := 0; i < len(e.Params); i++ {
		count++
	}
	return count
}

// HasTails returns whether the function has tails function.
func (e *FunctionInfo) HasTails() bool {
	return e != nil && e.Tails != nil
}

// FuncTail contains the tail function information
type FuncTail struct {
	Name     string
	Params   []Token
	Offset   []int
	Variadic bool
}

// IsParamEmpty checks if the parameter at the given index is empty (nil).
func (f FuncTail) IsParamEmpty(i int) bool {
	return f.Params[i] == UNKNOWN
}

// ExtFuncInfo is the structure for the extended golang function.
type ExtFuncInfo struct {
	Name     string
	Params   []reflect.Type
	Results  []reflect.Type
	Auto     []string
	Variadic bool
	Func     any
	CanWrite bool // If the function can update DB
}

// Call executes the function with the provided parameters.
// It takes a slice of any type as an argument which represents the parameters to be passed to the function.
// It returns a slice of any type which represents the return values of the function.
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

// AutoParamsCount returns the number of auto params.
func (e *ExtFuncInfo) AutoParamsCount() int {
	count := 0
	for i := 0; i < len(e.Params); i++ {
		if len(e.Auto[i]) > 0 {
			count++
		}
	}
	return count
}
