package compile

import (
	"reflect"
)

// ExtendData is used for the definition of the extended functions and variables
type ExtendData struct {
	Owner   *OwnerInfo
	Func    []ExtendFunc
	PreVar  []string
	Objects map[string]*Object
	Extern  bool // true if ignore not found identifiers object
}

type ExtendFunc struct {
	Name     string
	Func     any
	CanWrite bool
	AutoPars map[string]string
}

func (ext *ExtendData) MakeExtFunc() map[string]*Object {
	objects := make(map[string]*Object)
	for _, item := range ext.Func {
		obj := item.MakeObject()
		if obj != nil {
			objects[item.Name] = obj
		}
	}
	return objects
}

func (item *ExtendFunc) MakeObject() *Object {
	if item.ExtFuncInfo() != nil {
		return &Object{Type: ObjExtFunc, Value: item.ExtFuncInfo()}
	}
	return nil
}

func (item *ExtendFunc) ExtFuncInfo() *ExtFuncInfo {
	f := reflect.ValueOf(item.Func).Type()
	switch f.Kind() {
	case reflect.Func:
		data := &ExtFuncInfo{
			Name:     item.Name,
			Params:   make([]reflect.Type, f.NumIn()),
			Results:  make([]reflect.Type, f.NumOut()),
			Auto:     make([]string, f.NumIn()),
			Variadic: f.IsVariadic(),
			Func:     item.Func,
			CanWrite: item.CanWrite,
		}

		// populate Params, Auto, and Results
		for i := 0; i < f.NumIn(); i++ {
			if isauto, ok := item.AutoPars[f.In(i).String()]; ok {
				data.Auto[i] = isauto
			}
			data.Params[i] = f.In(i)
		}

		for i := 0; i < f.NumOut(); i++ {
			data.Results[i] = f.Out(i)
		}
		return data
	}
	return nil
}
