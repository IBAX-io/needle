package vm

import (
	"fmt"
	compile2 "github.com/IBAX-io/needle/compile"
	"reflect"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	statusNormal = iota
	statusReturn
	statusContinue
	statusBreak

	// Decimal is the constant string for decimal type
	Decimal = `decimal.Decimal`
	// Interface is the constant string for interface type
	Interface = `interface`
	File      = `*script.Map`

	brackets = `[]`

	maxArrayIndex = 1000000
	maxMapCount   = 100000
	maxCallDepth  = 1000
	memoryLimit   = 128 << 20 // 128 MB
	MaxErrLen     = 150
)

var sysVars = map[string]struct{}{
	sysVars_block:               {},
	sysVars_block_key_id:        {},
	sysVars_block_time:          {},
	sysVars_data:                {},
	sysVars_ecosystem_id:        {},
	sysVars_key_id:              {},
	sysVars_account_id:          {},
	sysVars_node_position:       {},
	sysVars_parent:              {},
	sysVars_original_contract:   {},
	sysVars_sc:                  {},
	sysVars_contract:            {},
	sysVars_stack:               {},
	sysVars_this_contract:       {},
	sysVars_time:                {},
	sysVars_type:                {},
	sysVars_txcost:              {},
	sysVars_txhash:              {},
	sysVars_guest_key:           {},
	sysVars_guest_account:       {},
	sysVars_black_hole_key:      {},
	sysVars_black_hole_account:  {},
	sysVars_white_hole_key:      {},
	sysVars_white_hole_account:  {},
	sysVars_gen_block:           {},
	sysVars_time_limit:          {},
	sysVars_pre_block_data_hash: {},
}

func isSysVar(name string) bool {
	if _, ok := sysVars[name]; ok || strings.HasPrefix(name, Extend_loop) {
		return true
	}
	return false
}

var (
	ErrMemoryLimit = errors.New("Memory limit exceeded")
	//ErrVMTimeLimit returns when the time limit exceeded
	ErrVMTimeLimit = errors.New(`time limit exceeded`)
)

// VMError represents error of VM
type VMError struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

type blockStack struct {
	Block  *compile2.CodeBlock
	Offset int
}

// ErrInfo stores info about current contract or function
type ErrInfo struct {
	Name string
	Line uint16
}

// Runtime is needed for the execution of the byte-code
type Runtime struct {
	stack     []any
	blocks    []*blockStack
	vars      []any
	extend    map[string]any
	vm        *VM
	cost      int64 //cost remaining
	err       error
	unwrap    bool
	limitName func(string) bool
	timeLimit bool
	callDepth uint16
	mem       int64
	memVars   map[any]int64
	errInfo   ErrInfo
}

// NewRuntime creates a new Runtime for the virtual machine
func NewRuntime(vm *VM, cost int64) *Runtime {
	return &Runtime{
		stack:     make([]any, 0, 1024),
		vm:        vm,
		cost:      cost,
		memVars:   make(map[any]int64),
		limitName: isSysVar,
	}
}

// SetCost sets the max cost of the execution.
func (rt *Runtime) SetCost(cost int64) {
	rt.cost = cost
}

func (rt *Runtime) SubCost(cost int64) error {
	if cost > 0 {
		rt.cost -= cost
	}
	if rt.cost < 0 {
		return fmt.Errorf("runtime cost limit overflow")
	}
	return nil
}

// Cost return the remain cost of the execution.
func (rt *Runtime) Cost() int64 {
	return rt.cost
}

