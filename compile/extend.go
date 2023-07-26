package compile

import (
	"reflect"
)

// ExtendData is used for the definition of the extended functions and variables
type ExtendData struct {
	Info    *OwnerInfo
	Func    []ExtendFunc
	PreVar  []string
	Objects map[string]*Object
	Extern  bool
}

type ExtendFunc struct {
	Name     string
	Func     any
	CanWrite bool
	AutoPars map[string]string
}

type ExtendBuilder struct {
	data ExtendData
}

func NewExtendBuilder() *ExtendBuilder {
	return &ExtendBuilder{}
}

func (b *ExtendBuilder) SetInfo(info *OwnerInfo) *ExtendBuilder {
	b.data.Info = info
	return b
}

func (b *ExtendBuilder) SetPreVar(vars []string) *ExtendBuilder {
	b.data.PreVar = vars
	return b
}

func (b *ExtendBuilder) SetFunc(fns []ExtendFunc) *ExtendBuilder {
	b.data.Func = fns
	return b
}

func (b *ExtendBuilder) SetExtern(extern bool) *ExtendBuilder {
	b.data.Extern = extern
	return b
}

func (b *ExtendBuilder) Build() *ExtendData {
	b.data.Objects = make(map[string]*Object)
	return &b.data
}

func (ext *ExtendData) MakeExtFunc() map[string]*Object {
	objects := make(map[string]*Object)
	for _, item := range ext.Func {
		obj := item.MakeObj()
		if obj != nil {
			objects[item.Name] = obj
		}
	}
	return objects
}

func (item *ExtendFunc) MakeObj() *Object {
	var obj *Object
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
		obj = NewObjInfo(ObjExtFunc, data)
	}

	return obj
}
