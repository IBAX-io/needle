package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/IBAX-io/needle/compile"

	"github.com/shopspring/decimal"
)

type instruction func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error)

var instructionTable = make(map[compile.CmdT]instruction)

type instructionCtx struct {
	ci         int
	size       int
	bin        any
	isContinue bool
	isBreak    bool
	isLoop     bool
	labels     []int
	top        []any
	assignVar  []*compile.VarInfo
	costRemain decimal.Decimal
}

func newInstructionCtx() *instructionCtx {
	return &instructionCtx{
		labels:    make([]int, 0),
		assignVar: make([]*compile.VarInfo, 0),
	}
}

func init() {
	instructionTable[compile.CmdPush] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.push(code.Value)
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
			ctx.isLoop = true
		}
		return
	}
	for i := compile.CmdExtend; i <= compile.CmdCallExtend; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			if err = rt.SubCost(CostExtend); err != nil {
				return
			}
			if val, ok := rt.extend[code.Value.(string)]; ok {
				if code.Cmd == compile.CmdCallExtend {
					err = rt.extendFunc(code.Value.(string))
					if err != nil {
						err = fmt.Errorf(`extend function %v %s`, code.Value, err)
						ctx.isLoop = true
						return
					}
				} else {
					switch varVal := val.(type) {
					case int:
						val = int64(varVal)
					}
					rt.stack.push(val)
				}
			} else {
				err = fmt.Errorf(`unknown extend identifier %v`, code.Value)
			}
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
			if code.Value.(*compile.ObjInfo).Type == compile.ObjectType_ExtFunc {
				finfo := code.Value.(*compile.ObjInfo).GetExtFuncInfo()
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
			err = rt.callFunc(code.Cmd, code.Value.(*compile.ObjInfo))
			return
		}
	}
	instructionTable[compile.CmdReturn] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusReturn
		return
	}
	instructionTable[compile.CmdIf] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if valueToBool(rt.stack.peek()) {
			status, err = rt.RunCode(code.Value.(*compile.CodeBlock))
		}
		return
	}
	instructionTable[compile.CmdElse] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if !valueToBool(rt.stack.peek()) {
			status, err = rt.RunCode(code.Value.(*compile.CodeBlock))
		}
		return
	}

	instructionTable[compile.CmdAssignVar] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.assignVar = code.Value.([]*compile.VarInfo)
		for _, item := range ctx.assignVar {
			if item.Owner == nil {
				if item.Obj.Type == compile.ObjectType_ExtVar {
					var n = item.Obj.GetExtendVariable().Name
					if rt.vm.AssertVar(n) {
						err = fmt.Errorf(eSysVar, n)
						ctx.isLoop = true
						return
					}
				}
			}
		}
		return
	}

	instructionTable[compile.CmdAssign] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		count := len(ctx.assignVar)
		for ivar, item := range ctx.assignVar {
			val := rt.stack.get(rt.stack.size() - count + ivar)
			if item.Owner == nil {
				if item.Obj.Type == compile.ObjectType_ExtVar {
					var n = item.Obj.GetExtendVariable().Name
					if v, ok := rt.extend[n]; ok && v != nil && reflect.TypeOf(v) != reflect.TypeOf(val) {
						err = fmt.Errorf("$%s (type %s) cannot be represented by the type %s", n, reflect.TypeOf(val), reflect.TypeOf(v))
						return
					}
					rt.setExtendVar(n, val)
				}
			} else {
				for i := len(rt.blocks) - 1; i >= 0; i-- {
					if item.Owner == rt.blocks[i].Block {
						k := rt.blocks[i].Offset + item.Obj.GetVariable().Index
						switch v := rt.blocks[i].Block.Vars[item.Obj.GetVariable().Index]; v.String() {
						case Decimal:
							var v decimal.Decimal
							v, err = ValueToDecimal(val)
							if err != nil {
								ctx.isLoop = true
								return
							}
							rt.setVar(k, v)
						default:
							if val != nil && v != reflect.TypeOf(val) {
								err = fmt.Errorf("variable '%v' (type %s) cannot be represented by the type %s", item.Obj.GetVariable().Name, reflect.TypeOf(val), v)
								break
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
		val := rt.stack.pop()
		if valueToBool(val) {
			status, err = rt.RunCode(code.Value.(*compile.CodeBlock))
			newci := ctx.labels[len(ctx.labels)-1]
			ctx.labels = ctx.labels[:len(ctx.labels)-1]
			if status == statusContinue {
				ctx.ci = newci - 1
				status = statusNormal
				//ctx.isContinue = true
				return
			}
			if status == statusBreak {
				status = statusNormal
				//ctx.isBreak = true
				return
			}
		}
		return
	}
	instructionTable[compile.CmdBreak] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		status = statusBreak
		return
	}
	instructionTable[compile.CmdGetIndex] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
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
	instructionTable[compile.CmdSetIndex] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		value := rt.stack.pop()
		key := rt.stack.pop()
		indextype := rt.stack.pop()
		rkt := reflect.TypeOf(key).String()
		itype := reflect.TypeOf(indextype).String()
		if isSelfAssignment(indextype, value) {
			err = errSelfAssignment
			ctx.isLoop = true
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
			if indextype.(*compile.Map).Size() > maxMapCount {
				err = errMaxMapCount
				break
			}
			if rkt != `string` {
				err = fmt.Errorf(eMapIndex, rkt)
				break
			}
			indextype.(*compile.Map).Set(key.(string), value)
			rt.stack.push(indextype)
		case itype[:2] == brackets:
			if rkt != `int64` {
				err = fmt.Errorf(eArrIndex, rkt)
				break
			}
			ind := key.(int64)
			if strings.Contains(itype, Interface) {
				slice := indextype.([]any)
				if int(ind) >= len(slice) {
					if ind > maxArrayIndex {
						err = errMaxArrayIndex
						break
					}
					slice = append(slice, make([]any, int(ind)-len(slice)+1)...)
					if indexInfo.Owner == nil { // Extend variable $varname
						rt.extend[indexInfo.Extend] = slice
					} else {
						rt.vars[indexKey] = slice
					}
				}
				slice[ind] = value
				rt.stack.push(slice)
			} else {
				slice := indextype.([]map[string]string)
				slice[ind] = value.(map[string]string)
				rt.stack.push(slice)
			}
		default:
			err = fmt.Errorf(`type %s doesn't support indexing`, itype)
		}
		if indexInfo.Owner == nil {
			rt.recalcMemExtendVar(indexInfo.Extend)
		} else {
			rt.recalculateMemVar(indexKey)
		}
		return
	}
	instructionTable[compile.CmdFuncName] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ifunc := code.Value.(compile.FuncNameCmd)
		mapoff := rt.stack.size() - 1 - ifunc.Count
		if rt.stack.get(mapoff) == nil {
			rt.stack.set(mapoff, make(map[string][]any))
		}
		params := make([]any, 0, ifunc.Count)
		for i := 0; i < ifunc.Count; i++ {
			cur := rt.stack.get(mapoff + 1 + i)
			if i == ifunc.Count-1 && rt.unwrap &&
				reflect.TypeOf(cur).String() == `[]interface {}` {
				params = append(params, cur.([]any)...)
				rt.unwrap = false
			} else {
				params = append(params, cur)
			}
		}
		rt.stack.get(mapoff).(map[string][]any)[ifunc.Name] = params
		rt.stack.resetByIdx(mapoff + 1)
		ctx.isContinue = true
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
			ctx.isLoop = true
			return
		}
		rt.stack.push(initMap)
		return
	}
	instructionTable[compile.CmdArrayInit] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initArray []any
		initArray, err = rt.getResultArray(code.Value.([]compile.MapItem))
		if err != nil {
			ctx.isLoop = true
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
		err = SetVMError(eType, rt.stack.peek())
		return
	}
	instructionTable[compile.CmdNot] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack.set(ctx.size-1, !valueToBool(ctx.top[0]))
		return
	}
	instructionTable[compile.CmdSign] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if code.Lexeme.Value.(compile.Token) == compile.Add {
			return
		}
		switch ctx.top[0].(type) {
		case float64:
			rt.stack.set(ctx.size-1, -ctx.top[0].(float64))
		case int64:
			rt.stack.set(ctx.size-1, -ctx.top[0].(int64))
		default:
			err = errUnsupportedType
			ctx.isLoop = true
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
				ctx.isLoop = true
				err = fmt.Errorf("assign op %s variable count must be 1", code.Cmd)
				return
			}
			y := ctx.top[0]
			item := ctx.assignVar[0]
			if item.Owner == nil {
				if item.Obj.Type != compile.ObjectType_ExtVar {
					err = fmt.Errorf("can not assign to %s", item.Obj.Type)
					return
				}
				var n = item.Obj.GetExtendVariable().Name
				var ret any
				ret, err = evaluateCmd(rt.extend[n], y, code.Cmd)
				if err != nil {
					ctx.isLoop = true
					return
				}
				rt.setExtendVar(n, ret)
				rt.stack.set(ctx.size-1, ret)
				return
			}
			for i := len(rt.blocks) - 1; i >= 0; i-- {
				if item.Owner == rt.blocks[i].Block {
					k := rt.blocks[i].Offset + item.Obj.GetVariable().Index
					var ret any
					ret, err = evaluateCmd(rt.vars[k], y, code.Cmd)
					if err != nil {
						ctx.isLoop = true
						return
					}
					rt.setVar(k, ret)
					rt.stack.set(ctx.size-1, ret)
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
			var ret any
			x := ctx.top[1]
			y := ctx.top[0]
			ret, err = evaluateCmd(x, y, code.Cmd)
			if err != nil {
				ctx.isLoop = true
				return
			}
			ctx.bin = ret
			return
		}
	}
}
