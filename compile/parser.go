package compile

import (
	"fmt"
)

func NewParser(lexemes Lexemes, ext *ExtendData) (*CodeBlock, error) {
	if ext == nil {
		ext = &ExtendData{
			Func:   make([]ExtendFunc, 0),
			PreVar: make([]string, 0),
			Info:   &OwnerInfo{StateID: 1},
		}
	}
	root := &CodeBlock{
		Objects: make(map[string]*Object),
		Type:    ext.Info.ObjectType(),
		Info:    ext.Info,
		PreVar:  ext.PreVar,
	}
	for s, info := range ext.Objects {
		root.Objects[s] = info
	}
	for s, info := range ext.MakeExtFunc() {
		root.Objects[s] = info
	}
	if len(lexemes) == 0 {
		return root, nil
	}

	curState := stateRoot
	stateStack := make([]stateType, 0)
	blocks := make(CodeBlocks, 1)
	blocks[0] = root
	fork := 0

	for i := 0; i < len(lexemes); i++ {
		lexeme := lexemes[i]
		comps, ok := stateTable[curState][lexeme.Type]
		if !ok {
			comps = stateTable[curState][0]
		}

		nextState := comps.next & 0xff
		if comps.hasState(stateFork) {
			fork = i
		}
		if comps.hasState(stateToFork) {
			i = fork
			fork = 0
			lexeme = lexemes[i]
		}
		if comps.hasState(stateStay) {
			curState = nextState
			i--
			continue
		}
		if nextState == stateEval {
			if comps.hasState(stateLabel) {
				blocks.peek().Code.push(newByteCode(CmdLabel, lexeme, 0))
			}

			curlen := len(blocks.peek().Code)
			if err := parserEval(&lexemes, &i, &blocks, ext.Extern); err != nil {
				return nil, fmt.Errorf("parser eval: %s", err)
			}
			if (comps.next&stateMustEval) > 0 && curlen == len(blocks.peek().Code) {
				return nil, fmt.Errorf("there is not eval expression")
			}

			nextState = curState
		}
		if comps.hasState(statePush) {
			stateStack = append(stateStack, curState)
			parent := blocks.peek()
			block := &CodeBlock{Objects: make(map[string]*Object), Parent: parent}
			parent.Children.push(block)
			blocks.push(block)
		}

		if comps.hasState(statePop) {
			if len(stateStack) == 0 {
				return nil, fnError(&blocks, errMustLBRACE, lexeme)
			}
			nextState, stateStack = stateStack[len(stateStack)-1], stateStack[:len(stateStack)-1]
			if len(blocks) >= 2 {
				prev := blocks.get(len(blocks) - 2)
				if len(prev.Code) > 0 && (*prev).Code[len((*prev).Code)-1].Cmd == CmdContinue {
					(*prev).Code = (*prev).Code[:len((*prev).Code)-1]
					prev = blocks.peek()
					(*prev).Code.push(newByteCode(CmdContinue, lexeme, 0))
				}
			}
			blocks = blocks[:len(blocks)-1]
		}
		if comps.hasState(stateToBlock) {
			nextState = stateBlock
		}
		if comps.hasState(stateToBody) {
			nextState = stateBody
		}

		if err := comps.fn(&blocks, nextState, lexeme); err != nil {
			return nil, fmt.Errorf("func handles: %s", err)
		}
		curState = nextState
	}
	if len(stateStack) > 0 {
		return nil, fnError(&blocks, errMustRBRACE, lexemes[len(lexemes)-1])
	}
	for _, item := range root.Objects {
		if item.Type == ObjContract {
			if cond, ok := item.GetCodeBlock().Objects[`conditions`]; ok {
				if cond.Type == ObjFunc && cond.GetFuncInfo().CanWrite {
					return nil, errCondWrite
				}
			}
		}
	}
	return root, nil
}

