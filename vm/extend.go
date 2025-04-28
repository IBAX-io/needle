package vm

import (
	"fmt"
	"reflect"
	"sync"
)

const (
	ExtendParentContract   = `parent_contract`   // parent contract name
	ExtendOriginalContract = `original_contract` // original contract name
	ExtendThisContract     = `this_contract`     // current contract name
	ExtendTimeLimit        = `time_limit`        // time limit for contract execution
	ExtendGenBlock         = `gen_block`         // true then we check the time limit
	ExtendTxCost           = `txcost`            // maximum cost limit of the transaction
	ExtendStack            = `stack`             // name of the contract stack

	ExtendSc = `sc` // implements the Stacker interface of struct
	ExtendRt = `rt` // runtime of the contract

	ExtendResult = `result` // result of the contract
)

// system variable cannot be changed through the contract
var sysVars = map[string]struct{}{
	ExtendParentContract:   {},
	ExtendOriginalContract: {},
	ExtendThisContract:     {},
	ExtendTimeLimit:        {},
	ExtendGenBlock:         {},
	ExtendTxCost:           {},
	ExtendStack:            {},
	ExtendRt:               {},
	ExtendSc:               {},
}

var sysExtendTypes = map[string]reflect.Type{
	ExtendParentContract:   reflect.TypeOf(""),
	ExtendOriginalContract: reflect.TypeOf(""),
	ExtendThisContract:     reflect.TypeOf(""),
	ExtendTimeLimit:        reflect.TypeOf(int64(0)),
	ExtendGenBlock:         reflect.TypeOf(true),
	ExtendTxCost:           reflect.TypeOf(int64(0)),
	ExtendStack:            reflect.TypeOf([]any{}),
	ExtendSc:               reflect.TypeOf((*Stacker)(nil)).Elem(),
	ExtendRt:               reflect.TypeOf(&Runtime{}),
	ExtendResult:           reflect.TypeOf((*any)(nil)).Elem(),
}

func GetSysVarsKeys() []string {
	keys := make([]string, 0, len(sysVars))
	for key := range sysVars {
		keys = append(keys, key)
	}
	return keys
}

type varInfo struct {
	value     any
	typ       reflect.Type
	isMutable bool
	isSystem  bool
}

type BaseExtendManager struct {
	mu      sync.RWMutex
	values  sync.Map
	memory  int64
	memVars map[string]int64
}

// NewExtendManager creates a new extend manager
func NewExtendManager() *BaseExtendManager {
	manager := &BaseExtendManager{
		memVars: make(map[string]int64),
	}
	for name, typ := range sysExtendTypes {
		defaultValue := reflect.Zero(typ).Interface()
		manager.registerVar(name, defaultValue, name == ExtendResult, true)
	}
	return manager
}

func (m *BaseExtendManager) registerVar(name string, value any, isMutable, isSystem bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, loaded := m.values.Load(name); loaded {
		return fmt.Errorf("variable %s already registered", name)
	}

	typ := reflect.TypeOf(value)

	m.values.Store(name, &varInfo{
		value:     value,
		typ:       typ,
		isMutable: isMutable,
		isSystem:  isSystem,
	})

	if value != nil {
		m.updateMemoryUsage(name, value)
	}

	return nil
}

// Set sets the value of the variable with the given key
func (m *BaseExtendManager) Set(key string, value any) error {
	m.mu.RLock()
	rawInfo, ok := m.values.Load(key)
	m.mu.RUnlock()
	if !ok {
		rawInfo = &varInfo{
			isMutable: true,
			typ:       reflect.TypeOf(value),
		}
	}
	info := rawInfo.(*varInfo)

	if !info.isMutable && key != ExtendResult {
		return fmt.Errorf("variable %s is immutable-", key)
	}

	valueType := reflect.TypeOf(value)

	if !valueType.AssignableTo(info.typ) {
		return fmt.Errorf("variable %s requires type %v, got %v", key, info.typ, valueType)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	info.value = value
	m.updateMemoryUsage(key, value)
	return nil
}

func (m *BaseExtendManager) Get(key string) (any, bool) {
	m.mu.RLock()
	rawInfo, ok := m.values.Load(key)
	m.mu.RUnlock()
	if !ok {
		return nil, false
	}
	return rawInfo.(*varInfo).value, true
}

func (m *BaseExtendManager) IsSysVar(key string) bool {
	rawInfo, ok := m.values.Load(key)
	if !ok {
		return false
	}
	return rawInfo.(*varInfo).isSystem
}

func (m *BaseExtendManager) GetImmutableVarKeys() []string {
	var immutableKeys []string
	m.values.Range(func(key, value any) bool {
		info := value.(*varInfo)
		if !info.isMutable {
			immutableKeys = append(immutableKeys, key.(string))
		}
		return true
	})

	return immutableKeys
}

func (m *BaseExtendManager) updateMemoryUsage(key string, newValue any) {
	oldMem := m.memVars[key]
	newMem := calcMem(newValue)

	m.memory += newMem - oldMem

	if newMem > 0 {
		m.memVars[key] = newMem
	} else {
		delete(m.memVars, key)
	}
}

type ExtendState struct {
	Values map[string]any
	Memory int64
}

func (m *BaseExtendManager) Snapshot() *ExtendState {
	m.mu.RLock()
	defer m.mu.RUnlock()

	state := &ExtendState{
		Values: make(map[string]any),
		Memory: m.memory,
	}

	m.values.Range(func(key, value any) bool {
		info := value.(*varInfo)
		state.Values[key.(string)] = info.value
		return true
	})

	return state
}

// Restore restores the state of the extend manager
func (m *BaseExtendManager) Restore(state *ExtendState) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.values = sync.Map{}
	m.memory = 0
	m.memVars = make(map[string]int64)

	m.memory = state.Memory
	for key, value := range state.Values {
		info, _ := m.values.Load(key)
		if info != nil {
			varInfo := info.(*varInfo)
			varInfo.value = value
			m.updateMemoryUsage(key, value)
		}
	}
}

func GetSysVarByKey[T any](m *BaseExtendManager, key string) T {
	if _, ok := sysVars[key]; !ok {
		var zero T
		return zero
	}
	value, _ := m.Get(key)

	if typedValue, ok := value.(T); ok {
		return typedValue
	}

	var zero T
	return zero
}

func GetSysVar[T any](rt *Runtime, key string) T {
	if typedValue, ok := rt.extend[key].(T); ok {
		return typedValue
	}

	var zero T
	return zero
}
