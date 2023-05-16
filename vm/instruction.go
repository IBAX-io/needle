package vm

import (
	"fmt"
	"github.com/IBAX-io/needle/compile"
	"reflect"
	"strings"

	"github.com/shopspring/decimal"
)

type instruction func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error)

var instructionTable = make(map[compile.CmdT]instruction)

type instructionCtx struct {
	labels     []int
	ci         int
	assignVar  []*compile.VarInfo
	top        []any
	isContinue bool
	isBreak    bool
	isLoop     bool
	size       int
	bin        any
}

func newInstructionCtx() *instructionCtx {
	return &instructionCtx{
		labels:    make([]int, 0),
		assignVar: make([]*compile.VarInfo, 0),
		//top:       make([]any, 8),
	}
}

func init() {
	instructionTable[compile.CmdPush] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.push(code.Value)
		return
	}
	instructionTable[compile.CmdVar] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ivar := code.Value.(*compile.VarInfo)
		var i int
		for i = len(rt.blocks) - 1; i >= 0; i-- {
			if ivar.Owner == rt.blocks[i].Block {
				rt.push(rt.vars[rt.blocks[i].Offset+ivar.Obj.GetVariable().Index])
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
					rt.push(val)
				}
			} else {
				err = fmt.Errorf(`unknown extend identifier %v`, code.Value)
			}
			return
		}
	}
	instructionTable[compile.CmdPushStr] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.push(code.Value.(string))
		return
	}
	for i := compile.CmdCall; i <= compile.CmdCallVariadic; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			var cost = int64(CostCall)
			if code.Value.(*compile.ObjInfo).Type == compile.ObjectType_ExtFunc {
				finfo := code.Value.(*compile.ObjInfo).GetExtFuncInfo()
				if rt.vm.ExtCost != nil {
					cost = rt.vm.ExtCost(finfo.Name)
					if cost == -1 {
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
		if valueToBool(rt.peek()) {
			status, err = rt.RunCode(code.Value.(*compile.CodeBlock))
		}
		return
	}
	instructionTable[compile.CmdElse] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if !valueToBool(rt.peek()) {
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
					if rt.limitName(n) {
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
			val := rt.stack[rt.len()-count+ivar]
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
		val := rt.peek()
		rt.resetByIdx(rt.len() - 1)
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
	instructionTable[compile.CmdIndex] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rvalue := reflect.ValueOf(rt.stack[ctx.size-2])
		rtype := reflect.TypeOf(rt.stack[ctx.size-2]).String()
		switch {
		case rtype == `*script.Map`:
			key := rt.getStack(ctx.size - 1)
			if reflect.TypeOf(key).String() != `string` {
				err = fmt.Errorf(eMapIndex, reflect.TypeOf(key).String())
				break
			}
			v, found := rt.stack[ctx.size-2].(*compile.Map).Get(key.(string))
			if found {
				rt.stack[ctx.size-2] = v
			} else {
				rt.stack[ctx.size-2] = nil
			}
			rt.resetByIdx(ctx.size - 1)
		case rtype[:2] == brackets:
			index := rt.getStack(ctx.size - 1)
			indexT := reflect.TypeOf(index).String()
			if indexT != `int64` {
				err = fmt.Errorf(eArrIndex, indexT)
				break
			}
			if int(index.(int64)) >= rvalue.Len() {
				err = fmt.Errorf("index out of range [%d] with length %d", index, rvalue.Len())
				break
			}
			v := rvalue.Index(int(index.(int64)))
			if v.IsValid() {
				rt.stack[ctx.size-2] = v.Interface()
			} else {
				rt.stack[ctx.size-2] = nil
			}
			rt.resetByIdx(ctx.size - 1)
		default:
			err = fmt.Errorf(`type %s doesn't support indexing`, rtype)
		}
		return
	}
	instructionTable[compile.CmdSetIndex] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		itype := reflect.TypeOf(rt.stack[ctx.size-3]).String()
		indexInfo := code.Value.(*compile.IndexInfo)
		var indexKey int
		if indexInfo.Owner != nil {
			for i := len(rt.blocks) - 1; i >= 0; i-- {
				if indexInfo.Owner == rt.blocks[i].Block {
					indexKey = rt.blocks[i].Offset + indexInfo.VarOffset
					break
				}
			}
		}
		if isSelfAssignment(rt.stack[ctx.size-3], rt.getStack(ctx.size-1)) {
			err = errSelfAssignment
			ctx.isLoop = true
			return
		}

		switch {
		case itype == `*script.Map`:
			if rt.stack[ctx.size-3].(*compile.Map).Size() > maxMapCount {
				err = errMaxMapCount
				break
			}
			if reflect.TypeOf(rt.stack[ctx.size-2]).String() != `string` {
				err = fmt.Errorf(eMapIndex, reflect.TypeOf(rt.stack[ctx.size-2]).String())
				break
			}
			rt.stack[ctx.size-3].(*compile.Map).Set(rt.stack[ctx.size-2].(string),
				reflect.ValueOf(rt.getStack(ctx.size-1)).Interface())
			rt.resetByIdx(ctx.size - 2)
		case itype[:2] == brackets:
			if reflect.TypeOf(rt.stack[ctx.size-2]).String() != `int64` {
				err = fmt.Errorf(eArrIndex, reflect.TypeOf(rt.stack[ctx.size-2]).String())
				break
			}
			ind := rt.stack[ctx.size-2].(int64)
			if strings.Contains(itype, Interface) {
				slice := rt.stack[ctx.size-3].([]any)
				if int(ind) >= len(slice) {
					if ind > maxArrayIndex {
						err = errMaxArrayIndex
						break
					}
					slice = append(slice, make([]any, int(ind)-len(slice)+1)...)
					indexInfo := code.Value.(*compile.IndexInfo)
					if indexInfo.Owner == nil { // Extend variable $varname
						rt.extend[indexInfo.Extend] = slice
					} else {
						rt.vars[indexKey] = slice
					}
					rt.stack[ctx.size-3] = slice
				}
				slice[ind] = rt.getStack(ctx.size - 1)
			} else {
				slice := rt.getStack(ctx.size - 3).([]map[string]string)
				slice[ind] = rt.getStack(ctx.size - 1).(map[string]string)
			}
			rt.resetByIdx(ctx.size - 2)
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
		mapoff := rt.len() - 1 - ifunc.Count
		if rt.stack[mapoff] == nil {
			rt.stack[mapoff] = make(map[string][]any)
		}
		params := make([]any, 0, ifunc.Count)
		for i := 0; i < ifunc.Count; i++ {
			cur := rt.stack[mapoff+1+i]
			if i == ifunc.Count-1 && rt.unwrap &&
				reflect.TypeOf(cur).String() == `[]interface {}` {
				params = append(params, cur.([]any)...)
				rt.unwrap = false
			} else {
				params = append(params, cur)
			}
		}
		rt.stack[mapoff].(map[string][]any)[ifunc.Name] = params
		rt.resetByIdx(mapoff + 1)
		ctx.isContinue = true
		return
	}
	instructionTable[compile.CmdUnwrapArr] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if reflect.TypeOf(rt.getStack(ctx.size-1)).String() == `[]interface {}` {
			rt.unwrap = true
		}
		return
	}
	instructionTable[compile.CmdMapInit] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initMap *compile.Map
		initMap, err = rt.getResultMap(code.Value.(*compile.Map))
		if err != nil {
			ctx.isLoop = true
			return
		}
		rt.push(initMap)
		return
	}
	instructionTable[compile.CmdArrayInit] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		var initArray []any
		initArray, err = rt.getResultArray(code.Value.([]compile.MapItem))
		if err != nil {
			ctx.isLoop = true
			return
		}
		rt.push(initArray)
		return
	}
	instructionTable[compile.CmdError] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		eType := "error"
		if code.Value.(compile.Token) == compile.ERRWARNING {
			eType = "warning"
		} else if code.Value.(compile.Token) == compile.ERRINFO {
			eType = "info"
		}
		err = SetVMError(eType, rt.peek())
		return
	}
	instructionTable[compile.CmdNot] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		rt.stack[ctx.size-1] = !valueToBool(ctx.top[0])
		return
	}
	instructionTable[compile.CmdSign] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if code.Lexeme.Value.(compile.Token) == compile.Add {
			return
		}
		switch ctx.top[0].(type) {
		case float64:
			rt.stack[ctx.size-1] = -ctx.top[0].(float64)
		case int64:
			rt.stack[ctx.size-1] = -ctx.top[0].(int64)
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
				err = fmt.Errorf("assign op variable count must be 1")
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
				rt.stack[ctx.size-1] = ret
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
					rt.stack[ctx.size-1] = ret
					break
				}
			}
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