func parserEval(lexemes *Lexemes, ind *int, block *CodeBlocks, extern bool) error {
	var indexInfo *IndexInfo
	i := *ind
	curBlock := block.peek()

	buffer := make(ByteCodes, 0, 20)
	bytecode := make(ByteCodes, 0, 100)
	parcount := make([]int, 0, 20)
	setIndex := false
	noMap := false
	prevLex := &Lexeme{}
main:
	for ; i < len(*lexemes); i++ {
		var cmd *ByteCode
		var call bool
		lexeme := (*lexemes)[i]
		if !noMap {
			if lexeme.Type == LBRACE {
				pMap, err := getInitMap(lexemes, &i, block, false)
				if err != nil {
					return err
				}
				bytecode.push(newByteCode(CmdMapInit, lexeme, pMap))
				continue
			}
			if lexeme.Type == LBRACK {
				pArray, err := getInitArray(lexemes, &i, block)
				if err != nil {
					return err
				}
				bytecode.push(newByteCode(CmdArrayInit, lexeme, pArray))
				continue
			}
		}
		noMap = false
		switch lexeme.Type {
		case COLON:
			var low, high = -1, -1
			if (*lexemes)[i-1].Type == LBRACK {
				low = SliceLow
			}
			if (*lexemes)[i+1].Type == RBRACK {
				high = SliceHigh
			}
			if (*lexemes)[i-1].Type == NUMBER {
				if (buffer.peek().Cmd == CmdSign && (*lexemes)[i-3].Type == LBRACK) || (*lexemes)[i-2].Type == LBRACK {
					if _, ok := (*lexemes)[i-1].Value.(int64); !ok {
						return fmt.Errorf("slice index must be integer")
					}
					low = SliceLowNum
				}
			}

			if (*lexemes)[i+1].Type == NUMBER && (*lexemes)[i+2].Type == RBRACK {
				if _, ok := (*lexemes)[i+1].Value.(int64); !ok {
					return fmt.Errorf("slice index must be integer")
				}
				high = SliceHighNum
			}
			if ((*lexemes)[i+1].Value == Sub || (*lexemes)[i+1].Value == Add) && (*lexemes)[i+2].Type == NUMBER && (*lexemes)[i+3].Type == RBRACK {
				if _, ok := (*lexemes)[i+2].Value.(int64); !ok {
					return fmt.Errorf("slice index must be integer")
				}
				high = SliceHighNum
			}
			if low != -1 && high != -1 {
				bytecode.push(newByteCode(CmdSliceColon, lexeme, &SliceItem{Index: [2]int{low, high}}))
			}
		case RBRACE, LBRACE:
			i--
			if prevLex.Type == COMMA || prevLex.Type == OPERATOR {
				return errEndExp
			}
			break main
		case NEWLINE:
			if i > 0 && ((*lexemes)[i-1].Type == COMMA || (*lexemes)[i-1].Type == OPERATOR && ((*lexemes)[i-1].Value != Inc && (*lexemes)[i-1].Value != Dec)) {
				continue main
			}
			for k := len(buffer) - 1; k >= 0; k-- {
				if buffer[k].Cmd == CmdSys {
					continue main
				}
			}
			break main
		case LPAREN, LBRACK:
			buffer.push(newByteCode(CmdSys, lexeme, uint16(0xff)))
		case COMMA:
			if len(parcount) > 0 {
				parcount[len(parcount)-1]++
			}
			for len(buffer) > 0 {
				prev := buffer.peek()
				if prev.Cmd == CmdSys && prev.Value.(uint16) == 0xff {
					break
				}
				bytecode.push(prev)
				buffer = buffer[:len(buffer)-1]
			}
		case RPAREN:
			noMap = true
			for {
				if len(buffer) == 0 {
					return fmt.Errorf("%s: there is not pair", lexeme.Type)
				}
				prev := buffer.pop()
				if prev.Value.(uint16) == 0xff {
					break
				}
				bytecode.push(prev)
			}
			if len(buffer) == 0 {
				continue
			}

			if prev := buffer.peek(); prev.Cmd == CmdFuncTail {
				buffer = buffer[:len(buffer)-1]
				fn := prev.Value.(FuncTailCmd)
				names := fn.FuncTail
				wantlen := len(names.Params)
				if names.Variadic {
					wantlen--
				}
				count := parcount[len(parcount)-1]
				if count != wantlen && (!names.Variadic || count < wantlen) {
					if names.Variadic {
						return fmt.Errorf("%s: at least %d params, but %d given", names.Name, wantlen, count)
					}
					return fmt.Errorf("%s: want %d params, but %d given", names.Name, wantlen, count)
				}
				fn.Count = count
				(*prev).Value = fn
				parcount = parcount[:len(parcount)-1]
				bytecode.push(prev)
			}
			var tail *ByteCode
			prev := buffer.peek()
			if prev.Cmd != CmdCall && prev.Cmd != CmdCallVariadic {
				continue
			}

			objInfo := prev.Value.(*Object)
			if (objInfo.Type == ObjFunc && objInfo.GetFuncInfo().CanWrite) ||
				(objInfo.Type == ObjExtFunc && objInfo.GetExtFuncInfo().CanWrite) {
				setWritable(block)
			}
			if objInfo.Type == ObjFunc && objInfo.GetFuncInfo().HasTails() {
				if len(bytecode) == 0 || bytecode[len(bytecode)-1].Cmd != CmdFuncTail {
					bytecode.push(newByteCode(CmdPush, lexeme, make(map[string][]any)))
				}
				if i < len(*lexemes)-4 && (*lexemes)[i+1].Type == DOT {
					if (*lexemes)[i+2].Type != IDENTIFIER {
						return fmt.Errorf(`must be the name of the tail`)
					}
					names := prev.Value.(*Object).GetFuncInfo().Tails
					if _, ok := names[(*lexemes)[i+2].Value.(string)]; !ok {
						if i < len(*lexemes)-5 && (*lexemes)[i+3].Type == LPAREN {
							objInfo, _ := findObj((*lexemes)[i+2].Value.(string), block)
							if objInfo != nil && (objInfo.Type == ObjFunc || objInfo.Type == ObjExtFunc) {
								tail = newByteCode(CmdCall, lexeme, objInfo)
							}
						}
						if tail == nil {
							return fmt.Errorf(`unknown function tail '%v'`, (*lexemes)[i+2].Value)
						}
					}
					if tail == nil {
						v, _ := names[(*lexemes)[i+2].Value.(string)]
						buffer.push(newByteCode(CmdFuncTail, lexeme, FuncTailCmd{FuncTail: v}))
						count := 0
						if (*lexemes)[i+4].Type != RPAREN {
							count++
						}
						parcount = append(parcount, count)
						i += 2
						break
					}
				}
			}
			count := parcount[len(parcount)-1]
			parcount = parcount[:len(parcount)-1]
			if prev.Value.(*Object).Type == ObjExtFunc {
				extFn := prev.Value.(*Object).GetExtFuncInfo()
				wantlen := len(extFn.Params) - extFn.AutoParamsCount()
				if extFn.Variadic {
					wantlen--
				}
				if count != wantlen && (!extFn.Variadic || count < wantlen) {
					return fmt.Errorf(eWrongParams, extFn.Name, wantlen)
				}
			}
			if prev.Cmd == CmdCallVariadic {
				bytecode.push(newByteCode(CmdPush, lexeme, count))
			}
			buffer = buffer[:len(buffer)-1]
			bytecode.push(prev)
			if tail != nil {
				buffer.push(tail)
				parcount = append(parcount, 1)
				i += 2
			}
		case RBRACK:
			noMap = true
			for {
				if len(buffer) == 0 {
					return fmt.Errorf("%s: there is not pair", lexeme.Type)
				}
				prev := buffer.pop()
				if prev.Value.(uint16) == 0xff {
					break
				}
				bytecode.push(prev)
			}
			if len(buffer) > 0 {
				if prev := buffer.peek(); prev.Cmd == CmdGetIndex {
					buffer = buffer[:len(buffer)-1]
					if i < len(*lexemes)-1 && (*lexemes)[i+1].Type == EQ {
						i++
						setIndex = true
						indexInfo = prev.Value.(*IndexInfo)
						noMap = false
						continue
					}
					bytecode.push(prev)
				}
			}
			if (*lexemes)[i+1].Type == LBRACK {
				return errMultiIndex
			}
		case OPERATOR:
			op, ok := operator[lexeme.Value.(Token)]
			if !ok {
				return fmt.Errorf(`unknown operator %v`, lexeme.Value)
			}
			var prevType Token
			if i > 0 {
				prevType = (*lexemes)[i-1].Type
			}
			if (op.Cmd == CmdSub || op.Cmd == CmdAdd) && (i == 0 || (prevType != NUMBER && prevType != IDENTIFIER &&
				prevType != EXTEND && prevType != LITERAL && prevType != RBRACE &&
				prevType != RBRACK && prevType != RPAREN)) {
				op.Cmd = CmdSign
				op.Priority = uint16(CmdUnary)
			} else if prevLex.Type == OPERATOR && op.Priority != uint16(CmdUnary) {
				return errOper
			}
			byteOper := newByteCode(op.Cmd, lexeme, op.Priority)
			for {
				if len(buffer) == 0 {
					buffer.push(byteOper)
					break
				}
				prev := buffer[len(buffer)-1]
				if prev.Value.(uint16) >= op.Priority && op.Priority != uint16(CmdUnary) && prev.Cmd != CmdSys {
					if prev.Value.(uint16) == uint16(CmdUnary) { // Right to left
						unar := len(buffer) - 1
						for ; unar > 0 && buffer[unar-1].Value.(uint16) == uint16(CmdUnary); unar-- {
						}
						bytecode = append(bytecode, buffer[unar:]...)
						buffer = buffer[:unar]
					} else {
						bytecode.push(prev)
						buffer = buffer[:len(buffer)-1]
					}
				} else {
					buffer.push(byteOper)
					break
				}
			}
		case NUMBER, LITERAL:
			noMap = true
			cmd = newByteCode(CmdPush, lexeme, lexeme.Value)
		case EXTEND:
			noMap = true
			if i < len(*lexemes)-2 {
				if (*lexemes)[i+1].Type == LPAREN {
					count := 0
					if (*lexemes)[i+2].Type != RPAREN {
						count++
					}
					parcount = append(parcount, count)
					buffer.push(newByteCode(CmdCallExtend, lexeme, lexeme.Value.(string)))
					call = true
				}
			}
			if !call {
				cmd = newByteCode(CmdExtend, lexeme, lexeme.Value.(string))
				if i < len(*lexemes)-1 && (*lexemes)[i+1].Type == LBRACK {
					buffer.push(newByteCode(CmdGetIndex, lexeme, &IndexInfo{Extend: lexeme.Value.(string)}))
				}
			}
		case IDENTIFIER:
			noMap = true
			obj, owner := findObj(lexeme.Value.(string), block)
			if obj == nil && (!extern || i > *ind || i >= len(*lexemes)-2 || (*lexemes)[i+1].Type != LPAREN) {
				return fmt.Errorf(eUnknownIdent, fmt.Sprintf(`%s[%s]`, lexeme.Value, lexeme.Position()))
			}
			if i < len(*lexemes)-2 {
				if (*lexemes)[i+1].Type == LPAREN {
					var (
						isContract  bool
						objContract *ContractInfo
					)
					if extern && obj == nil {
						obj = &Object{Type: ObjContract}
					}
					if obj == nil || (obj.Type != ObjExtFunc && obj.Type != ObjFunc &&
						obj.Type != ObjContract) {
						return fmt.Errorf(`unknown %s %s`, lexeme.Type, lexeme.Value)
					}

					if obj.Type == ObjContract {
						if obj.Value != nil {
							objContract = obj.GetContractInfo()
						}
						obj, owner = findObj(`ExecContract`, block)
						if obj == nil {
							return fmt.Errorf(eUnknownIdent, fmt.Sprintf(`%s[%s]`, lexeme.Value, lexeme.Position()))
						}
						isContract = true
					}
					cmd := CmdCall
					if obj.GetVariadic() {
						cmd = CmdCallVariadic
					}
					count := 0
					if (*lexemes)[i+2].Type != RPAREN {
						count++
					}
					buffer.push(newByteCode(cmd, lexeme, obj))
					if isContract {
						name := StateName(block.ParentOwner().StateID, lexeme.Value.(string))
						for j := len(*block) - 1; j >= 0; j-- {
							topBlock := (*block)[j]
							if topBlock.Type == ObjContract {
								if name == topBlock.GetContractInfo().Name {
									return errRecursion
								}
								topBlock.GetContractInfo().Used[name] = true
							}
						}
						if objContract != nil && objContract.CanWrite {
							setWritable(block)
						}
						bytecode.push(newByteCode(CmdPush, lexeme, name))
						if count == 0 {
							count = 2
							bytecode.push(newByteCode(CmdPush, lexeme, ""))
							bytecode.push(newByteCode(CmdPush, lexeme, ""))
						}
						count++
					}
					parcount = append(parcount, count)
					call = true
				}
				if (*lexemes)[i+1].Type == LBRACK {
					if obj == nil || obj.Type != ObjVar {
						return fmt.Errorf(`unknown variable %v`, lexeme.Value)
					}
					buffer.push(newByteCode(CmdGetIndex, lexeme, &IndexInfo{VarOffset: obj.GetVariable().Index, Owner: owner}))
				}
			}
			if !call {
				if obj.Type != ObjVar {
					return fmt.Errorf(`unknown variable %v`, lexeme.Value)
				}
				cmd = newByteCode(CmdVar, lexeme, &VarInfo{Obj: obj, Owner: owner})
			}
		case TAIL:
			cmd = newByteCode(CmdUnwrapArr, lexeme, 0)
		}
		if lexeme.Type != NEWLINE {
			prevLex = lexeme
		}
		if lexeme.Type&0xff == KEYWORD && lexeme.Value.(string) == Keyword2Str(TAIL) {
			cmd = newByteCode(CmdUnwrapArr, lexeme, 0)
		}
		if cmd != nil {
			bytecode.push(cmd)
		}
	}
	*ind = i
	if prevLex.Type == OPERATOR && (prevLex.Value != Inc && prevLex.Value != Dec) {
		return errEndExp
	}
	for i := len(buffer) - 1; i >= 0; i-- {
		if buffer[i].Cmd == CmdSys {
			return fmt.Errorf("%s: there is not pair", buffer[i].Lexeme.Type)
		}
		bytecode.push(buffer[i])
	}
	if setIndex {
		bytecode.push(newByteCode(CmdSetIndex, &Lexeme{Line: bytecode[len(bytecode)-1].Lexeme.Line}, indexInfo))
	}
	curBlock.Code = append(curBlock.Code, bytecode...)
	return nil
}

