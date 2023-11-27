package compile

import (
	"fmt"
)

type Parser struct {
	parent, root *CodeBlock
	lexemes      Lexemes
	conf         *CompConfig
	Debug        bool
}

func NewParser(lexemes Lexemes, conf *CompConfig) *Parser {
	var p = &Parser{
		lexemes: lexemes,
		conf:    conf,
	}
	p.init()
	return p
}

func (p *Parser) init() {
	if p.conf == nil {
		p.conf = &CompConfig{
			Func:   make([]ExtendFunc, 0),
			PreVar: make([]string, 0),
			Owner:  &OwnerInfo{StateID: 1},
		}
	}
	p.root = &CodeBlock{
		Objects: make(map[string]*Object),
		Type:    p.conf.Owner.ObjectType(),
		Info:    p.conf.Owner,
		PreVar:  p.conf.PreVar,
	}
	p.parent = &CodeBlock{
		Objects: make(map[string]*Object),
	}
	for s, info := range p.conf.Objects {
		p.parent.Objects[s] = info
	}
	for s, info := range p.conf.MakeExtFunc() {
		p.root.Objects[s] = info
	}
}

func (p *Parser) Parse() (*CodeBlock, error) {
	if len(p.lexemes) == 0 {
		return p.root, nil
	}
	curState := stateRoot
	stateStack := make([]stateType, 0)
	blocks := make(CodeBlocks, 1)
	blocks[0] = p.root
	fork := 0

	for i := 0; i < len(p.lexemes); i++ {
		lexeme := p.lexemes[i]
		st, found := stateTable[curState]
		if !found {
			st = stateTable[stateRoot]
		}
		comps, ok := st[lexeme.Type]
		if !ok {
			comps = st[UNKNOWN]
		}

		nextState := comps.next & 0xff
		if comps.hasState(stateFlagFork) {
			fork = i
		}
		if comps.hasState(stateFlagToFork) {
			i, fork = fork, 0
			lexeme = p.lexemes[i]
		}
		if comps.hasState(stateFlagStay) {
			curState = nextState
			i--
			continue
		}
		if nextState == stateEval {
			if comps.hasState(stateFlagLabel) {
				blocks.peek().Code.push(newByteCode(CmdLabel, lexeme, 0))
			}

			curlen := len(blocks.peek().Code)
			if err := p.parserEval(&i, &blocks); err != nil {
				return nil, fmt.Errorf("parser eval: %s", err)
			}
			if comps.hasState(stateFlagMustEval) && curlen == len(blocks.peek().Code) {
				return nil, fmt.Errorf("there is not eval expression in %s", lexeme.Position())
			}

			nextState = curState
		}
		if comps.hasState(stateFlagPush) {
			stateStack = append(stateStack, curState)
			parent := blocks.peek()
			block := &CodeBlock{Objects: make(map[string]*Object), Parent: parent}
			parent.Children.push(block)
			blocks.push(block)
		}

		if comps.hasState(stateFlagPop) {
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
		if comps.hasState(stateFlagToBlock) {
			nextState = stateBlock
		}
		if comps.hasState(stateFlagToBody) {
			nextState = stateBody
		}

		if err := comps.fn(&blocks, nextState, lexeme); err != nil {
			return nil, fmt.Errorf("func handles: %s[%s]", err, lexeme.Position())
		}
		curState = nextState
	}
	if len(stateStack) > 0 {
		return nil, fnError(&blocks, errMustRBRACE, p.lexemes[len(p.lexemes)-1])
	}
	for _, item := range p.root.Objects {
		if item.Type == ObjContract {
			if cond, ok := item.GetCodeBlock().Objects[`conditions`]; ok {
				if cond.Type == ObjFunc && cond.GetFuncInfo().CanWrite && p.conf.IgnoreObj != IgnoreIdent {
					return nil, fmt.Errorf("%s %w", item.GetContractInfo().Name, errCondWrite)
				}
			}
		}
	}
	return p.root, nil
}

func (p *Parser) parserEval(ind *int, block *CodeBlocks) error {
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
	for ; i < len(p.lexemes); i++ {
		var cmd *ByteCode
		var call bool
		lexeme := p.lexemes[i]
		if !noMap {
			if lexeme.Type == LBRACE {
				if prevLex.Type == IF {
					return fmt.Errorf("if must be followed by a condition")
				}
				pMap, err := p.parseInitMap(&i, block, false)
				if err != nil {
					return fmt.Errorf("get init map: %w", err)
				}
				bytecode.push(newByteCode(CmdMapInit, lexeme, pMap))
				continue
			}
			if lexeme.Type == LBRACK {
				pArray, err := p.parseInitArray(&i, block)
				if err != nil {
					return fmt.Errorf("get init array: %w", err)
				}
				bytecode.push(newByteCode(CmdArrayInit, lexeme, pArray))
				continue
			}
		}
		noMap = false
		switch lexeme.Type {
		case COLON:
			err := p.sliceStmt(i, buffer, bytecode)
			if err != nil {
				return err
			}
		case RBRACE, LBRACE:
			i--
			if prevLex.Type == COMMA || prevLex.Type == OPERATOR {
				return errEndExp
			}
			break main
		case NEWLINE:
			if i > 0 && (p.lexemes[i-1].Type == COMMA || p.lexemes[i-1].Type == OPERATOR && (p.lexemes[i-1].Value != Inc && p.lexemes[i-1].Value != Dec)) {
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
				if i < len(p.lexemes)-4 && p.lexemes[i+1].Type == DOT {
					if p.lexemes[i+2].Type != IDENTIFIER {
						return fmt.Errorf(`must be the name of the tail`)
					}
					names := prev.Value.(*Object).GetFuncInfo().Tails
					if _, ok := names[p.lexemes[i+2].Value.(string)]; !ok {
						if i < len(p.lexemes)-5 && p.lexemes[i+3].Type == LPAREN {
							objInfo, _ := p.findObj(p.lexemes[i+2].Value.(string), block)
							if objInfo != nil && (objInfo.Type == ObjFunc || objInfo.Type == ObjExtFunc) {
								tail = newByteCode(CmdCall, lexeme, objInfo)
							}
						}
						if tail == nil {
							return fmt.Errorf(`unknown function tail '%v'`, p.lexemes[i+2].Value)
						}
					}
					if tail == nil {
						v, _ := names[p.lexemes[i+2].Value.(string)]
						buffer.push(newByteCode(CmdFuncTail, lexeme, FuncTailCmd{FuncTail: v}))
						count := 0
						if p.lexemes[i+4].Type != RPAREN {
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
				if count != wantlen && (!extFn.Variadic || count < wantlen) && p.conf.IgnoreObj != IgnoreIdent {
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
					if i < len(p.lexemes)-1 && p.lexemes[i+1].Type == EQ {
						i++
						setIndex = true
						indexInfo = prev.Value.(*IndexInfo)
						noMap = false
						continue
					}
					bytecode.push(prev)
				}
			}
			if p.lexemes[i+1].Type == LBRACK {
				return errMultiIndex
			}
		case OPERATOR:
			op, ok := operatorPriority[lexeme.Value.(Token)]
			if !ok {
				return fmt.Errorf(`unknown operator %v`, lexeme.Value)
			}
			var prevType Token
			if i > 0 {
				prevType = p.lexemes[i-1].Type
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
			if i < len(p.lexemes)-2 {
				if p.lexemes[i+1].Type == LPAREN {
					count := 0
					if p.lexemes[i+2].Type != RPAREN {
						count++
					}
					parcount = append(parcount, count)
					buffer.push(newByteCode(CmdCallExtend, lexeme, lexeme.Value.(string)))
					call = true
				}
			}
			if !call {
				cmd = newByteCode(CmdExtend, lexeme, lexeme.Value.(string))
				if i < len(p.lexemes)-1 && p.lexemes[i+1].Type == LBRACK {
					buffer.push(newByteCode(CmdGetIndex, lexeme, &IndexInfo{Extend: lexeme.Value.(string)}))
				}
			}
		case IDENTIFIER:
			noMap = true
			obj, owner := p.findObj(lexeme.Value.(string), block)
			if obj == nil && (p.conf.IgnoreObj != IgnoreIdent || i >= len(p.lexemes)-2 || p.lexemes[i+1].Type != LPAREN) {
				return fmt.Errorf(eUnknownIdent, fmt.Sprintf(`%s[%s]`, lexeme.Value, lexeme.Position()))
			}
			if i < len(p.lexemes)-2 {
				if p.lexemes[i+1].Type == LPAREN {
					var (
						isContract  bool
						objContract *ContractInfo
					)
					if p.conf.IgnoreObj == IgnoreIdent && obj == nil {
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
						obj, owner = p.findObj(`ExecContract`, block)
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
					if p.lexemes[i+2].Type != RPAREN {
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
					if lexeme.Value.(string) == "CallContract" {
						count++
						bytecode.push(newByteCode(CmdPush, lexeme, block.ParentOwner().StateID))
					}
					parcount = append(parcount, count)
					call = true
				}
				if p.lexemes[i+1].Type == LBRACK {
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

func (p *Parser) sliceStmt(i int, buffer, bytecode ByteCodes) error {
	lexeme := p.lexemes[i]
	var low, high = -1, -1
	if p.lexemes[i-1].Type == LBRACK {
		low = SliceLow
	}
	if p.lexemes[i+1].Type == RBRACK {
		high = SliceHigh
	}
	if p.lexemes[i-1].Type == NUMBER {
		if (buffer.peek().Cmd == CmdSign && p.lexemes[i-3].Type == LBRACK) || p.lexemes[i-2].Type == LBRACK {
			if _, ok := p.lexemes[i-1].Value.(int64); !ok {
				return fmt.Errorf("slice index must be integer")
			}
			low = SliceLowNum
		}
	}

	if p.lexemes[i+1].Type == NUMBER && p.lexemes[i+2].Type == RBRACK {
		if _, ok := p.lexemes[i+1].Value.(int64); !ok {
			return fmt.Errorf("slice index must be integer")
		}
		high = SliceHighNum
	}
	if (p.lexemes[i+1].Value == Sub || p.lexemes[i+1].Value == Add) && p.lexemes[i+2].Type == NUMBER && p.lexemes[i+3].Type == RBRACK {
		if _, ok := p.lexemes[i+2].Value.(int64); !ok {
			return fmt.Errorf("slice index must be integer")
		}
		high = SliceHighNum
	}
	if low != -1 && high != -1 {
		bytecode.push(newByteCode(CmdSliceColon, lexeme, &SliceItem{Index: [2]int{low, high}}))
	}
	return nil
}

func (p *Parser) findObj(name string, block *CodeBlocks) (obj *Object, owner *CodeBlock) {
	statename := StateName(block.ParentOwner().StateID, name)
	for _, n := range []string{name, statename} {
		if obj, owner = findVar(n, block); obj != nil {
			return
		}
	}
	for _, n := range []string{name, statename} {
		ret, ok := p.parent.Objects[n]
		if !ok {
			continue
		}
		return ret, p.parent
	}
	return
}

func (p *Parser) parseInitValue(ind *int, block *CodeBlocks) (value MapItem, err error) {
	var (
		subArr []MapItem
		subMap *Map
	)
	i := *ind
	lexeme := p.lexemes[i]

	switch lexeme.Type {
	case LBRACK:
		subArr, err = p.parseInitArray(&i, block)
		if err == nil {
			value = MapItem{Type: MapArray, Value: subArr}
		}
	case LBRACE:
		subMap, err = p.parseInitMap(&i, block, false)
		if err == nil {
			value = MapItem{Type: MapMap, Value: subMap}
		}
	case EXTEND:
		value = MapItem{Type: MapExtend, Value: lexeme.Value}
	case IDENTIFIER:
		objInfo, tobj := p.findObj(lexeme.Value.(string), block)
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

func (p *Parser) parseInitMap(ind *int, block *CodeBlocks, oneItem bool) (*Map, error) {
	var next int
	if !oneItem {
		next = 1
	}
	i := *ind + next
	key := ``
	ret := NewMap()
	state := MustKey
main:
	for ; i < len(p.lexemes); i++ {
		lexeme := p.lexemes[i]
		switch lexeme.Type {
		case NEWLINE:
			continue
		case RBRACE:
			if state == MustColon {
				return nil, fmt.Errorf("%w[%s]", errUnexpColon, lexeme.Position())
			}
			if state == MustValue {
				return nil, fmt.Errorf("%w[%s]", errUnexpValue, lexeme.Position())
			}
			if state == MustKey {
				return nil, fmt.Errorf("%w[%s]", errUnexpKey, lexeme.Position())
			}
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
				return nil, fmt.Errorf("%w[%s]", errUnexpComma, lexeme.Position())
			}
			state = MustKey
		case MustColon:
			if lexeme.Type != COLON {
				return nil, fmt.Errorf("%w[%s]", errUnexpColon, lexeme.Position())
			}
			state = MustValue
		case MustKey:
			switch lexeme.Type & 0xff {
			case IDENTIFIER, LITERAL:
				key = lexeme.Value.(string)
			case EXTEND:
				key = "$" + lexeme.Value.(string)
			case KEYWORD:
				for ikey, v := range KeywordValue {
					if Keyword2Str(v) == fmt.Sprint(lexeme.Value) {
						key = ikey
						if v == FUNC && i < len(p.lexemes)-1 && p.lexemes[i+1].Type&0xff == IDENTIFIER {
							continue main
						}
						break
					}
				}
			default:
				return nil, fmt.Errorf("%w[%s]", errUnexpKey, lexeme.Position())
			}

			state = MustColon
		case MustValue:
			mapi, err := p.parseInitValue(&i, block)
			if err != nil {
				return nil, err
			}
			ret.Set(key, mapi)
			state = MustComma
		}
	}
	if ret.IsEmpty() && (state == MustKey) {
		return nil, fmt.Errorf("%w[%s]", errUnexpKey, p.lexemes[i].Position())
	}
	if i == len(p.lexemes) {
		return nil, errUnclosedMap
	}
	*ind = i
	return ret, nil
}

func (p *Parser) parseInitArray(ind *int, block *CodeBlocks) ([]MapItem, error) {
	i := *ind + 1
	ret := make([]MapItem, 0)
	state := MustValue
main:
	for ; i < len(p.lexemes); i++ {
		lexeme := p.lexemes[i]
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
			if i+1 < len(p.lexemes) && p.lexemes[i+1].Type == COLON {
				subMap, err := p.parseInitMap(&i, block, true)
				if err != nil {
					return nil, err
				}
				ret = append(ret, MapItem{Type: MapMap, Value: subMap})
			} else {
				arri, err := p.parseInitValue(&i, block)
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
	if i == len(p.lexemes) {
		return nil, errUnclosedArray
	}
	*ind = i
	return ret, nil
}
