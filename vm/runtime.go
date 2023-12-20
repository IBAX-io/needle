package vm

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/IBAX-io/needle/compile"
	"github.com/pkg/errors"
)

// Runtime is needed for the execution of the byte-code
type Runtime struct {
	stack      *Stack
	blocks     []*blockStack
	vars       []any
	extend     map[string]any
	vm         *VM
	costRemain int64
	costLimit  int64
	used       map[string]struct{}
	unwrap     bool
	timeLimit  bool
	callDepth  uint16
	mem        int64
	memVars    map[any]int64
	errInfo    ExtFuncErr
}

// NewRuntime creates a new Runtime for the virtual machine
func NewRuntime(vm *VM, extend map[string]any, cost int64) *Runtime {
	return &Runtime{
		stack:      newStack(),
		vm:         vm,
		costRemain: cost,
		costLimit:  cost,
		extend:     extend,
		memVars:    make(map[any]int64),
		used:       make(map[string]struct{}),
	}
}

// Stacker represents interface for working with call stack
type Stacker interface {
	AppendStack(fn string) error
	PopStack(fn string)
}

// SetCost sets the max cost of the execution.
func (rt *Runtime) SetCost(cost int64) {
	rt.costRemain = cost
}

func (rt *Runtime) SubCost(cost int64) error {
	if cost > 0 {
		rt.costRemain -= cost
	}
	if rt.costRemain < 0 {
		return fmt.Errorf("runtime cost limit overflow")
	}
	return nil
}

// CostRemain return the remain cost of the execution.
func (rt *Runtime) CostRemain() int64 {
	return rt.costRemain
}

func (rt *Runtime) CostUsed() int64 {
	return rt.costLimit - rt.costRemain
}

// Run executes CodeBlock with the extended variables and functions
func (rt *Runtime) Run(block *compile.CodeBlock) (ret []any, err error) {
	defer func() {
		if r := recover(); r != nil {
			//rt.vm.logger.WithFields(log.Fields{"type": PanicRecoveredError, "error_info": r, "stack": string(debug.Stack())}).Error("runtime panic error")
			err = fmt.Errorf("runtime panic: %v", r)
		}
	}()
	info := block.GetFuncInfo()
	if info == nil {
		return nil, fmt.Errorf("the block is not a function")
	}
	var (
		genBlock bool
		timer    *time.Timer
	)
	genBlock = rt.loadExtendBy(ExtendGenBlock).genBlock
	timeOver := func() {
		rt.timeLimit = false
	}
	if genBlock {
		timer = time.AfterFunc(time.Millisecond*time.Duration(rt.loadExtendBy(ExtendTimeLimit).timeLimit), timeOver)
	}
	defer func() {
		if genBlock {
			timer.Stop()
		}
	}()

	if _, err = rt.RunCode(block); err != nil {
		return
	}
	if rt.stack.size() < len(info.Results) {
		var keyNames []string
		for i := 0; i < len(info.Results); i++ {
			keyNames = append(keyNames, info.Results[i].String())
		}
		err = fmt.Errorf("func '%s' not enough arguments to return, need %s", info.Name, strings.Join(keyNames, "|"))
	}
	off := rt.stack.size() - len(info.Results)
	for i := 0; i < len(info.Results) && off >= 0; i++ {
		ret = append(ret, rt.stack.get(off+i))
	}
	return
}

