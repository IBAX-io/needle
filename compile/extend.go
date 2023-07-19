package compile

import (
	"reflect"
)

// ExtendData is used for the definition of the extended functions and variables
type ExtendData struct {
	Info    *OwnerInfo
	Func    []ExtendFunc
	PreVar  []string
	Objects map[string]*ObjInfo
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

func (b *ExtendBuilder) Build() *ExtendData {
	b.data.Objects = make(map[string]*ObjInfo)
	return &b.data
}

func (ext *ExtendData) MakeExtFunc() map[string]*ObjInfo {
	objects := make(map[string]*ObjInfo)
	for _, item := range ext.Func {
		obj := item.MakeObj()
		if obj != nil {
			objects[item.Name] = obj
		}
	}
	return objects
}

func (item *ExtendFunc) MakeObj() *ObjInfo {
	var obj *ObjInfo
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
			//if fobj.Out(i).String() != "error" {
			//	if !SupportedType(fobj.Out(i)) {
			//		log.Panicf("unsupported output type %s for function %s", fobj.Out(i), item.Name)
			//	}
			//}
			data.Results[i] = fobj.Out(i)
		}
		obj = NewObjInfo(ObjExtFunc, data)
	}

	return obj
}