// Run executes CodeBlock with the extended variables and functions
func (rt *Runtime) Run(block *compile2.CodeBlock, extend map[string]any) (ret []any, err error) {
	defer func() {
		if r := recover(); r != nil {
			//rt.vm.logger.WithFields(log.Fields{"type": PanicRecoveredError, "error_info": r, "stack": string(debug.Stack())}).Error("runtime panic error")
			err = fmt.Errorf(`runtime panic: %v`, r)
		}
	}()
	info := block.GetFuncInfo()
	if info == nil {
		return nil, fmt.Errorf("func info is nil")
	}
	rt.extend = extend
	var (
		genBlock bool
		timer    *time.Timer
	)
	if gen, ok := extend[Extend_gen_block]; ok {
		genBlock, _ = gen.(bool)
	}
	timeOver := func() {
		rt.timeLimit = false
	}
	if genBlock {
		timer = time.AfterFunc(time.Millisecond*time.Duration(extend[Extend_time_limit].(int64)), timeOver)
	}
	if _, err = rt.RunCode(block); err == nil {
		if rt.len() < len(info.Results) {
			var keyNames []string
			for i := 0; i < len(info.Results); i++ {
				keyNames = append(keyNames, info.Results[i].String())
			}
			err = fmt.Errorf("not enough arguments to return, need [%s]", strings.Join(keyNames, "|"))
		}
		off := rt.len() - len(info.Results)
		for i := 0; i < len(info.Results) && off >= 0; i++ {
			ret = append(ret, rt.stack[off+i])
		}
	}
	if genBlock {
		timer.Stop()
	}
	return
}

// RunCode executes CodeBlock
func (rt *Runtime) RunCode(block *compile2.CodeBlock) (status int, err error) {
	var cmd *compile2.ByteCode
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf(`runtime run code crashed: %v`, r)
		}
		if err != nil && !strings.HasPrefix(err.Error(), `{`) {
			var curContract, line string
			if block.IsParentContract() {
				stack := block.Parent.GetContractInfo()
				curContract = stack.Name
			}
			if stack, ok := rt.extend[Extend_stack].([]any); ok {
				curContract = stack[len(stack)-1].(string)
			}

			line = "]"
			if cmd != nil {
				line = fmt.Sprintf(":%d]", cmd.Line)
			}

			if len(rt.errInfo.Name) > 0 && rt.errInfo.Name != `ExecContract` {
				err = fmt.Errorf("%s [%s %s%s", err, rt.errInfo.Name, curContract, line)
				rt.errInfo.Name = ``
			} else {
				out := err.Error()
				if strings.HasSuffix(out, `]`) {
					prev := strings.LastIndexByte(out, ' ')
					if strings.HasPrefix(out[prev+1:], curContract+`:`) {
						out = out[:prev+1]
					} else {
						out = out[:len(out)-1] + ` `
					}
				} else {
					out += ` [`
				}
				err = fmt.Errorf(`%s%s%s`, out, curContract, line)
			}
		}
	}()
	rt.blocks = append(rt.blocks, &blockStack{Block: block, Offset: len(rt.vars)})
	var namemap map[string][]any
	if block.Type == compile2.ObjectType_Func && block.GetFuncInfo().Names != nil {
		if rt.peek() != nil {
			namemap = rt.peek().(map[string][]any)
		}
		rt.resetByIdx(rt.len() - 1)
	}
	start := rt.len()
	varoff := len(rt.vars)
	for vkey, vpar := range block.Vars {
		if err = rt.SubCost(1); err != nil {
			break
		}
		var value any
		if block.Type == compile2.ObjectType_Func && vkey < len(block.GetFuncInfo().Params) {
			value = rt.stack[start-len(block.GetFuncInfo().Params)+vkey]
		} else {
			value = reflect.New(vpar).Elem().Interface()
			if vpar == reflect.TypeOf(&compile2.Map{}) {
				value = compile2.NewMap()
			} else if vpar == reflect.TypeOf([]any{}) {
				value = make([]any, 0, len(rt.vars)+1)
			}
		}
		rt.addVar(value)
	}
	if err != nil {
		return
	}
	if namemap != nil {
		for key, item := range namemap {
			params := (*block.GetFuncInfo().Names)[key]
			for i, value := range item {
				if params.Variadic && i >= len(params.Params)-1 {
					off := varoff + params.Offset[len(params.Params)-1]
					rt.setVar(off, append(rt.vars[off].([]any), value))
				} else {
					rt.setVar(varoff+params.Offset[i], value)
				}
			}
		}
	}
	if block.Type == compile2.ObjectType_Func {
		start -= len(block.GetFuncInfo().Params)
	}
	var (
		assign []*compile2.VarInfo
		tmpInt int64
		tmpDec decimal.Decimal
	)
	top := make([]any, 8)
	labels := make([]int, 0)
	ctx := &instructionCtx{
		labels: labels,
		top:    top,
		assign: assign,
		tmpDec: tmpDec,
		tmpInt: tmpInt,
	}
