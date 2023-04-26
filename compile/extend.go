package compile

import "reflect"

// ExtendData is used for the definition of the extended functions and variables
type ExtendData struct {
	Objects    map[string]any
	AutoPars   map[string]string
	WriteFuncs map[string]struct{}
}

func NewExtendData(objects map[string]any, autoPars map[string]string, writeFuncs map[string]struct{}) *ExtendData {
	return &ExtendData{Objects: objects, AutoPars: autoPars, WriteFuncs: writeFuncs}
}

func (ext *ExtendData) MakeObj() map[string]*ObjInfo {
	objects := make(map[string]*ObjInfo)
	for key, item := range ext.Objects {
		fobj := reflect.ValueOf(item).Type()
		switch fobj.Kind() {
		case reflect.Func:
			_, canWrite := ext.WriteFuncs[key]
			data := &ExtFuncInfo{
				Name:     key,
				Params:   make([]reflect.Type, fobj.NumIn()),
				Results:  make([]reflect.Type, fobj.NumOut()),
				Auto:     make([]string, fobj.NumIn()),
				Variadic: fobj.IsVariadic(),
				Func:     item,
				CanWrite: canWrite,
			}

			// populate Params, Auto, and Results
			for i := 0; i < fobj.NumIn(); i++ {
				if isAuto, ok := ext.AutoPars[fobj.In(i).String()]; ok {
					data.Auto[i] = isAuto
				}
				data.Params[i] = fobj.In(i)
			}
			for i := 0; i < fobj.NumOut(); i++ {
				data.Results[i] = fobj.Out(i)
			}
			objects[key] = &ObjInfo{Type: ObjectType_ExtFunc, Value: data}
		}
	}
	return objects
}