// RunCode executes CodeBlock
func (rt *Runtime) RunCode(block *compile.CodeBlock) (status int, err error) {
	var cmd *compile.ByteCode
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf(`runcode panic: %v`, r)
		}

		if err != nil {
			if !errors.As(err, &VMError{}) {
				var name, line string
				if block.Parent != nil && block.Parent.Type == compile.ObjFunc {
					name = block.Parent.GetFuncInfo().Name
				}
				if block.Type == compile.ObjFunc {
					name = block.GetFuncInfo().Name
				}
				if block.IsParentContract() {
					stack := block.Parent.GetContractInfo()
					name = stack.Name
				}

				if stack, ok := rt.extend[ExtendStack].([]any); ok {
					name = stack[len(stack)-1].(string)
				}

				line = "]"
				if cmd != nil && cmd.Lexeme != nil {
					line = fmt.Sprintf(":%d", cmd.Lexeme.Line)
					if cmd.Lexeme.Column != 0 {
						line += fmt.Sprintf(":%d", cmd.Lexeme.Column)
					}
					line += "]"
				}

				if len(rt.errInfo.Name) > 0 && rt.errInfo.Name != `ExecContract` {
					err = fmt.Errorf("%s [%s %s%s", err, rt.errInfo.Name, name, line)
					rt.errInfo.Name = ``
				} else {
					out := err.Error()
					if strings.HasSuffix(out, `]`) {
						prev := strings.LastIndexByte(out, ' ')
						if strings.HasPrefix(out[prev+1:], name+`:`) {
							out = out[:prev+1]
						} else {
							out = out[:len(out)-1] + ` `
						}
					} else {
						out += ` [`
					}
					err = fmt.Errorf(`%s%s%s`, out, name, line)
				}
			}
		}
	}()
	rt.pushBlock(&blockStack{Block: block, Offset: len(rt.vars)})
	var names map[string][]any
	if block.Type == compile.ObjFunc && block.GetFuncInfo().HasTails() {
		ret := rt.stack.pop()
		if ret != nil {
			names, _ = ret.(map[string][]any)
		}
	}
	start := rt.stack.size()

	if err = rt.SubCost(int64(len(block.Vars))); err != nil {
		return
	}
	rt.addVarBy(block)
	err = rt.setVarBy(block, names)
	if err != nil {
		return
	}
	if block.Type == compile.ObjFunc {
		start -= len(block.GetFuncInfo().Params)
	}
	if start < 0 {
		err = fmt.Errorf("not enough arguments in call to")
		return
	}

	ctx := newInstructionCtx()
	for ctx.ci = 0; ctx.ci < len(block.Code); ctx.ci++ {
		if err = rt.SubCost(1); err != nil {
			break
		}
		if rt.timeLimit {
			err = ErrVMTimeLimit
			break
		}
		if rt.mem > MemoryLimit {
			err = ErrMemoryLimit
			break
		}
		cmd = block.Code[ctx.ci]
		instruction, ok := instructionTable[cmd.Cmd]
		if !ok {
			err = fmt.Errorf(`unknown command '%s'`, cmd.Cmd)
			break
		}
		status, err = instruction(rt, cmd, ctx)
		if err != nil {
			break
		}
		if block.Type != compile.ObjDefault && (status == statusContinue || status == statusBreak) {
			err = fmt.Errorf(`%s is outside of loop block`, cmd.Cmd)
			break
		}
		if status == statusReturn || status == statusContinue || status == statusBreak {
			break
		}
	}
	if err != nil {
		return
	}
	last := rt.popBlock()
	if status == statusReturn {
		if last.Block.Type != compile.ObjFunc {
			return
		}
		funcInfo := last.Block.GetFuncInfo()
		refRes := funcInfo.Results
		local := rt.stack.peekFromTo(start, rt.stack.size())
		if len(refRes) > len(local) {
			err = fmt.Errorf("not enough arguments to return")
			return
		}
		rt.stack.resetByIdx(start)
		ret := local[len(local)-len(refRes):]
		for i, v := range ret {
			if refRes[i] != reflect.TypeOf(v) {
				err = fmt.Errorf("func '%s' return index[%d] %v (type %v) cannot be represented by the type %s", last.Block.GetFuncInfo().Name, i, v, reflect.TypeOf(v), refRes[i])
				return
			}
			rt.stack.push(v)
			start++
		}
		status = statusNormal
	}
	rt.stack.resetByIdx(start)
	return
}

func (rt *Runtime) callFunc(obj *compile.Object) (err error) {
	var (
		count, in int
	)
	variadic := obj.GetVariadic()
	in = obj.GetParamsLen()
	if rt.unwrap && variadic && rt.stack.size() > 1 &&
		reflect.TypeOf(rt.stack.get(rt.stack.size()-2)).String() == `[]interface {}` {
		count = rt.stack.pop().(int)
		arr := rt.stack.pop().([]any)
		rt.stack.pushN(arr)
		rt.stack.push(count - 1 + len(arr))
	}
	rt.unwrap = false
	count = in
	if variadic {
		count = rt.stack.pop().(int)
	}

	var (
		result []reflect.Value
		limit  int
		finfo  = obj.GetExtFuncInfo()
		foo    = reflect.ValueOf(finfo.Func)
		pars   = make([]reflect.Value, in)
	)
	stack, ok := rt.extend[ExtendSc].(Stacker)
	if ok {
		if err := stack.AppendStack(finfo.Name); err != nil {
			return err
		}
	}
	rt.extend[ExtendRt] = rt
	auto := finfo.AutoParamsCount()
	size := rt.stack.size()
	shift := size - count + auto
	if finfo.Variadic {
		shift = size - count
		count += auto
		limit = count - in + 1
	}
	i := count
	for ; i > limit; i-- {
		index := count - i
		if len(finfo.Auto[index]) > 0 {
			value, ok := rt.extend[finfo.Auto[index]]
			if !ok {
				return fmt.Errorf("func %q auto param %q not found", finfo.Name, finfo.Auto[index])
			}
			pars[index] = reflect.ValueOf(value)
			auto--
		} else {
			pars[index] = reflect.ValueOf(rt.stack.get(size - i + auto))
		}
		if !pars[index].IsValid() {
			pars[index] = reflect.Zero(finfo.Params[index])
		}
	}
	if i > 0 {
		if size-i >= 0 {
			pars[in-1] = reflect.ValueOf(rt.stack.peekFromTo(size-i, size))
		} else {
			if !pars[in-1].IsValid() {
				pars[in-1] = reflect.Zero(finfo.Params[in-1])
			}
		}
	}
	if finfo.Variadic {
		if i == 0 {
			pars = []reflect.Value{reflect.Zero(finfo.Params[in-1])}
		}
		result = foo.CallSlice(pars)
	} else {
		result = foo.Call(pars)
	}
	if shift < 0 {
		shift = 0
	}
	rt.stack.resetByIdx(shift)
	if stack != nil {
		stack.PopStack(finfo.Name)
	}

	for i, ret := range result {
		// first return value of every extend function that makes queries to DB is cost
		if _, ok := rt.vm.FuncCallsDB[finfo.Name]; ok && i == 0 {
			if !ret.CanInt() {
				err = fmt.Errorf("invalid type of first return parameter")
				return
			}
			if err = rt.SubCost(ret.Int()); err != nil {
				return
			}
			continue
		}
		if finfo.Results[i].String() == `error` {
			if ret.Interface() != nil {
				rt.errInfo = ExtFuncErr{Name: finfo.Name, Value: ret.Interface()}
				return rt.errInfo
			}
		} else {
			rt.stack.push(ret.Interface())
		}
	}
	return
}

