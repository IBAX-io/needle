package compiler

import (
	"reflect"
)

type IgnoreLevel int

const (
	IgnoreNone IgnoreLevel = iota
	// IgnoreIdent ignore not found identifiers object.
	// If it is not found, it will be treated as a contract object.
	IgnoreIdent
)

// Config is used for the definition of the extended functions and variables
type Config struct {
	Owner     *OwnerInfo
	Func      []ExtendFunc
	PreVar    []string
	IgnoreObj IgnoreLevel

	objects map[string]*Object
}

// EnsureDefault sets default values and returns the Config.
func (cfg *Config) EnsureDefault() *Config {
	if cfg.objects == nil {
		cfg.objects = make(map[string]*Object)
	}

	if cfg.Owner == nil {
		cfg.Owner = &OwnerInfo{StateId: 1}
	}
	if cfg.PreVar == nil {
		cfg.PreVar = make([]string, 0)
	}
	if cfg.Func == nil {
		cfg.Func = make([]ExtendFunc, 0)
	}
	return cfg
}

func (cfg *Config) SetObjects(src map[string]*Object) {
	cfg.objects = src
}

// MakeExtFunc returns a map of the object of the extended functions
func (cfg *Config) MakeExtFunc() map[string]*Object {
	objects := make(map[string]*Object)
	for _, item := range cfg.Func {
		obj := item.MakeObject()
		if obj != nil {
			objects[item.Name] = obj
		}
	}
	return objects
}

// ExtendFunc is used for the definition of the extended functions
type ExtendFunc struct {
	Name string
	// Func is the function to be called, it must be a function type
	Func     any
	CanWrite bool
	AutoPars map[string]string
}

// MakeObject returns an Object if the ExtendFuncInfo is not nil.
func (item *ExtendFunc) MakeObject() *Object {
	if v := item.ExtFuncInfo(); v != nil {
		return NewObject(v)
	}
	return nil
}

// ExtFuncInfo returns an ExtFuncInfo if the Func is a function type.
func (item *ExtendFunc) ExtFuncInfo() *ExtFuncInfo {
	f := reflect.ValueOf(item.Func).Type()
	switch f.Kind() {
	default:
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