main:
	//for ci := 0; ci < len(block.Code); ci++ {
	for ctx.ci = 0; ctx.ci < len(block.Code); ctx.ci++ {
		if err = rt.SubCost(1); err != nil {
			break
		}
		if rt.timeLimit {
			err = ErrVMTimeLimit
			break
		}

		if rt.mem > memoryLimit {
			err = ErrMemoryLimit
			break
		}

		//todo new
		cmd = block.Code[ctx.ci]
		ctx.size = rt.len()
		ctx.isContinue = false
		ctx.isBreak = false
		ctx.isLoop = false
		if ctx.size < int(cmd.Cmd>>8) {
			err = fmt.Errorf(`stack is empty`)
			break
		}
		for i := 1; i <= int(cmd.Cmd>>8); i++ {
			ctx.top[i-1] = rt.stack[ctx.size-i]
		}
		instruction, ok := instructionTable[cmd.Cmd]
		if ok {
			status, err = instruction(rt, cmd, ctx)
			if ctx.isContinue {
				continue
			}
			if ctx.isBreak {
				break
			}
			if ctx.isLoop {
				break main
			}
		} else {
			err = fmt.Errorf(`unknown command %d`, cmd.Cmd)
		}
		//todo end

		//todo add old there

		if err != nil {
			break
		}
		if status == statusReturn || status == statusContinue || status == statusBreak {
			break
		}
		//todo new
		if (cmd.Cmd >> 8) == 2 {
			rt.stack[ctx.size-2] = ctx.bin
			rt.resetByIdx(ctx.size - 1)
		}
		//todo end

		//if (cmd.Cmd >> 8) == 2 {
		//	rt.stack[size-2] = bin
		//	rt.resetByIdx(size - 1)
		//}
	}
	if err != nil {
		return
	}
	last := rt.popBlock()
	if status == statusReturn {
		if last.Block.Type == compile2.ObjectType_Func {
			lastResults := last.Block.GetFuncInfo().Results
			if len(lastResults) > rt.len() {
				var keyNames []string
				for i := 0; i < len(lastResults); i++ {
					keyNames = append(keyNames, lastResults[i].String())
				}
				err = fmt.Errorf("func '%s' not enough arguments to return, need [%s]", last.Block.GetFuncInfo().Name, strings.Join(keyNames, "|"))
				return
			}
			stackCpy := make([]any, rt.len())
			copy(stackCpy, rt.stack)
			var index int
			for count := len(lastResults); count > 0; count-- {
				val := stackCpy[len(stackCpy)-1-index]
				if val != nil && lastResults[count-1] != reflect.TypeOf(val) {
					err = fmt.Errorf("function '%s' return index[%d] (type %s) cannot be represented by the type %s", last.Block.GetFuncInfo().Name, count-1, reflect.TypeOf(val), lastResults[count-1])
					return
				}
				rt.stack[start] = rt.stack[rt.len()-count]
				start++
				index++
			}
			status = statusNormal
		} else {
			return
		}
	}

	rt.resetByIdx(start)
	return
}

