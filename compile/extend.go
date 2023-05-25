package compile

import (
	"reflect"
)

// ExtendData is used for the definition of the extended functions and variables
type ExtendData struct {
	Info   *OwnerInfo
	Func   []ExtendFunc
	PreVar []string
}

type ExtendFunc struct {
	Name     string
	Func     any
	CanWrite bool
	AutoPars map[string]string
}

func NewExtendData(info *OwnerInfo, fns []ExtendFunc, vars []string) *ExtendData {
	if info == nil {
		info = &OwnerInfo{StateID: 1}
	}
	return &ExtendData{Info: info, Func: fns, PreVar: vars}
}

func (ext *ExtendData) MakeExtFunc() map[string]*ObjInfo {
	objects := make(map[string]*ObjInfo)
	for _, item := range ext.Func {
		fobj := reflect.ValueOf(item.Func).Type()
		switch fobj.Kind() {
		case reflect.Func:
			data := &ExtFuncInfo{
				Name:     item.Name,
				Params:   make([]reflect.Type, fobj.NumIn()),
				Results:  make([]reflect.Type, fobj.NumOut()),
				Auto:     make([]string, fobj.NumIn()),
				Variadic: fobj.IsVariadic(),
				Func:     item.Func,
				CanWrite: item.CanWrite,
			}

			// populate Params, Auto, and Results
			for i := 0; i < fobj.NumIn(); i++ {
				if isauto, ok := item.AutoPars[fobj.In(i).String()]; ok {
					data.Auto[i] = isauto
				}
				data.Params[i] = fobj.In(i)
			}
			for i := 0; i < fobj.NumOut(); i++ {
				data.Results[i] = fobj.Out(i)
			}
			objects[item.Name] = NewObjInfo(ObjectType_ExtFunc, data)
		}
	}
	return objects
}

func (ext *ExtendData) MakePreVar() []string {
	return ext.PreVar
}