func (rt *Runtime) callObjFunc(obj *compile.Object) error {
	var (
		count    int
		imap     map[string][]any
		finfo    = obj.GetFuncInfo()
		in       = len(finfo.Params)
		variadic = finfo.Variadic
	)
	if rt.unwrap && variadic && rt.stack.size() > 1 &&
		reflect.TypeOf(rt.stack.get(rt.stack.size()-2)).String() == `[]interface {}` {
		count = rt.stack.pop().(int)
		arr := rt.stack.pop().([]any)
		rt.stack.pushN(arr)
		rt.stack.push(count - 1 + len(arr))
	}
	rt.unwrap = false
	count = in
	if variadic {
		count = rt.stack.pop().(int)
	}

	if finfo.HasTails() {
		imap, _ = rt.stack.pop().(map[string][]any)
	}
	if finfo.Variadic {
		parcount := count + 1 - in
		if parcount < 0 {
			return errWrongCountPars
		}
		pars := make([]any, parcount)
		shift := rt.stack.size() - parcount
		for i := parcount; i > 0; i-- {
			pars[i-1] = rt.stack.get(shift + i - 1)
		}
		rt.stack.resetByIdx(shift)
		rt.stack.push(pars)
	}
	if rt.stack.size() < len(finfo.Params) {
		return fmt.Errorf("func '%s' wrong number of parameters, expected %d, got %d", finfo.Name, len(finfo.Params), rt.stack.size())
	}
	if len(finfo.Params) == 0 {
		rt.stack.pushN(nil)
	}
	for i, v := range finfo.Params {
		offset := rt.stack.size() - in + i
		stack := rt.stack.get(offset)
		if v.Kind() == reflect.Int64 && reflect.TypeOf(stack).Kind() == reflect.Float64 {
			val, _ := ValueToInt(stack)
			rt.stack.set(offset, val)
		}
		if reflect.TypeOf(stack) != v {
			return fmt.Errorf("func '%s' param: cannot use (type %T) as the type %s", finfo.Name, stack, v)
		}
	}
	if finfo.HasTails() {
		rt.stack.push(imap)
	}
	_, err := rt.RunCode(obj.GetCodeBlock())
	return err
}

