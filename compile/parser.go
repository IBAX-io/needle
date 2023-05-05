package compile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var extern bool // extern is mode of compilation

func NewParser(lexemes Lexemes, owner *OwnerInfo) (*CodeBlock, error) {
	if owner == nil {
		owner = &OwnerInfo{StateID: 1}
	}
	root := NewCodeBlock(owner)
	//root := &CodeBlock{Objects: make(map[string]*ObjInfo), Info: owner}
	if len(lexemes) == 0 {
		return root, nil
	}

	curState := stateRoot
	stateStack := make([]stateType, 0)
	blocksStack := make(CodeBlocks, 1)
	blocksStack[0] = root
	fork := 0

	for i := 0; i < len(lexemes); i++ {
		var (
			newState compileState
			ok       bool
		)
		lexeme := lexemes[i]
		if newState, ok = stateTable[curState][lexeme.Type]; !ok {
			newState = stateTable[curState][0]
		}

		nextState := newState.newState & 0xff
		if newState.hasState(stateFork) {
			fork = i
		}
		if newState.hasState(stateToFork) {
			i = fork
			fork = 0
			lexeme = lexemes[i]
		}
		if newState.hasState(stateStay) {
			curState = nextState
			i--
			continue
		}
		if nextState == stateEval {
			if newState.hasState(stateLabel) {
				blocksStack.peek().Code.push(newByteCode(CmdLabel, lexeme, lexeme.Line, 0))
			}

			curlen := len(blocksStack.peek().Code)
			if err := parserEval(&lexemes, &i, &blocksStack); err != nil {
				return nil, fmt.Errorf("parser eval: %s", err)
			}
			if (newState.newState&stateMustEval) > 0 && curlen == len(blocksStack.peek().Code) {
				return nil, fmt.Errorf("there is not eval expression")
			}

			nextState = curState
		}
		if newState.hasState(statePush) {
			stateStack = append(stateStack, curState)
			top := blocksStack.peek()
			block := &CodeBlock{Objects: make(map[string]*ObjInfo), Parent: top}
			top.Children.push(block)
			blocksStack.push(block)
		}

		if newState.hasState(statePop) {
			if len(stateStack) == 0 {
				return nil, fnError(&blocksStack, errMustLCurly, lexeme)
			}
			nextState = stateStack[len(stateStack)-1]
			stateStack = stateStack[:len(stateStack)-1]
			if len(blocksStack) >= 2 {
				prev := blocksStack.get(len(blocksStack) - 2)
				if len(prev.Code) > 0 && (*prev).Code[len((*prev).Code)-1].Cmd == CmdContinue {
					(*prev).Code = (*prev).Code[:len((*prev).Code)-1]
					prev = blocksStack.peek()
					(*prev).Code.push(newByteCode(CmdContinue, lexeme, lexeme.Line, 0))
				}
			}
			blocksStack = blocksStack[:len(blocksStack)-1]
		}
		if newState.hasState(stateToBlock) {
			nextState = stateBlock
		}
		if newState.hasState(stateToBody) {
			nextState = stateBody
		}
		if err := newState.funcHandle(&blocksStack, nextState, lexeme); err != nil {
			lexeme.GetLogger().WithFields(log.Fields{"type": ParseError, "nextState": nextState, "err": err, "lex_value": lexeme.Value}).Errorf("func handles")
			return nil, fmt.Errorf("func handles: %s", err)
		}
		curState = nextState
	}
	if len(stateStack) > 0 {
		return nil, fnError(&blocksStack, errMustRCurly, lexemes[len(lexemes)-1])
	}
	for _, item := range root.Objects {
		if item.Type == ObjectType_Contract {
			if cond, ok := item.GetCodeBlock().Objects[`conditions`]; ok {
				if cond.Type == ObjectType_Func && cond.GetCodeBlock().GetFuncInfo().CanWrite {
					return nil, errCondWrite
				}
			}
		}
	}
	return root, nil
}

