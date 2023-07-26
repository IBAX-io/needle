package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/IBAX-io/needle/compile"

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

type instruction func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error)

type instructionCtx struct {
	ci        int
	labels    []int
	assignVar []*compile.VarInfo
}

func newInstructionCtx() *instructionCtx {
	return &instructionCtx{
		labels:    make([]int, 0),
		assignVar: make([]*compile.VarInfo, 0),
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

var instructionTable = make(map[compile.CmdT]instruction)

func init() {
	for i := compile.CmdExtend; i <= compile.CmdCallExtend; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			if err = rt.SubCost(CostExtend); err != nil {
				return
			}
			val, ok := rt.extend[code.Value.(string)]
			if !ok {
				err = fmt.Errorf(`unknown extend identifier %v`, code.Value)
				return
			}
			if code.Cmd == compile.CmdCallExtend {
				err = rt.extendFunc(code.Value.(string))
				if err != nil {
					err = fmt.Errorf(`extend function %v %s`, code.Value, err)
				}
				return
			}
			switch varVal := val.(type) {
			case int:
				val = int64(varVal)
			}
			rt.stack.push(val)
			return
		}
	}
	instructionTable[compile.CmdPushStr] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(code.Value.(string))
		return
	}
	for i := compile.CmdCall; i <= compile.CmdCallVariadic; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			var cost = int64(CostCall)
			if code.Value.(*compile.Object).Type == compile.ObjExtFunc {
				finfo := code.Value.(*compile.Object).GetExtFuncInfo()
				if rt.vm.ExtCost != nil {
					cost = rt.vm.ExtCost(finfo.Name)
					if cost < 1 {
						cost = CostCall
					}
				}
			}
			if err = rt.SubCost(cost); err != nil {
				return
			}

			err = rt.callFunc(code.Value.(*compile.Object))
			return
		}
	}
	instructionTable[compile.CmdReturn] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusReturn
		return
	}
	instructionTable[compile.CmdIf] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if valueToBool(rt.stack.peek()) {
			rt.stack.pop()
			return rt.RunCode(code.Value.(*compile.CodeBlock))
		}
		return
	}
	instructionTable[compile.CmdElse] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if !valueToBool(rt.stack.pop()) {
			return rt.RunCode(code.Value.(*compile.CodeBlock))
		}
		return
	}

	instructionTable[compile.CmdAssignVar] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.assignVar = code.Value.([]*compile.VarInfo)
		for _, item := range ctx.assignVar {
			if item.Owner == nil && item.Obj.Type == compile.ObjExtVar {
				var n = item.Obj.GetExtendVariable().Name
				if rt.vm.AssertVar(n) {
					err = fmt.Errorf(eSysVar, n)
					return
				}
			}
		}
		return
	}

	instructionTable[compile.CmdAssign] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		count := len(ctx.assignVar)
		if count > rt.stack.size() {
			err = fmt.Errorf("not enough values to assign")
			return
		}
		cut := count
		preCode := rt.peekBlock().Block.Code[ctx.ci-1]
		if preCode.Cmd == compile.CmdCall || preCode.Cmd == compile.CmdCallVariadic {
			objInfo := preCode.Value.(*compile.Object)
			resultsLen := objInfo.GetResultsLen()
			if objInfo.Type == compile.ObjExtFunc || objInfo.Type == compile.ObjFunc {
				if count > resultsLen {
					err = fmt.Errorf("assignments count mismatch: %d = %d", count, resultsLen)
					return
				}
				cut = resultsLen
			}
		}
		local := rt.stack.peekN(cut)
		rt.stack.resetByIdx(rt.stack.size() - cut)
		for ivar, item := range ctx.assignVar {
			val := local[ivar]
			if item.Owner == nil {
				if item.Obj.Type == compile.ObjExtVar {
					var n = item.Obj.GetExtendVariable().Name
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
						switch v := rt.blocks[i].Block.Vars[item.Obj.GetVariable().Index]; v.String() {
						case "float64":
							var d decimal.Decimal
							d, err = ValueToDecimal(val)
							if err != nil {
								return
							}
							rt.setVar(k, d.InexactFloat64())
						case "decimal.Decimal":
							var d decimal.Decimal
							d, err = ValueToDecimal(val)
							if err != nil {
								return
							}
							rt.setVar(k, d)
						default:
							if val != nil && v != reflect.TypeOf(val) {
								err = fmt.Errorf("variable '%v' (type %v) cannot be represented by the type %s", item.Obj.GetVariable().Name, reflect.TypeOf(val), v)
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
	instructionTable[compile.CmdLabel] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.labels = append(ctx.labels, ctx.ci)
		return
	}

	instructionTable[compile.CmdContinue] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusContinue
		return
	}
	instructionTable[compile.CmdWhile] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if !valueToBool(rt.stack.pop()) {
			return
		}
		status, err = rt.RunCode(code.Value.(*compile.CodeBlock))
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
	instructionTable[compile.CmdBreak] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusBreak
		return
	}
	instructionTable[compile.CmdGetIndex] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.stack.CheckDepth(2); err != nil {
			return
		}
		ind := rt.stack.pop()
		value := rt.stack.pop()
		rv := reflect.ValueOf(value)
		rkt := reflect.TypeOf(ind).String()
		rvt := reflect.TypeOf(value).String()
		switch {
		case rvt == `*compile.Map`:
			if rkt != `string` {
				err = fmt.Errorf(eMapIndex, rkt)
				break
			}
			v, found := value.(*compile.Map).Get(ind.(string))
			if found {
				rt.stack.push(v)
			} else {
				rt.stack.push(nil)
			}
		case rvt[:2] == brackets:
			if rkt != `int64` {
				err = fmt.Errorf(eArrIndex, rkt)
				break
			}
			if rv.Len() <= int(ind.(int64)) {
				err = fmt.Errorf("index out of range [%d] with length %d", ind, rv.Len())
				break
			}
			v := rv.Index(int(ind.(int64)))
			if v.IsValid() {
				rt.stack.push(v.Interface())
			} else {
				rt.stack.push(nil)
			}
		default:
			err = fmt.Errorf(`type %s doesn't support indexing`, rvt)
		}
		return
	}
	instructionTable[compile.CmdVar] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ivar := code.Value.(*compile.VarInfo)
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

	instructionTable[compile.CmdSetIndex] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if err = rt.stack.CheckDepth(3); err != nil {
			return 0, err
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
		indexInfo := code.Value.(*compile.IndexInfo)
		if indexInfo.Owner != nil {
			for i := len(rt.blocks) - 1; i >= 0; i-- {
				if indexInfo.Owner == rt.blocks[i].Block {
					indexKey = rt.blocks[i].Offset + indexInfo.VarOffset
					break
				}
			}
		}
		switch {
		case itype == `*compile.Map`:
			if indextype.(*compile.Map).Size() > MaxMapCount {
				err = errMaxMapCount
				return
			}
			if rkt != `string` {
				err = fmt.Errorf(eMapIndex, rkt)
				return
			}
			indextype.(*compile.Map).Set(key.(string), value)
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
				//rt.stack.push(slice)
			} else {
				slice := indextype.([]map[string]string)
				slice[ind] = value.(map[string]string)
				//rt.stack.push(slice)
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
	instructionTable[compile.CmdPush] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(code.Value)
		return
	}
	instructionTable[compile.CmdFuncTail] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		f := code.Value.(compile.FuncTailCmd)
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

	instructionTable[compile.CmdUnwrapArr] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if reflect.TypeOf(rt.stack.peek()).String() != `[]interface {}` {
			err = fmt.Errorf(`invalid use of '...'`)
			return
		}
		rt.unwrap = true
		return
	}
	instructionTable[compile.CmdMapInit] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initMap *compile.Map
		initMap, err = rt.getResultMap(code.Value.(*compile.Map))
		if err != nil {
			return
		}
		rt.stack.push(initMap)
		return
	}
	instructionTable[compile.CmdArrayInit] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initArray []any
		initArray, err = rt.getResultArray(code.Value.([]compile.MapItem))
		if err != nil {
			return
		}
		rt.stack.push(initArray)
		return
	}
	instructionTable[compile.CmdError] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		eType := "error"
		if code.Value.(compile.Token) == compile.ERRWARNING {
			eType = "warning"
		} else if code.Value.(compile.Token) == compile.ERRINFO {
			eType = "info"
		}
		err = VMError{Type: eType, Err: rt.stack.pop()}
		return
	}
	instructionTable[compile.CmdNot] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(!valueToBool(rt.stack.pop()))
		return
	}
	instructionTable[compile.CmdSign] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if code.Lexeme.Value.(compile.Token) == compile.Add {
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
	for _, c := range []compile.CmdT{
		compile.CmdInc, compile.CmdDec,
		compile.CmdAssignAdd, compile.CmdAssignSub,
		compile.CmdAssignMul, compile.CmdAssignDiv, compile.CmdAssignMod,
		compile.CmdAssignAnd, compile.CmdAssignOr, compile.CmdAssignXor,
		compile.CmdAssignLShift, compile.CmdAssignRShift,
	} {
		instructionTable[c] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			if len(ctx.assignVar) != 1 {
				err = fmt.Errorf("assign op %s variable count must be 1", code.Cmd)
				return
			}
			var y any
			if code.Cmd == compile.CmdInc || code.Cmd == compile.CmdDec {
				y = 1
			} else {
				y = rt.stack.pop()
			}
			item := ctx.assignVar[0]
			if item.Owner == nil {
				if item.Obj.Type != compile.ObjExtVar {
					err = fmt.Errorf("can not assign to %s", item.Obj.Type)
					return
				}
				var n = item.Obj.GetExtendVariable().Name
				var ret any
				ret, err = evaluateCmd(rt.extend[n], y, code.Cmd)
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
					ret, err = evaluateCmd(rt.vars[k], y, code.Cmd)
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
	for _, c := range []compile.CmdT{
		compile.CmdAdd, compile.CmdSub, compile.CmdMul, compile.CmdDiv, compile.CmdMod,
		compile.CmdOr, compile.CmdAnd,
		compile.CmdEqual, compile.CmdNotEq,
		compile.CmdLess, compile.CmdGrEq,
		compile.CmdGreat, compile.CmdLessEq,
		compile.CmdShiftL, compile.CmdShiftR,
		compile.CmdBitAnd, compile.CmdBitOr, compile.CmdBitXor,
	} {
		instructionTable[c] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			var z any
			y := rt.stack.pop()
			x := rt.stack.pop()
			z, err = evaluateCmd(x, y, code.Cmd)
			if err != nil {
				return
			}
			rt.stack.push(z)
			return
		}
	}
}