func (rt *Runtime) extendFunc(name string) error {
	f, ok := rt.extend[name]
	foo := reflect.ValueOf(f)
	if !ok || foo.Kind() != reflect.Func {
		return fmt.Errorf(`unknown extend function $%s`, name)
	}
	variadic := foo.Type().IsVariadic()
	count := foo.Type().NumIn()
	last := count
	stack := rt.stack.popN(count)
	size := len(stack)
	if variadic {
		last--
		if size < last {
			return fmt.Errorf("parameter expected at least %d, got %d", last, size)
		}
	} else {
		if size != count {
			return fmt.Errorf("parameter expected %d, got %d", count, size)
		}
	}
	var result []reflect.Value
	pars := make([]reflect.Value, count)
	var lastType reflect.Type

	for i, vs := range stack {
		ftyp := foo.Type()
		v := reflect.ValueOf(vs)
		var ityp reflect.Type
		if i < last {
			ityp = ftyp.In(i)
		} else {
			ityp = ftyp.In(last).Elem()
			lastType = ityp
		}
		if !v.IsValid() {
			stack[i] = reflect.Zero(ityp)
			continue
		}
		vtyp := v.Type()
		if !vtyp.AssignableTo(ityp) {
			k := vtyp.Kind()
			if (k == reflect.Ptr || k == reflect.Interface) && !v.IsNil() && vtyp.Elem().AssignableTo(ityp) {
				stack[i] = v.Elem()
				continue
			}
			if reflect.PtrTo(vtyp).AssignableTo(ityp) && v.CanAddr() {
				stack[i] = v.Addr()
				continue
			}
			return fmt.Errorf("can't with %s as argument %d, need %s", vtyp, i+1, ityp)
		}
	}
	if foo.Type().IsVariadic() {
		var arr = make([]any, 0)
		for _, v := range stack[last:] {
			arr = append(arr, v)
		}
		if len(stack) <= last {
			stack = append(stack, reflect.Value{})
		}
		for i := 0; i < count; i++ {
			if i >= last {
				if lastType == reflect.TypeOf((*any)(nil)).Elem() {
					pars[i] = reflect.ValueOf(arr)
				} else {
					lastSlice := reflect.MakeSlice(foo.Type().In(last), len(arr), len(arr))
					for y, v := range arr {
						lastSlice.Index(y).Set(reflect.ValueOf(v))
					}
					pars[i] = lastSlice
				}
				continue
			}
			pars[i] = reflect.ValueOf(stack[i])
		}
		result = foo.CallSlice(pars)
	} else {
		for i := 0; i < count; i++ {
			pars[i] = reflect.ValueOf(stack[i])
		}
		result = foo.Call(pars)
	}

	for i, ret := range result {
		if foo.Type().Out(i).String() == `error` {
			if ret.Interface() != nil {
				return ret.Interface().(error)
			}
		} else {
			rt.stack.push(ret.Interface())
		}
	}
	return nil
}

func (rt *Runtime) setExtendVar(k string, v any) {
	rt.extend[k] = v
	rt.recalculateMemExtendVar(k)
}

func (rt *Runtime) recalculateMemExtendVar(k string) {
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

func (rt *Runtime) getResultValue(item compile.MapItem) (value any, err error) {
	switch item.Type {
	case compile.MapConst:
		value = item.Value
	case compile.MapExtend:
		var ok bool
		value, ok = rt.extend[item.Value.(string)]
		if !ok {
			err = fmt.Errorf(`unknown extend identifier '$%s'`, item.Value)
		}
	case compile.MapVar:
		ivar := item.Value.(*compile.VarInfo)
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
	case compile.MapMap:
		value, err = rt.getResultMap(item.Value.(*compile.Map))
	case compile.MapArray:
		value, err = rt.getResultArray(item.Value.([]compile.MapItem))
	}
	return
}

func (rt *Runtime) getResultArray(cmd []compile.MapItem) ([]any, error) {
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

func (rt *Runtime) getResultMap(cmd *compile.Map) (*compile.Map, error) {
	initMap := compile.NewMap()
	for _, key := range cmd.Keys() {
		val, _ := cmd.Get(key)
		value, err := rt.getResultValue(val.(compile.MapItem))
		if err != nil {
			return nil, err
		}
		initMap.Set(key, value)
	}
	return initMap, nil
}

func (rt *Runtime) addVarBy(block *compile.CodeBlock) {
	for key, par := range block.Vars {
		var value any
		if block.Type == compile.ObjFunc && key < len(block.GetFuncInfo().Params) {
			value = rt.stack.getAndDel(rt.stack.size() - len(block.GetFuncInfo().Params) + key)
		} else {
			value = reflect.New(par).Elem().Interface()
			if par == reflect.TypeOf(&compile.Map{}) {
				value = compile.NewMap()
			} else if par == reflect.TypeOf([]any{}) {
				value = make([]any, 0, len(rt.vars)+1)
			}
		}
		rt.addVar(value)
	}
}

func (rt *Runtime) setVarBy(block *compile.CodeBlock, names map[string][]any) error {
	varoff := len(rt.vars) - len(block.Vars)
	for key, item := range names {
		params := block.GetFuncInfo().Tails[key]
		for i, value := range item {
			var ind int
			if params.Variadic && i >= len(params.Params)-1 {
				ind = varoff + params.Offset[len(params.Params)-1]
				value = append(rt.vars[ind].([]any), value)
			} else {
				ind = varoff + params.Offset[i]
				refx := reflect.TypeOf(value)
				refy := reflect.TypeOf(rt.vars[ind])
				if refx != refy {
					return fmt.Errorf("func tail '%s' param: cannot use (type %s) as the type %s", key, refx, refy)
				}
			}
			rt.setVar(ind, value)
		}
	}
	return nil
}