func findObj(name string, block *CodeBlocks) (obj *Object, owner *CodeBlock) {
	statename := StateName(block.ParentOwner().StateID, name)
	for _, n := range []string{name, statename} {
		if obj, owner = findVar(n, block); obj != nil {
			return
		}
	}
	return
}

func getInitValue(lexemes *Lexemes, ind *int, block *CodeBlocks) (value MapItem, err error) {
	var (
		subArr []MapItem
		subMap *Map
	)
	i := *ind
	lexeme := (*lexemes)[i]

	switch lexeme.Type {
	case LBRACK:
		subArr, err = getInitArray(lexemes, &i, block)
		if err == nil {
			value = MapItem{Type: MapArray, Value: subArr}
		}
	case LBRACE:
		subMap, err = getInitMap(lexemes, &i, block, false)
		if err == nil {
			value = MapItem{Type: MapMap, Value: subMap}
		}
	case EXTEND:
		value = MapItem{Type: MapExtend, Value: lexeme.Value}
	case IDENTIFIER:
		objInfo, tobj := findObj(lexeme.Value.(string), block)
		if objInfo == nil {
			err = fmt.Errorf(eUnknownIdent, lexeme.Value)
		} else {
			value = MapItem{Type: MapVar, Value: &VarInfo{Obj: objInfo, Owner: tobj}}
		}
	case NUMBER, LITERAL:
		value = MapItem{Type: MapConst, Value: lexeme.Value}
	default:
		err = errUnexpValue
	}
	*ind = i
	return
}

