package compile

import (
	"reflect"
)

type IgnoreLevel int

const (
	IgnoreNone IgnoreLevel = iota
	//IgnoreIdent ignore not found identifiers object
	IgnoreIdent
)

// CompConfig is used for the definition of the extended functions and variables
type CompConfig struct {
	Owner     *OwnerInfo
	Func      []ExtendFunc
	PreVar    []string
	Objects   map[string]*Object
	IgnoreObj IgnoreLevel
}

type ExtendFunc struct {
	Name string
	// Func is the function to be called, it must be a function type
	Func     any
	CanWrite bool
	AutoPars map[string]string
}

func setDefault(conf *CompConfig) {
	if conf == nil {
		conf = &CompConfig{
			Objects: make(map[string]*Object),
			Owner:   &OwnerInfo{StateID: 1},
			PreVar:  make([]string, 0),
			Func:    make([]ExtendFunc, 0),
		}
	}
	if conf.Objects == nil {
		conf.Objects = make(map[string]*Object)
	}
	if conf.Owner == nil {
		conf.Owner = &OwnerInfo{StateID: 1}
	}
	if conf.PreVar == nil {
		conf.PreVar = make([]string, 0)
	}
	if conf.Func == nil {
		conf.Func = make([]ExtendFunc, 0)
	}
}

func (cfg *CompConfig) MakeConfig() *CompConfig {
	setDefault(cfg)
	return cfg
}

func (cfg *CompConfig) MakeExtFunc() map[string]*Object {
	objects := make(map[string]*Object)
	for _, item := range cfg.Func {
		obj := item.MakeObject()
		if obj != nil {
			objects[item.Name] = obj
		}
	}
	return objects
}

func (item *ExtendFunc) MakeObject() *Object {
	if v := item.ExtFuncInfo(); v != nil {
		return NewObject(ObjExtFunc, v)
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