func (rt *Runtime) callFunc(cmd uint16, obj *compile2.ObjInfo) (err error) {
	var (
		count, in int
	)
	if rt.callDepth >= maxCallDepth {
		return fmt.Errorf("max call depth of recursive call")
	}

	rt.callDepth++
	defer func() {
		rt.callDepth--
	}()

	size := rt.len()
	in = obj.GetParamsLen()
	if rt.unwrap && cmd == compile2.CmdCallVariadic && size > 1 &&
		reflect.TypeOf(rt.stack[size-2]).String() == `[]interface {}` {
		count = rt.getStack(size - 1).(int)
		arr := rt.getStack(size - 2).([]any)
		rt.resetByIdx(size - 2)
		for _, item := range arr {
			rt.push(item)
		}
		rt.push(count - 1 + len(arr))
		size = rt.len()
	}
	rt.unwrap = false
	if cmd == compile2.CmdCallVariadic {
		count = rt.getStack(size - 1).(int)
		size--
	} else {
		count = in
	}
	if obj.Type == compile2.ObjectType_Func {
		var imap map[string][]any
		finfo := obj.GetCodeBlock().GetFuncInfo()
		if finfo.Names != nil {
			if rt.getStack(size-1) != nil {
				imap = rt.getStack(size - 1).(map[string][]any)
			}
			rt.resetByIdx(size - 1)
			size = rt.len()
		}
		if cmd == compile2.CmdCallVariadic {
			parcount := count + 1 - in
			if parcount < 0 {
				log.WithFields(log.Fields{"type": VMErr}).Error(errWrongCountPars)
				return errWrongCountPars
			}
			pars := make([]any, parcount)
			shift := size - parcount
			for i := parcount; i > 0; i-- {
				pars[i-1] = rt.stack[size+i-parcount-1]
			}
			rt.resetByIdx(shift)
			rt.push(pars)
		}
		if rt.len() < len(finfo.Params) {
			log.WithFields(log.Fields{"type": VMErr}).Error(errWrongCountPars)
			return errWrongCountPars
		}
		for i, v := range finfo.Params {
			switch v.Kind() {
			case reflect.String, reflect.Int64:
				offset := rt.len() - in + i
				if v.Kind() == reflect.Int64 {
					rv := reflect.ValueOf(rt.stack[offset])
					switch rv.Kind() {
					case reflect.Float64:
						val, _ := ValueToInt(rt.stack[offset])
						rt.stack[offset] = val
					}
				}
				if reflect.TypeOf(rt.stack[offset]) != v {
					log.WithFields(log.Fields{"type": VMErr}).Error(fmt.Sprintf(eTypeParam, i+1))
					return fmt.Errorf(eTypeParam, i+1)
				}
			}
		}
		if finfo.Names != nil {
			rt.push(imap)
		}
		_, err = rt.RunCode(obj.GetCodeBlock())
		return
	}

	var (
		stack  Stacker
		ok     bool
		result []reflect.Value
		limit  = 0
		finfo  = obj.GetExtFuncInfo()
		foo    = reflect.ValueOf(finfo.Func)
		pars   = make([]reflect.Value, in)
	)
	if stack, ok = rt.extend[Extend_sc].(Stacker); ok {
		if err := stack.AppendStack(finfo.Name); err != nil {
			return err
		}
	}
	rt.extend[Extend_rt] = rt
	auto := 0
	for k := 0; k < in; k++ {
		if len(finfo.Auto[k]) > 0 {
			auto++
		}
	}
	shift := size - count + auto
	if finfo.Variadic {
		shift = size - count
		count += auto
		limit = count - in + 1
	}
	i := count
	for ; i > limit; i-- {
		if len(finfo.Auto[count-i]) > 0 {
			pars[count-i] = reflect.ValueOf(rt.extend[finfo.Auto[count-i]])
			auto--
		} else {
			pars[count-i] = reflect.ValueOf(rt.stack[size-i+auto])
		}
		if !pars[count-i].IsValid() {
			pars[count-i] = reflect.Zero(reflect.TypeOf(``))
		}
	}
	if i > 0 && size-i >= 0 {
		pars[in-1] = reflect.ValueOf(rt.stack[size-i : size])
	} else {
		if !pars[in-1].IsValid() {
			pars[in-1] = reflect.Zero(finfo.Params[in-1])
		}
	}
	if finfo.Name == `ExecContract` && (pars[2].Kind() != reflect.String || !pars[3].IsValid()) {
		return fmt.Errorf(`unknown function %v`, pars[1])
	}
	if finfo.Variadic {
		result = foo.CallSlice(pars)
	} else {
		result = foo.Call(pars)
	}
	if shift < 0 {
		shift = 0
	}
	rt.resetByIdx(shift)
	if stack != nil {
		stack.PopStack(finfo.Name)
	}

	for i, ret := range result {
		// first return value of every extend function that makes queries to DB is cost
		if _, ok := rt.vm.FuncCallsDB[finfo.Name]; ok && i == 0 {
			if err = rt.SubCost(ret.Int()); err != nil {
				return
			}
			continue
		}
		if finfo.Results[i].String() == `error` {
			if ret.Interface() != nil {
				rt.errInfo = ErrInfo{Name: finfo.Name}
				return ret.Interface().(error)
			}
		} else {
			rt.push(ret.Interface())
		}
	}
	return
}

