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
	tmpInt     int64
	tmpDec     decimal.Decimal
	top        []any
	operands   []any
	sp         int // stack position of operands
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
		operands: make([]any, 8),
		//sp:        -1,
		//tmpDec:    decimal.Zero,
		//tmpInt:    int64(0),
	}
}

func (ctx *instructionCtx) pushOperands(i int, v any) {
	ctx.sp = i
	ctx.operands[ctx.sp] = v
	//ctx.operands = append(ctx.operands, v)
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
		return
	}

	instructionTable[compile.CmdAssign] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		count := len(ctx.assignVar)
		for ivar, item := range ctx.assignVar {
			val := rt.stack[rt.len()-count+ivar]
			if item.Owner == nil {
				if item.Obj.Type == compile.ObjectType_ExtVar {
					var n = item.Obj.GetExtendVariable().Name
					if rt.limitName(n) {
						err = fmt.Errorf(eSysVar, n)
						ctx.isLoop = true
						return
					}
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

	instructionTable[compile.CmdAssignMod] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		if len(ctx.assignVar) > 1 {
			ctx.isLoop = true
			err = fmt.Errorf("mod assign not support")
			return
		}
		y := ctx.top[0]
		item := ctx.assignVar[0]
		if item.Owner == nil {
			if item.Obj.Type == compile.ObjectType_ExtVar {
				var n = item.Obj.GetExtendVariable().Name
				if rt.limitName(n) {
					err = fmt.Errorf(eSysVar, n)
					ctx.isLoop = true
					return
				}
				x, ok := rt.extend[n]
				xt := reflect.TypeOf(x)
				yt := reflect.TypeOf(y)
				if ok && x != nil && xt != yt {
					err = fmt.Errorf("$%s (type %s) cannot be represented by the type %s", n, yt, xt)
					return
				}
				var ret any
				ret, err = evaluateCmd(x, y, compile.CmdAssignMod.String())
				if err != nil {
					ctx.isLoop = true
					return
				}
				rt.setExtendVar(n, ret)
				rt.stack[ctx.size-1] = ret
			}
			return
		}
		for i := len(rt.blocks) - 1; i >= 0; i-- {
			if item.Owner == rt.blocks[i].Block {
				k := rt.blocks[i].Offset + item.Obj.GetVariable().Index
				switch v := rt.blocks[i].Block.Vars[item.Obj.GetVariable().Index]; v.String() {
				case Decimal:
					var yD decimal.Decimal
					yD, err = ValueToDecimal(y)
					if err != nil {
						ctx.isLoop = true
						return
					}
					if yD.IsZero() {
						err = errDivZero
						ctx.isLoop = true
						return
					}
					x := rt.vars[k]
					ret := x.(decimal.Decimal).Mod(yD)
					rt.setVar(k, ret)
					rt.stack[ctx.size-1] = ret
				default:
					if y != nil && v != reflect.TypeOf(y) {
						err = fmt.Errorf("variable '%v' (type %s) cannot be represented by the type %s", item.Obj.GetVariable().Name, reflect.TypeOf(y), v)
						break
					}
					switch y.(type) {
					case int64:
						if y.(int64) == 0 {
							err = errDivZero
							ctx.isLoop = true
							return
						}
						ret := rt.vars[k].(int64) % y.(int64)
						rt.setVar(k, ret)
						rt.stack[ctx.size-1] = ret
					default:
						ctx.isLoop = true
						err = fmt.Errorf(`invalid operation: the operator %s is not defined on %s`, compile.CmdAssignMod, v)
						return
					}
				}
				break
			}
		}
		return
	}
	for i := compile.CmdInc; i <= compile.CmdDec; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			return
		}
	}
	instructionTable[compile.CmdAdd] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		switch ctx.top[1].(type) {
		case string:
			switch ctx.top[0].(type) {
			case string:
				ctx.bin = ctx.top[1].(string) + ctx.top[0].(string)
			case int64:
				if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
					ctx.bin = ctx.tmpInt + ctx.top[0].(int64)
				}
			case float64:
				ctx.bin = ValueToFloat(ctx.top[1]) + ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case float64:
			switch ctx.top[0].(type) {
			case string, int64, float64:
				ctx.bin = ctx.top[1].(float64) + ValueToFloat(ctx.top[0])
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case int64:
			switch ctx.top[0].(type) {
			case string, int64:
				if ctx.tmpInt, err = ValueToInt(ctx.top[0]); err == nil {
					ctx.bin = ctx.top[1].(int64) + ctx.tmpInt
				}
			case float64:
				ctx.bin = ValueToFloat(ctx.top[1]) + ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		default:
			if reflect.TypeOf(ctx.top[1]).String() == Decimal &&
				reflect.TypeOf(ctx.top[0]).String() == Decimal {
				ctx.bin = ctx.top[1].(decimal.Decimal).Add(ctx.top[0].(decimal.Decimal))
			} else {
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		}
		return
	}
	instructionTable[compile.CmdSub] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		switch ctx.top[1].(type) {
		case string:
			switch ctx.top[0].(type) {
			case int64:
				if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
					ctx.bin = ctx.tmpInt - ctx.top[0].(int64)
				}
			case float64:
				ctx.bin = ValueToFloat(ctx.top[1]) - ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case float64:
			switch ctx.top[0].(type) {
			case string, int64, float64:
				ctx.bin = ctx.top[1].(float64) - ValueToFloat(ctx.top[0])
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case int64:
			switch ctx.top[0].(type) {
			case int64, string:
				if ctx.tmpInt, err = ValueToInt(ctx.top[0]); err == nil {
					ctx.bin = ctx.top[1].(int64) - ctx.tmpInt
				}
			case float64:
				ctx.bin = ValueToFloat(ctx.top[1]) - ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		default:
			if reflect.TypeOf(ctx.top[1]).String() == Decimal &&
				reflect.TypeOf(ctx.top[0]).String() == Decimal {
				ctx.bin = ctx.top[1].(decimal.Decimal).Sub(ctx.top[0].(decimal.Decimal))
			} else {
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		}
		return
	}
	instructionTable[compile.CmdMul] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		switch ctx.top[1].(type) {
		case string:
			switch ctx.top[0].(type) {
			case int64:
				if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
					ctx.bin = ctx.tmpInt * ctx.top[0].(int64)
				}
			case float64:
				ctx.bin = ValueToFloat(ctx.top[1]) * ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case float64:
			switch ctx.top[0].(type) {
			case string, int64, float64:
				ctx.bin = ctx.top[1].(float64) * ValueToFloat(ctx.top[0])
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case int64:
			switch ctx.top[0].(type) {
			case int64, string:
				if ctx.tmpInt, err = ValueToInt(ctx.top[0]); err == nil {
					ctx.bin = ctx.top[1].(int64) * ctx.tmpInt
				}
			case float64:
				ctx.bin = ValueToFloat(ctx.top[1]) * ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		default:
			if reflect.TypeOf(ctx.top[1]).String() == Decimal &&
				reflect.TypeOf(ctx.top[0]).String() == Decimal {
				ctx.bin = ctx.top[1].(decimal.Decimal).Mul(ctx.top[0].(decimal.Decimal))
			} else {
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		}
		return
	}
	instructionTable[compile.CmdDiv] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		switch ctx.top[1].(type) {
		case string:
			switch v := ctx.top[0].(type) {
			case int64:
				if v == 0 {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
					ctx.bin = ctx.tmpInt / v
				}
			case float64:
				if v == 0 {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				ctx.bin = ValueToFloat(ctx.top[1]) / v
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case float64:
			switch ctx.top[0].(type) {
			case string, int64, float64:
				vFloat := ValueToFloat(ctx.top[0])
				if vFloat == 0 {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				ctx.bin = ctx.top[1].(float64) / vFloat
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case int64:
			switch ctx.top[0].(type) {
			case int64, string:
				if ctx.tmpInt, err = ValueToInt(ctx.top[0]); err == nil {
					if ctx.tmpInt == 0 {
						err = errDivZero
						ctx.isLoop = true
						return
					}
					ctx.bin = ctx.top[1].(int64) / ctx.tmpInt
				}
			case float64:
				if ctx.top[0].(float64) == 0 {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				ctx.bin = ValueToFloat(ctx.top[1]) / ctx.top[0].(float64)
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		default:
			if reflect.TypeOf(ctx.top[1]).String() == Decimal &&
				reflect.TypeOf(ctx.top[0]).String() == Decimal {
				if ctx.top[0].(decimal.Decimal).Cmp(decimal.Zero) == 0 {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				ctx.bin = ctx.top[1].(decimal.Decimal).Div(ctx.top[0].(decimal.Decimal)).Floor()
			} else {
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		}
		return
	}
	instructionTable[compile.CmdMod] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		switch ctx.top[1].(type) {
		case string:
			switch y := ctx.top[0].(type) {
			case int64:
				if y == 0 {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
					ctx.bin = ctx.tmpInt % y
				}
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		case int64:
			switch ctx.top[0].(type) {
			case int64, string:
				if ctx.tmpInt, err = ValueToInt(ctx.top[0]); err == nil {
					if ctx.tmpInt == 0 {
						err = errDivZero
						ctx.isLoop = true
						return
					}
					ctx.bin = ctx.top[1].(int64) % ctx.tmpInt
				}
			default:
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		default:
			if reflect.TypeOf(ctx.top[1]).String() == Decimal &&
				reflect.TypeOf(ctx.top[0]).String() == Decimal {
				if ctx.top[0].(decimal.Decimal).IsZero() {
					err = errDivZero
					ctx.isLoop = true
					return
				}
				ctx.bin = ctx.top[1].(decimal.Decimal).Mod(ctx.top[0].(decimal.Decimal)).Floor()
			} else {
				err = errUnsupportedType
				ctx.isLoop = true
				return
			}
		}
		return
	}
	instructionTable[compile.CmdAnd] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.bin = valueToBool(ctx.top[1]) && valueToBool(ctx.top[0])
		return
	}
	instructionTable[compile.CmdOr] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
		ctx.bin = valueToBool(ctx.top[1]) || valueToBool(ctx.top[0])

		return
	}
	for i := compile.CmdEqual; i <= compile.CmdNotEq; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			if ctx.top[1] == nil || ctx.top[0] == nil {
				ctx.bin = ctx.top[0] == ctx.top[1]
			} else {
				switch ctx.top[1].(type) {
				case string:
					switch ctx.top[0].(type) {
					case int64:
						if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
							ctx.bin = ctx.tmpInt == ctx.top[0].(int64)
						}
					case float64:
						ctx.bin = ValueToFloat(ctx.top[1]) == ctx.top[0].(float64)
					default:
						if reflect.TypeOf(ctx.top[0]).String() == Decimal {
							if ctx.tmpDec, err = ValueToDecimal(ctx.top[1]); err != nil {
								ctx.isLoop = true
								return
							}
							ctx.bin = ctx.tmpDec.Cmp(ctx.top[0].(decimal.Decimal)) == 0
						} else {
							ctx.bin = ctx.top[1].(string) == ctx.top[0].(string)
						}
					}
				case float64:
					ctx.bin = ctx.top[1].(float64) == ValueToFloat(ctx.top[0])
				case int64:
					switch ctx.top[0].(type) {
					case int64:
						ctx.bin = ctx.top[1].(int64) == ctx.top[0].(int64)
					case float64:
						ctx.bin = ValueToFloat(ctx.top[1]) == ctx.top[0].(float64)
					default:
						err = errUnsupportedType
						ctx.isLoop = true
						return
					}
				case bool:
					switch ctx.top[0].(type) {
					case bool:
						ctx.bin = ctx.top[1].(bool) == ctx.top[0].(bool)
					default:
						err = errUnsupportedType
						ctx.isLoop = true
						return
					}
				default:
					if ctx.tmpDec, err = ValueToDecimal(ctx.top[0]); err != nil {
						ctx.isLoop = true
						return
					}
					ctx.bin = ctx.top[1].(decimal.Decimal).Cmp(ctx.tmpDec) == 0
				}
			}
			if code.Cmd == compile.CmdNotEq {
				ctx.bin = !ctx.bin.(bool)
			}
			return
		}
	}
	for i := compile.CmdLess; i <= compile.CmdNotLess; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			switch ctx.top[1].(type) {
			case string:
				switch ctx.top[0].(type) {
				case int64:
					if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
						ctx.bin = ctx.tmpInt < ctx.top[0].(int64)
					}
				case float64:
					ctx.bin = ValueToFloat(ctx.top[1]) < ctx.top[0].(float64)
				default:
					if reflect.TypeOf(ctx.top[0]).String() == Decimal {
						if ctx.tmpDec, err = ValueToDecimal(ctx.top[1]); err != nil {
							ctx.isLoop = true
							return
						}
						ctx.bin = ctx.tmpDec.Cmp(ctx.top[0].(decimal.Decimal)) < 0
					} else {
						ctx.bin = ctx.top[1].(string) < ctx.top[0].(string)
					}
				}
			case float64:
				ctx.bin = ctx.top[1].(float64) < ValueToFloat(ctx.top[0])
			case int64:
				switch ctx.top[0].(type) {
				case int64:
					ctx.bin = ctx.top[1].(int64) < ctx.top[0].(int64)
				case float64:
					ctx.bin = ValueToFloat(ctx.top[1]) < ctx.top[0].(float64)
				default:
					err = errUnsupportedType
					ctx.isLoop = true
					return
				}
			default:
				if ctx.tmpDec, err = ValueToDecimal(ctx.top[0]); err != nil {
					ctx.isLoop = true
					return
				}
				ctx.bin = ctx.top[1].(decimal.Decimal).Cmp(ctx.tmpDec) < 0
			}
			if code.Cmd == compile.CmdNotLess {
				ctx.bin = !ctx.bin.(bool)
			}
			return
		}
	}
	for i := compile.CmdGreat; i <= compile.CmdNotGreat; i++ {
		instructionTable[i] = func(rt *Runtime, code *compile.ByteCode, ctx *instructionCtx) (status int, err error) {
			switch ctx.top[1].(type) {
			case string:
				switch ctx.top[0].(type) {
				case int64:
					if ctx.tmpInt, err = ValueToInt(ctx.top[1]); err == nil {
						ctx.bin = ctx.tmpInt > ctx.top[0].(int64)
					}
				case float64:
					ctx.bin = ValueToFloat(ctx.top[1]) > ctx.top[0].(float64)
				default:
					if reflect.TypeOf(ctx.top[0]).String() == Decimal {
						if ctx.tmpDec, err = ValueToDecimal(ctx.top[1]); err != nil {
							ctx.isLoop = true
							return
						}
						ctx.bin = ctx.tmpDec.Cmp(ctx.top[0].(decimal.Decimal)) > 0
					} else {
						ctx.bin = ctx.top[1].(string) > ctx.top[0].(string)
					}
				}
			case float64:
				ctx.bin = ctx.top[1].(float64) > ValueToFloat(ctx.top[0])
			case int64:
				switch ctx.top[0].(type) {
				case int64:
					ctx.bin = ctx.top[1].(int64) > ctx.top[0].(int64)
				case float64:
					ctx.bin = ValueToFloat(ctx.top[1]) > ctx.top[0].(float64)
				default:
					err = errUnsupportedType
					ctx.isLoop = true
					return
				}
			default:
				if ctx.tmpDec, err = ValueToDecimal(ctx.top[0]); err != nil {
					ctx.isLoop = true
					return
				}
				ctx.bin = ctx.top[1].(decimal.Decimal).Cmp(ctx.tmpDec) > 0
			}
			if code.Cmd == compile.CmdNotGreat {
				ctx.bin = !ctx.bin.(bool)
			}
			return
		}
	}
}