func getInitMap(lexemes *Lexemes, ind *int, block *CodeBlocks, oneItem bool) (*Map, error) {
	var next int
	if !oneItem {
		next = 1
	}
	i := *ind + next
	key := ``
	ret := NewMap()
	state := MustKey
main:
	for ; i < len(*lexemes); i++ {
		lexeme := (*lexemes)[i]
		switch lexeme.Type {
		case NEWLINE:
			continue
		case RBRACE:
			break main
		case COMMA, RBRACK:
			if oneItem {
				*ind = i - 1
				return ret, nil
			}
		}
		switch state {
		case MustComma:
			if lexeme.Type != COMMA {
				return nil, errUnexpComma
			}
			state = MustKey
		case MustColon:
			if lexeme.Type != COLON {
				return nil, errUnexpColon
			}
			state = MustValue
		case MustKey:
			switch lexeme.Type & 0xff {
			case IDENTIFIER, LITERAL:
				key = lexeme.Value.(string)
			case EXTEND:
				key = `$` + lexeme.Value.(string)
			case KEYWORD:
				for ikey, v := range KeywordValue {
					if fmt.Sprint(v) == fmt.Sprint(lexeme.Value) {
						key = ikey
						if v == FUNC && i < len(*lexemes)-1 && (*lexemes)[i+1].Type&0xff == IDENTIFIER {
							continue main
						}
						break
					}
				}
			default:
				return nil, errUnexpKey
			}
			state = MustColon
		case MustValue:
			mapi, err := getInitValue(lexemes, &i, block)
			if err != nil {
				return nil, err
			}
			ret.Set(key, mapi)
			state = MustComma
		}
	}
	if ret.IsEmpty() && state == MustKey {
		return nil, errUnexpKey
	}
	if i == len(*lexemes) {
		return nil, errUnclosedMap
	}
	*ind = i
	return ret, nil
}

func getInitArray(lexemes *Lexemes, ind *int, block *CodeBlocks) ([]MapItem, error) {
	i := *ind + 1
	ret := make([]MapItem, 0)
	state := MustValue
main:
	for ; i < len(*lexemes); i++ {
		lexeme := (*lexemes)[i]
		switch lexeme.Type {
		case NEWLINE:
			continue
		case RBRACK:
			break main
		}
		switch state {
		case MustComma:
			if lexeme.Type != COMMA {
				return nil, errUnexpComma
			}
			state = MustValue
		case MustValue:
			if i+1 < len(*lexemes) && (*lexemes)[i+1].Type == COLON {
				subMap, err := getInitMap(lexemes, &i, block, true)
				if err != nil {
					return nil, err
				}
				ret = append(ret, MapItem{Type: MapMap, Value: subMap})
			} else {
				arri, err := getInitValue(lexemes, &i, block)
				if err != nil {
					return nil, err
				}
				ret = append(ret, arri)
			}
			state = MustComma
		}
	}
	if len(ret) > 0 && state == MustValue {
		return nil, errUnexpValue
	}
	if i == len(*lexemes) {
		return nil, errUnclosedArray
	}
	*ind = i
	return ret, nil
}