func parserEval(lexemes *Lexemes, ind *int, block *CodeBlocks) error {
	var indexInfo *IndexInfo
	i := *ind
	curBlock := block.peek()

	buffer := make(ByteCodes, 0, 20)
	bytecode := make(ByteCodes, 0, 100)
	parcount := make([]int, 0, 20)
	setIndex := false
	noMap := false
	prevLex := Token(0)
main:
	for ; i < len(*lexemes); i++ {
		var cmd *ByteCode
		var call bool
		lexeme := (*lexemes)[i]
		logger := lexeme.GetLogger()
		if !noMap {
			if lexeme.Type == LBRACE {
				pMap, err := getInitMap(lexemes, &i, block, false)
				if err != nil {
					return err
				}
				bytecode.push(newByteCode(CmdMapInit, lexeme, lexeme.Line, pMap))
				continue
			}
			if lexeme.Type == LBRACK {
				pArray, err := getInitArray(lexemes, &i, block)
				if err != nil {
					return err
				}
				bytecode.push(newByteCode(CmdArrayInit, lexeme, lexeme.Line, pArray))
				continue
			}
		}
		noMap = false
		switch lexeme.Type {
		case RBRACE, LBRACE:
			i--
			if prevLex == COMMA || prevLex == OPERATOR {
				return errEndExp
			}
			break main
		case NEWLINE:
			if i > 0 && ((*lexemes)[i-1].Type == COMMA || (*lexemes)[i-1].Type == OPERATOR) {
				continue main
			}
			for k := len(buffer) - 1; k >= 0; k-- {
				if buffer[k].Cmd == CmdSys {
					continue main
				}
			}
			break main
		case LPAREN:
			buffer.push(newByteCode(CmdSys, lexeme, lexeme.Line, uint16(0xff)))
		case LBRACK:
			buffer.push(newByteCode(CmdSys, lexeme, lexeme.Line, uint16(0xff)))
		case COMMA:
			if len(parcount) > 0 {
				parcount[len(parcount)-1]++
			}
			for len(buffer) > 0 {
				prev := buffer[len(buffer)-1]
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
				prev := buffer[len(buffer)-1]
				buffer = buffer[:len(buffer)-1]
				if prev.Value.(uint16) == 0xff {
					break
				}
				bytecode.push(prev)
			}
			if len(buffer) > 0 {
				if prev := buffer[len(buffer)-1]; prev.Cmd == CmdFuncName {
					buffer = buffer[:len(buffer)-1]
					(*prev).Value = FuncNameCmd{Name: prev.Value.(FuncNameCmd).Name,
						Count: parcount[len(parcount)-1]}
					parcount = parcount[:len(parcount)-1]
					bytecode.push(prev)
				}
				var tail *ByteCode
				if prev := buffer[len(buffer)-1]; prev.Cmd == CmdCall || prev.Cmd == CmdCallVariadic {
					objInfo := prev.Value.(*ObjInfo)
					if (objInfo.Type == ObjectType_Func && objInfo.GetCodeBlock().GetFuncInfo().CanWrite) ||
						(objInfo.Type == ObjectType_ExtFunc && objInfo.GetExtFuncInfo().CanWrite) {
						setWritable(block)
					}
					if objInfo.Type == ObjectType_Func && objInfo.GetCodeBlock().GetFuncInfo().Names != nil {
						if len(bytecode) == 0 || bytecode[len(bytecode)-1].Cmd != CmdFuncName {
							bytecode.push(newByteCode(CmdPush, lexeme, lexeme.Line, nil))
						}
						if i < len(*lexemes)-4 && (*lexemes)[i+1].Type == DOT {
							if (*lexemes)[i+2].Type != IDENTIFIER {
								log.WithFields(log.Fields{"type": ParseError}).Error("must be the name of the tail")
								return fmt.Errorf(`must be the name of the tail`)
							}
							names := prev.Value.(*ObjInfo).GetCodeBlock().GetFuncInfo().Names
							if _, ok := (*names)[(*lexemes)[i+2].Value.(string)]; !ok {
								if i < len(*lexemes)-5 && (*lexemes)[i+3].Type == LPAREN {
									objInfo, _ := findObj((*lexemes)[i+2].Value.(string), block)
									if objInfo != nil && (objInfo.Type == ObjectType_Func || objInfo.Type == ObjectType_ExtFunc) {
										tail = newByteCode(CmdCall, lexeme, lexeme.Line, objInfo)
									}
								}
								if tail == nil {
									log.WithFields(log.Fields{"type": ParseError, "tail": (*lexemes)[i+2].Value.(string)}).Error("unknown function tail")
									return fmt.Errorf(`unknown function tail '%s'`, (*lexemes)[i+2].Value.(string))
								}
							}
							if tail == nil {
								buffer.push(newByteCode(CmdFuncName, lexeme, lexeme.Line, FuncNameCmd{Name: (*lexemes)[i+2].Value.(string)}))
								count := 0
								if (*lexemes)[i+3].Type != RPAREN {
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
					if prev.Value.(*ObjInfo).Type == ObjectType_ExtFunc {
						var errtext string
						extinfo := prev.Value.(*ObjInfo).GetExtFuncInfo()
						wantlen := len(extinfo.Params)
						for _, v := range extinfo.Auto {
							if len(v) > 0 {
								wantlen--
							}
						}
						if count != wantlen && (!extinfo.Variadic || count < wantlen) {
							errtext = fmt.Sprintf(eWrongParams, extinfo.Name, wantlen)
							logger.WithFields(log.Fields{"error": errtext, "type": ParseError}).Error(errtext)
							return fmt.Errorf(errtext)
						}
					}
					if prev.Cmd == CmdCallVariadic {
						bytecode.push(newByteCode(CmdPush, lexeme, lexeme.Line, count))
					}
					buffer = buffer[:len(buffer)-1]
					bytecode.push(prev)
					if tail != nil {
						buffer.push(tail)
						parcount = append(parcount, 1)
						i += 2
					}
				}
			}
		case RBRACK:
			noMap = true
			for {
				if len(buffer) == 0 {
					return fmt.Errorf("%s: there is not pair", lexeme.Type)
				}
				prev := buffer[len(buffer)-1]
				buffer = buffer[:len(buffer)-1]
				if prev.Value.(uint16) == 0xff {
					break
				}
				bytecode.push(prev)
			}
			if len(buffer) > 0 {
				if prev := buffer[len(buffer)-1]; prev.Cmd == CmdIndex {
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
			if op.Cmd == CmdSub && (i == 0 || (prevType != NUMBER && prevType != IDENTIFIER &&
				prevType != EXTEND && prevType != LITERAL && prevType != RBRACE &&
				prevType != RBRACK && prevType != RPAREN)) {
				op.Cmd = CmdSign
				op.Priority = uint16(CmdUnary)
			} else if prevLex == OPERATOR && op.Priority != uint16(CmdUnary) {
				return errOper
			}
			//buffer is stack
			byteOper := newByteCode(op.Cmd, lexeme, lexeme.Line, op.Priority)
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
			cmd = newByteCode(CmdPush, lexeme, lexeme.Line, lexeme.Value)
		case EXTEND:
			noMap = true
			if i < len(*lexemes)-2 {
				if (*lexemes)[i+1].Type == LPAREN {
					count := 0
					if (*lexemes)[i+2].Type != RPAREN {
						count++
					}
					parcount = append(parcount, count)
					buffer.push(newByteCode(CmdCallExtend, lexeme, lexeme.Line, lexeme.Value.(string)))
					call = true
				}
			}
			if !call {
				cmd = newByteCode(CmdExtend, lexeme, lexeme.Line, lexeme.Value.(string))
				if i < len(*lexemes)-1 && (*lexemes)[i+1].Type == LBRACK {
					buffer.push(newByteCode(CmdIndex, lexeme, lexeme.Line, &IndexInfo{Extend: lexeme.Value.(string)}))
				}
			}
		case IDENTIFIER:
			noMap = true
			objInfo, tobj := findObj(lexeme.Value.(string), block)
			if objInfo == nil && (!extern || i > *ind || i >= len(*lexemes)-2 || (*lexemes)[i+1].Type != LPAREN) {
				return fmt.Errorf(eUnknownIdent, fmt.Sprintf(`%s[%d:%d]`, lexeme.Value, lexeme.Line, lexeme.Column))
			}
			if i < len(*lexemes)-2 {
				if (*lexemes)[i+1].Type == LPAREN {
					var (
						isContract  bool
						objContract *CodeBlock
					)
					if extern && objInfo == nil {
						objInfo = &ObjInfo{Type: ObjectType_Contract}
					}

					if objInfo == nil || (objInfo.Type != ObjectType_ExtFunc && objInfo.Type != ObjectType_Func &&
						objInfo.Type != ObjectType_Contract) {
						logger.WithFields(log.Fields{"lex_value": lexeme.Value, "type": ParseError}).Error("unknown function")
						return fmt.Errorf(`unknown %s %s`, lexeme.Type, lexeme.Value)
					}

					if objInfo.Type == ObjectType_Contract {
						if objInfo.Value != nil {
							objContract = objInfo.GetCodeBlock()
						}
						objInfo, tobj = findObj(`ExecContract`, block)
						if objInfo == nil {
							return fmt.Errorf(eUnknownIdent, fmt.Sprintf(`%s[%d:%d]`, lexeme.Value, lexeme.Line, lexeme.Column))
						}
						isContract = true
					}
					cmd := CmdCall
					if (objInfo.Type == ObjectType_ExtFunc && objInfo.GetExtFuncInfo().Variadic) ||
						(objInfo.Type == ObjectType_Func && objInfo.GetCodeBlock().GetFuncInfo().Variadic) {
						cmd = CmdCallVariadic
					}
					count := 0
					if (*lexemes)[i+2].Type != RPAREN {
						count++
					}
					buffer.push(newByteCode(cmd, lexeme, lexeme.Line, objInfo))
					if isContract {
						name := StateName(block.ParentOwner().StateID, lexeme.Value.(string))
						for j := len(*block) - 1; j >= 0; j-- {
							topblock := (*block)[j]
							if topblock.Type == ObjectType_Contract {
								if name == topblock.GetContractInfo().Name {
									return errRecursion
								}
								topblock.GetContractInfo().Used[name] = true
							}
						}
						if objContract != nil && objContract.GetContractInfo().CanWrite {
							setWritable(block)
						}
						bytecode.push(newByteCode(CmdPush, lexeme, lexeme.Line, name))
						if count == 0 {
							count = 2
							bytecode.push(newByteCode(CmdPush, lexeme, lexeme.Line, ""))
							bytecode.push(newByteCode(CmdPush, lexeme, lexeme.Line, ""))
						}
						count++
					}
					if lexeme.Value.(string) == `CallContract` {
						count++
						bytecode.push(newByteCode(CmdPush, lexeme, lexeme.Line, block.ParentOwner().StateID))
					}
					parcount = append(parcount, count)
					call = true
				}
				if (*lexemes)[i+1].Type == LBRACK {
					if objInfo == nil || objInfo.Type != ObjectType_Var {
						logger.WithFields(log.Fields{"lex_value": lexeme.Value, "type": ParseError}).Error("unknown variable")
						return fmt.Errorf(`unknown variable %s`, lexeme.Value.(string))
					}
					buffer.push(newByteCode(CmdIndex, lexeme, lexeme.Line, &IndexInfo{VarOffset: objInfo.GetVariable().Index, Owner: tobj}))
				}
			}
			if !call {
				if objInfo.Type != ObjectType_Var {
					return fmt.Errorf(`unknown variable %s`, lexeme.Value.(string))
				}
				cmd = newByteCode(CmdVar, lexeme, lexeme.Line, &VarInfo{Obj: objInfo, Owner: tobj})
			}
		default:
			//fmt.Println("other-", lexeme.Type, lexeme.Value)
		}
		if lexeme.Type != NEWLINE {
			prevLex = lexeme.Type
		}
		if lexeme.Type&0xff == KEYWORD {
			if lexeme.Value.(Token) == TAIL {
				cmd = newByteCode(CmdUnwrapArr, lexeme, lexeme.Line, 0)
			}
		}
		if cmd != nil {
			bytecode.push(cmd)
		}
	}
	*ind = i
	if prevLex == OPERATOR {
		return errEndExp
	}
	for i := len(buffer) - 1; i >= 0; i-- {
		if buffer[i].Cmd == CmdSys {
			return fmt.Errorf("%s: there is not pair", buffer[i].Lexeme.Type)
		}
		bytecode.push(buffer[i])
	}
	if setIndex {
		bytecode.push(newByteCode(CmdSetIndex, nil, 0, indexInfo))
	}
	curBlock.Code = append(curBlock.Code, bytecode...)
	return nil
}

func findObj(name string, block *CodeBlocks) (obj *ObjInfo, owner *CodeBlock) {
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
						if v == keyFunc && i < len(*lexemes)-1 && (*lexemes)[i+1].Type&0xff == IDENTIFIER {
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