func (rt *Runtime) extendFunc(name string) error {
	var (
		ok bool
		f  any
	)
	if f, ok = rt.extend[name]; !ok || reflect.ValueOf(f).Kind() != reflect.Func {
		return fmt.Errorf(`unknown function %s`, name)
	}
	size := rt.len()
	foo := reflect.ValueOf(f)

	count := foo.Type().NumIn()
	pars := make([]reflect.Value, count)
	for i := count; i > 0; i-- {
		pars[count-i] = reflect.ValueOf(rt.stack[size-i])
	}
	result := foo.Call(pars)

	rt.resetByIdx(size - count)
	for i, iret := range result {
		if foo.Type().Out(i).String() == `error` {
			if iret.Interface() != nil {
				return iret.Interface().(error)
			}
		} else {
			rt.push(iret.Interface())
		}
	}
	return nil
}

func (rt *Runtime) setExtendVar(k string, v any) {
	rt.extend[k] = v
	rt.recalcMemExtendVar(k)
}

func (rt *Runtime) recalcMemExtendVar(k string) {
	mem := calcMem(rt.extend[k])
	rt.mem += mem - rt.memVars[k]
	rt.memVars[k] = mem
}

func (rt *Runtime) addVar(v any) {
	rt.vars = append(rt.vars, v)
	mem := calcMem(v)
	rt.memVars[len(rt.vars)-1] = mem
	rt.mem += mem
}

func (rt *Runtime) setVar(k int, v any) {
	rt.vars[k] = v
	rt.recalculateMemVar(k)
}

func (rt *Runtime) recalculateMemVar(k int) {
	mem := calcMem(rt.vars[k])
	rt.mem += mem - rt.memVars[k]
	rt.memVars[k] = mem
}

func (rt *Runtime) getResultValue(item compile2.MapItem) (value any, err error) {
	switch item.Type {
	case compile2.MapConst:
		value = item.Value
	case compile2.MapExtend:
		var ok bool
		value, ok = rt.extend[item.Value.(string)]
		if !ok {
			err = fmt.Errorf(`unknown extend identifier %s`, item.Value)
		}
	case compile2.MapVar:
		ivar := item.Value.(*compile2.VarInfo)
		var i int
		for i = len(rt.blocks) - 1; i >= 0; i-- {
			if ivar.Owner == rt.blocks[i].Block {
				value = rt.vars[rt.blocks[i].Offset+ivar.Obj.GetVariable().Index]
				break
			}
		}
		if i < 0 {
			err = fmt.Errorf(eWrongVar, ivar.Obj.Value)
		}
	case compile2.MapMap:
		value, err = rt.getResultMap(item.Value.(*compile2.Map))
	case compile2.MapArray:
		value, err = rt.getResultArray(item.Value.([]compile2.MapItem))
	}
	return
}

func (rt *Runtime) getResultArray(cmd []compile2.MapItem) ([]any, error) {
	initArr := make([]any, 0)
	for _, val := range cmd {
		value, err := rt.getResultValue(val)
		if err != nil {
			return nil, err
		}
		initArr = append(initArr, value)
	}
	return initArr, nil
}

func (rt *Runtime) getResultMap(cmd *compile2.Map) (*compile2.Map, error) {
	initMap := compile2.NewMap()
	for _, key := range cmd.Keys() {
		val, _ := cmd.Get(key)
		value, err := rt.getResultValue(val.(compile2.MapItem))
		if err != nil {
			return nil, err
		}
		initMap.Set(key, value)
	}
	return initMap, nil
}
