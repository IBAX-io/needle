package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/IBAX-io/needle/compiler"

	"github.com/shopspring/decimal"
)

const (
	brackets = `[]`
)

const (
	statusNormal = iota
	statusReturn
	statusContinue
	statusBreak
)

type instruction func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error)

type instructionCtx struct {
	ci        int
	labels    []int
	assignVar []*compiler.VarInfo
	slice     *compiler.SliceItem
	ifCond    bool
}

func newInstructionCtx() *instructionCtx {
	return &instructionCtx{
		labels:    make([]int, 0),
		assignVar: make([]*compiler.VarInfo, 0),
	}
}

func (c *instructionCtx) popLabel() int {
	if len(c.labels) == 0 {
		return 0
	}
	label := c.labels[len(c.labels)-1]
	c.labels = c.labels[:len(c.labels)-1]
	return label
}

var instructionTable = make(map[compiler.Cmd]instruction)

func init() {
	instructionTable[compiler.CmdCallExtend] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.SubCost(CostExtend); err != nil {
			return
		}
		err = rt.callExtendFunc(code.Value.(string))
		if err != nil {
			err = fmt.Errorf("extend function '$%s' error: %v", code.Value.(string), err)
		}
		return
	}
	instructionTable[compiler.CmdExtend] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.SubCost(CostExtend); err != nil {
			return
		}
		val, ok := rt.extend[code.Value.(string)]
		if !ok {
			err = fmt.Errorf("unknown extend identifier '$%v'", code.Value)
			return
		}
		rt.stack.push(val)
		return
	}
	instructionTable[compiler.CmdPushStr] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(code.Value.(string))
		return
	}
	for i := compiler.CmdCall; i <= compiler.CmdCallVariadic; i++ {
		instructionTable[i] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
			cost := int64(CostCall)
			if code.Object().Type == compiler.ObjExtFunc {
				finfo := code.Object().GetExtFuncInfo()
				if rt.vm.ExtCost != nil {
					cost = rt.vm.ExtCost(finfo.Name)
					if cost < 0 {
						cost = CostCall
					}
				}
			}
			if err = rt.SubCost(cost); err != nil {
				return
			}
			if rt.callDepth >= MaxCallDepth {
				err = fmt.Errorf("max call depth of recursive call")
				return
			}

			rt.callDepth++
			defer func() {
				rt.callDepth--
			}()
			obj := code.Object()
			if obj.Type == compiler.ObjFunction {
				err = rt.callObjFunc(obj)
				return
			}
			err = rt.callFunc(code.Object())
			return
		}
	}
	instructionTable[compiler.CmdReturn] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusReturn
		return
	}
	instructionTable[compiler.CmdIf] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.stack.CheckDepth(1); err != nil {
			return
		}
		ret := rt.stack.pop()
		if !valueToBool(ret) {
			if ctx.ifCond {
				ctx.ifCond = false
			}
			if len(rt.peekBlock().Block.Code) > ctx.ci+1 &&
				rt.peekBlock().Block.Code[ctx.ci+1].Cmd == compiler.CmdElse {
				rt.stack.push(ret)
			}
			return
		}
		ctx.ifCond = true
		return rt.RunCode(code.CodeBlock())
	}

	instructionTable[compiler.CmdElse] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if ctx.ifCond {
			ctx.ifCond = false
			return
		}
		if err = rt.stack.CheckDepth(1); err != nil {
			return
		}
		ret := rt.stack.pop()
		if !valueToBool(ret) {
			return rt.RunCode(code.CodeBlock())
		}
		return
	}

	instructionTable[compiler.CmdAssignVar] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.assignVar = code.VarInfos()
		for _, item := range ctx.assignVar {
			if item.Owner == nil && item.Obj.Type == compiler.ObjExtVar {
				n := item.Obj.GetName()
				if rt.vm.AssertVar(n) {
					err = fmt.Errorf(eSysVar, n)
					return
				}
			}
		}
		return
	}

	instructionTable[compiler.CmdAssign] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		count := len(ctx.assignVar)
		if count > rt.stack.size() {
			err = fmt.Errorf("not enough stack to assign")
			return
		}
		cut := count
		preCode := rt.peekBlock().Block.Code[ctx.ci-1]
		if preCode.Cmd == compiler.CmdCall || preCode.Cmd == compiler.CmdCallVariadic {
			objInfo := preCode.Object()
			resultsLen := objInfo.GetResultsLen()
			if objInfo.Type == compiler.ObjExtFunc || objInfo.Type == compiler.ObjFunction {
				if count > resultsLen {
					err = fmt.Errorf("assignments count mismatch: %d = %d", count, resultsLen)
					return
				}
				if objInfo.Type == compiler.ObjExtFunc {
					if _, ok := rt.vm.FuncCallsDB[objInfo.GetExtFuncInfo().Name]; ok {
						resultsLen--
					}
				}
				cut = resultsLen
			}
		}
		local := rt.stack.peekN(cut)
		rt.stack.resetByIdx(rt.stack.size() - cut)
		for ivar, item := range ctx.assignVar {
			val := local[ivar]
			if item.Owner == nil {
				if item.Obj.Type == compiler.ObjExtVar {
					n := item.Obj.GetName()
					if v, ok := rt.extend[n]; ok && v != nil && reflect.TypeOf(v) != reflect.TypeOf(val) {
						err = fmt.Errorf("$%s (type %v) cannot be represented by the type %s", n, reflect.TypeOf(val), reflect.TypeOf(v))
						return
					}
					rt.setExtendVar(n, val)
				}
			} else {
				for i := len(rt.blocks) - 1; i >= 0; i-- {
					if item.Owner == rt.blocks[i].Block {
						k := rt.blocks[i].Offset + item.Obj.GetVariable().Index
						switch v := rt.blocks[i].Block.Vars[item.Obj.GetVariable().Index]; v {
						case compiler.FLOAT:
							var d float64
							d, err = ValueToFloat(val)
							if err != nil {
								return
							}
							rt.setVar(k, d)
						case compiler.MONEY:
							var d decimal.Decimal
							d, err = ValueToDecimal(val)
							if err != nil {
								return
							}
							rt.setVar(k, d)
						default:
							if !v.EqualsType(val) {
								err = fmt.Errorf("variable '%v' (type %v) cannot be represented by the type %s",
									item.Obj.GetName(), reflect.TypeOf(val), v.ReflectType())
								return
							}
							rt.setVar(k, val)
						}
						break
					}
				}
			}
		}
		ctx.assignVar = ctx.assignVar[:0]
		return
	}
	instructionTable[compiler.CmdLabel] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.labels = append(ctx.labels, ctx.ci)
		return
	}

	instructionTable[compiler.CmdContinue] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusContinue
		return
	}
	instructionTable[compiler.CmdWhile] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if !valueToBool(rt.stack.pop()) {
			return
		}
		status, err = rt.RunCode(code.CodeBlock())
		newci := ctx.popLabel()
		if status == statusContinue {
			ctx.ci = newci - 1
			status = statusNormal
		}
		if status == statusBreak {
			status = statusNormal
		}
		return
	}
	instructionTable[compiler.CmdBreak] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusBreak
		return
	}
	instructionTable[compiler.CmdSliceColon] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.slice = code.SliceItem()
		return
	}
	instructionTable[compiler.CmdGetIndex] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.stack.CheckDepth(1); err != nil {
			return
		}
		var (
			ind, value any
			low, high  int
			rkt        string
			rv         reflect.Value
		)

		if ctx.slice != nil {
			if ctx.slice.Index[1] == compiler.SliceHighNum {
				h, ok := rt.stack.pop().(int64)
				if !ok {
					err = fmt.Errorf("slice high must be int64")
					return
				}
				high = int(h)
			}
			if ctx.slice.Index[0] == compiler.SliceLowNum {
				l, ok := rt.stack.pop().(int64)
				if !ok {
					err = fmt.Errorf("slice low must be int64")
					return
				}
				low = int(l)
			}
			rkt = "slice"
		} else {
			ind = rt.stack.pop()
			rkt = reflect.TypeOf(ind).String()
		}
		value = rt.stack.pop()
		rv = reflect.ValueOf(value)
		rvt := reflect.TypeOf(value).String()
		switch {
		case rvt == `*compiler.Map`:
			if rkt != `string` {
				err = fmt.Errorf(eMapIndex, rkt)
				break
			}
			v, found := value.(*compiler.Map).Get(ind.(string))
			if found {
				rt.stack.push(v)
			} else {
				rt.stack.push(nil)
			}
		case rvt[:2] == brackets:
			if rkt != `int64` && rkt != `slice` {
				err = fmt.Errorf(eArrIndex, rkt)
				break
			}

			var ret reflect.Value
			if rkt == `slice` {
				if ctx.slice.Index[1] != compiler.SliceHighNum {
					high = rv.Len()
				}
				if low < 0 {
					err = fmt.Errorf("invalid slice index must be non-negative")
					return
				}
				if low > high {
					err = fmt.Errorf("invalid index values, must be low <= high")
					return
				}
				if high > rv.Len() {
					err = fmt.Errorf("index out of range [%d:%d] with length %d", low, high, rv.Len())
					return
				}
				ret = rv.Slice(low, high)
			} else {
				if ind.(int64) < 0 {
					err = fmt.Errorf("invalid index must be non-negative")
					return
				}
				if rv.Len() <= int(ind.(int64)) {
					err = fmt.Errorf("index out of range [%d] with length %d", ind, rv.Len())
					return
				}
				ret = rv.Index(int(ind.(int64)))
			}
			if ret.IsValid() {
				rt.stack.push(ret.Interface())
			} else {
				rt.stack.push(nil)
			}
			ctx.slice = nil
		case rvt == "string":
			if rkt != `int64` && rkt != `slice` {
				err = fmt.Errorf(eArrIndex, rkt)
				break
			}
			if rkt == `slice` {
				if ctx.slice.Index[1] != compiler.SliceHighNum {
					high = rv.Len()
				}
				if low < 0 {
					err = fmt.Errorf("invalid slice index must be non-negative")
					return
				}
				if low > high {
					err = fmt.Errorf("invalid index values, must be low <= high")
					return
				}
				if high > rv.Len() {
					err = fmt.Errorf("index out of range [%d:%d] with length %d", low, high, rv.Len())
					return
				}
				rt.stack.push(string([]rune(rv.String())[low:high]))
				break
			}
			if ind.(int64) < 0 {
				err = fmt.Errorf("invalid index must be non-negative")
				return
			}
			if rv.Len() <= int(ind.(int64)) {
				err = fmt.Errorf("index out of range [%d] with length %d", ind, rv.Len())
				return
			}
			rt.stack.push(string([]rune(rv.String())[ind.(int64)]))
			ctx.slice = nil
		default:
			err = fmt.Errorf(`type %s doesn't support indexing`, rvt)
		}
		return
	}
	instructionTable[compiler.CmdVar] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		ivar := code.VarInfo()
		var i int
		for i = len(rt.blocks) - 1; i >= 0; i-- {
			if ivar.Owner == rt.blocks[i].Block {
				rt.stack.push(rt.vars[rt.blocks[i].Offset+ivar.Obj.GetVariable().Index])
				break
			}
		}
		if i < 0 {
			err = fmt.Errorf(`wrong var %v`, ivar.Obj.Value)
		}
		return
	}
	instructionTable[compiler.CmdSetIndex] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.stack.CheckDepth(3); err != nil {
			return
		}

		value := rt.stack.pop()
		key := rt.stack.pop()
		indextype := rt.stack.pop()
		rkt := reflect.TypeOf(key).String()
		itype := reflect.TypeOf(indextype).String()
		if isSelfAssignment(indextype, value) {
			err = errSelfAssignment
			return
		}
		var indexKey int
		indexInfo := code.IndexInfo()
		if indexInfo.Owner != nil {
			for i := len(rt.blocks) - 1; i >= 0; i-- {
				if indexInfo.Owner == rt.blocks[i].Block {
					indexKey = rt.blocks[i].Offset + indexInfo.VarOffset
					break
				}
			}
		}
		switch {
		case itype == `*compiler.Map`:
			if indextype.(*compiler.Map).Size() > MaxMapCount {
				err = errMaxMapCount
				return
			}
			if rkt != `string` {
				err = fmt.Errorf(eMapIndex, rkt)
				return
			}
			indextype.(*compiler.Map).Set(key.(string), value)
			rt.stack.push(indextype)
		case itype[:2] == brackets:
			if rkt != `int64` {
				err = fmt.Errorf(eArrIndex, rkt)
				return
			}
			ind := key.(int64)
			if strings.Contains(itype, "interface") {
				slice := indextype.([]any)
				if int(ind) >= len(slice) {
					if ind > MaxArrayIndex {
						err = errMaxArrayIndex
						return
					}
					slice = append(slice, make([]any, int(ind)-len(slice)+1)...)
					if indexInfo.Owner == nil { // Extend variable $varname
						rt.extend[indexInfo.Extend] = slice
					} else {
						rt.vars[indexKey] = slice
					}
				}
				slice[ind] = value
				// rt.stack.push(slice)
			} else {
				slice := indextype.([]map[string]string)
				slice[ind] = value.(map[string]string)
				// rt.stack.push(slice)
			}
		default:
			err = fmt.Errorf(`type %s doesn't support indexing`, itype)
			return
		}
		if indexInfo.Owner == nil {
			rt.recalculateMemExtendVar(indexInfo.Extend)
		} else {
			rt.recalculateMemVar(indexKey)
		}
		return
	}
	instructionTable[compiler.CmdPush] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(code.Value)
		return
	}
	instructionTable[compiler.CmdFuncTail] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		f := code.FuncTailCmd()
		off := rt.stack.size() - f.Count
		if off < 0 {
			err = fmt.Errorf("not enough stack to assign for tail function '%v'", f.FuncTail.Name)
			return
		}
		if err = rt.stack.CheckDepth(f.Count + 1); err != nil {
			return
		}

		params := rt.stack.popN(f.Count)
		names := rt.stack.pop()
		if names == nil {
			names = make(map[string][]any)
		}

		if rt.unwrap && len(params) > 0 && reflect.TypeOf(params[len(params)-1]).String() == `[]interface {}` {
			params = append(params[:len(params)-1], params[len(params)-1].([]any)...)
			rt.unwrap = false
		}
		names.(map[string][]any)[f.FuncTail.Name] = params
		rt.stack.push(names)
		return
	}

	instructionTable[compiler.CmdUnwrapArr] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if reflect.TypeOf(rt.stack.peek()).String() != `[]interface {}` {
			err = fmt.Errorf(`invalid use of '...'`)
			return
		}
		rt.unwrap = true
		return
	}
	instructionTable[compiler.CmdMapInit] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initMap *compiler.Map
		initMap, err = rt.getResultMap(code.Map())
		if err != nil {
			return
		}
		rt.stack.push(initMap)
		return
	}
	instructionTable[compiler.CmdArrayInit] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initArray []any
		initArray, err = rt.getResultArray(code.MapItems())
		if err != nil {
			return
		}
		rt.stack.push(initArray)
		return
	}
	instructionTable[compiler.CmdError] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		err = VMError{Type: code.Value.(string), Err: rt.stack.pop(), Line: code.Lexeme.Line, Column: code.Lexeme.Column}
		return
	}
	instructionTable[compiler.CmdNot] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(!valueToBool(rt.stack.pop()))
		return
	}
	instructionTable[compiler.CmdSign] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
		if code.Lexeme.Token() == compiler.Add {
			return
		}
		z := rt.stack.pop()
		switch z.(type) {
		case float64:
			rt.stack.push(-z.(float64))
		case int64:
			rt.stack.push(-z.(int64))
		default:
			err = errUnsupportedType
			return
		}
		return
	}
	for _, c := range []compiler.Cmd{
		compiler.CmdInc, compiler.CmdDec,
		compiler.CmdAssignAdd, compiler.CmdAssignSub,
		compiler.CmdAssignMul, compiler.CmdAssignDiv, compiler.CmdAssignMod,
		compiler.CmdAssignAnd, compiler.CmdAssignOr, compiler.CmdAssignXor,
		compiler.CmdAssignLShift, compiler.CmdAssignRShift,
	} {
		instructionTable[c] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
			if len(ctx.assignVar) != 1 {
				err = fmt.Errorf("assign op %s variable count must be 1", code.Cmd)
				return
			}
			var y any
			if code.Cmd == compiler.CmdInc || code.Cmd == compiler.CmdDec {
				y = 1
			} else {
				y = rt.stack.pop()
			}
			item := ctx.assignVar[0]
			if item.Owner == nil {
				if item.Obj.Type != compiler.ObjExtVar {
					err = fmt.Errorf("can not assign to %s", item.Obj.Type)
					return
				}
				n := item.Obj.GetName()
				var ret any
				ret, err = operationExpr(rt.extend[n], y, code.Cmd)
				if err != nil {
					return
				}
				rt.setExtendVar(n, ret)
				return
			}
			for i := len(rt.blocks) - 1; i >= 0; i-- {
				if item.Owner == rt.blocks[i].Block {
					k := rt.blocks[i].Offset + item.Obj.GetVariable().Index
					var ret any
					ret, err = operationExpr(rt.vars[k], y, code.Cmd)
					if err != nil {
						return
					}
					rt.setVar(k, ret)
					break
				}
			}
			ctx.assignVar = ctx.assignVar[:0]
			return
		}
	}
	for _, c := range []compiler.Cmd{
		compiler.CmdAdd, compiler.CmdSub, compiler.CmdMul, compiler.CmdDiv, compiler.CmdMod,
		compiler.CmdOr, compiler.CmdAnd,
		compiler.CmdEqual, compiler.CmdNotEq,
		compiler.CmdLess, compiler.CmdGrEq,
		compiler.CmdGreat, compiler.CmdLessEq,
		compiler.CmdShiftL, compiler.CmdShiftR,
		compiler.CmdBitAnd, compiler.CmdBitOr, compiler.CmdBitXor,
	} {
		instructionTable[c] = func(rt *Runtime, code *compiler.ByteCode, ctx *instructionCtx) (status int, err error) {
			var z any
			y := rt.stack.pop()
			x := rt.stack.pop()
			z, err = operationExpr(x, y, code.Cmd)
			if err != nil {
				return
			}
			rt.stack.push(z)
			return
		}
	}
}
