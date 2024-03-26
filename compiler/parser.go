package compiler

import (
	"fmt"
)

// Parser is a parser for the compiler.
type Parser struct {
	parent, root *CodeBlock
	inputs       Lexemes
	conf         *CompConfig
	blocks       CodeBlocks
	scanner
}

// NewParser returns a new parser for the given lexemes and configuration.
func NewParser(lexemes Lexemes, conf *CompConfig) *Parser {
	p := &Parser{
		inputs: lexemes,
		conf:   conf,
	}
	p.init()
	return p
}

func (p *Parser) init() {
	if p.conf == nil {
		p.conf = &CompConfig{
			Func:   make([]ExtendFunc, 0),
			PreVar: make([]string, 0),
			Owner:  &OwnerInfo{StateId: 1},
		}
	}
	p.root = &CodeBlock{
		Objects:        make(map[string]*Object),
		Type:           ObjOwner,
		Info:           p.conf.Owner,
		PredeclaredVar: p.conf.PreVar,
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
	if p.blocks == nil {
		p.blocks = make(CodeBlocks, 1)
	}
	p.blocks[0] = p.root
}

// Parse parses the lexemes and returns the root code block.
// It handles different states and transitions between them based on the lexemes.
func (p *Parser) Parse() (*CodeBlock, error) {
	if len(p.inputs) == 0 {
		return p.root, nil
	}
	curState := stateRoot
	stateStack := make([]stateType, 0)
	bls := p.blocks
	fork := 0

	for ; p.i < len(p.inputs); p.i++ {
		p.set(p.i)
		st, found := stateTable[curState]
		if !found {
			st = stateTable[stateRoot]
		}
		lex := p.lex
		comps, ok := st[lex.Type]
		if !ok {
			comps = st[UNKNOWN]
		}

		nextState := comps.next & 0xff
		if comps.hasState(stateFlagFork) {
			fork = p.i
		}
		if comps.hasState(stateFlagToFork) {
			p.i = fork
			lex = p.inputs[fork]
			fork = 0
		}
		if comps.hasState(stateFlagStay) {
			curState = nextState
			p.i--
			continue
		}
		if nextState == stateEval {
			if comps.hasState(stateFlagLabel) {
				bls.peek().Code.push(newByteCode(CmdLabel, lex, ""))
			}

			curlen := len(bls.peek().Code)
			if err := p.parseEval(&bls); err != nil {
				return nil, fmt.Errorf("parse evaluate error: %s", err)
			}
			if comps.hasState(stateFlagMustEval) && curlen == len(bls.peek().Code) {
				return nil, p.syntaxError("there is not eval expression")
			}

			nextState = curState
		}
		if comps.hasState(stateFlagPush) {
			stateStack = append(stateStack, curState)
			parent := bls.peek()
			block := &CodeBlock{Objects: make(map[string]*Object), Parent: parent}
			parent.Children.push(block)
			bls.push(block)
		}

		if comps.hasState(stateFlagPop) {
			if len(stateStack) == 0 {
				return nil, handleError(&bls, errMustLBRACE, lex)
			}
			nextState, stateStack = stateStack[len(stateStack)-1], stateStack[:len(stateStack)-1]
			if len(bls) >= 2 {
				prev := bls.get(len(bls) - 2)
				if len(prev.Code) > 0 && (*prev).Code[len((*prev).Code)-1].Cmd == CmdContinue {
					(*prev).Code = (*prev).Code[:len((*prev).Code)-1]
					prev = bls.peek()
					(*prev).Code.push(newByteCode(CmdContinue, lex, ""))
				}
			}
			bls = bls[:len(bls)-1]
		}
		if comps.hasState(stateFlagToBlock) {
			nextState = stateBlock
		}
		if comps.hasState(stateFlagToBody) {
			nextState = stateBody
		}

		if err := comps.handle(&bls, nextState, lex); err != nil {
			return nil, lex.errorWrap(fmt.Errorf("func handles error: %s", err))
		}
		curState = nextState
	}
	if len(stateStack) > 0 {
		return nil, handleError(&bls, errMustRBRACE, p.inputs[len(p.inputs)-1])
	}
	for _, item := range p.root.Objects {
		if item.Type == ObjContract {
			if cond, ok := item.GetCodeBlock().Objects[`conditions`]; ok {
				if cond.Type == ObjFunction && cond.GetFunctionInfo().CanWrite && p.conf.IgnoreObj != IgnoreIdent {
					return nil, fmt.Errorf("%s %w", item.GetContractInfo().Name, errCondWrite)
				}
			}
		}
	}
	return p.root, nil
}

func (p *Parser) parseEval(block *CodeBlocks) error {
	var indexInfo *IndexInfo
	curBlock := block.peek()
	buf := make(ByteCodes, 0, 20)
	bc := make(ByteCodes, 0, 100)
	parcount := make([]int, 0, 20)
	setIndex := false
	noMap := false
	prevLex := &Lexeme{}
	p.ns = false
main:
	for ; p.i < len(p.inputs); p.i++ {
		p.set(p.i)
		if p.ns && p.lex.Type != NEWLINE {
			return p.syntaxErrorExpected("expected semicolon or newline")
		}
		var cmd *ByteCode
		var call bool
		if !noMap {
			if p.lex.Type == LBRACE {
				if prevLex.Type == IF {
					return p.syntaxError("if must be followed by a condition")
				}
				pMap, err := p.parseInitMap(block, false)
				if err != nil {
					return err
				}
				bc.push(newByteCode(CmdMapInit, p.lex, pMap))
				continue
			}
			if p.lex.Type == LBRACK {
				pArray, err := p.parseInitArray(block)
				if err != nil {
					return err
				}
				bc.push(newByteCode(CmdArrayInit, p.lex, pArray))
				continue
			}
		}
		noMap = false
		switch p.lex.Type {
		default:
		case DOT:
			if p.nextN(1).Type != IDENTIFIER {
				return p.syntaxError("must be the name of the dot")
			}
		case EQ:
			if prevLex.Type == LPAREN || prevLex.Type == COMMA {
				return p.syntaxErrorExpected("expected expression")
			}
		case COLON:
			err := p.parserSliceStmt(&buf, &bc)
			if err != nil {
				return err
			}
		case RBRACE, LBRACE:
			p.i--
			if prevLex.Type == COMMA || prevLex.Type == OPERATOR || prevLex.Type == DOT {
				return p.syntaxErrorWrap(errEndExp)
			}
			break main
		case NEWLINE:
			if p.prevN(1).Type == COMMA || p.prevN(1).Type == OPERATOR &&
				(p.prevN(1).Value != Inc && p.prevN(1).Value != Dec) {
				continue main
			}
			for k := len(buf) - 1; k >= 0; k-- {
				if buf[k].Cmd == CmdSys {
					continue main
				}
			}
			break main
		case LPAREN, LBRACK:
			buf.push(newByteCode(CmdSys, p.lex, uint16(0xff)))
		case COMMA:
			if p.nextN(1).Type == RBRACK || p.nextN(1).Type == RPAREN {
				return p.syntaxError("unexpected trailing comma")
			}

			if len(parcount) > 0 {
				parcount[len(parcount)-1]++
			}
			for len(buf) > 0 {
				prev := buf.peek()
				if prev.Cmd == CmdSys && prev.Value.(uint16) == 0xff {
					break
				}
				bc.push(prev)
				buf.pop()
			}
		case RPAREN:
			noMap = true
			for {
				if buf.empty() {
					return p.syntaxError(fmt.Sprintf("there is not %s paren pair", p.lex.Type))
				}
				prev := buf.pop()
				if prev.Cmd == CmdSys && prev.Value.(uint16) == 0xff && prev.Lexeme.Type == LPAREN {
					break
				}
				bc.push(prev)
			}
			if buf.empty() {
				continue
			}
			prev := buf.peek()
			if fn := prev.FuncTailCmd(); fn != nil {
				buf.pop()
				names := fn.FuncTail
				wantlen := len(names.Params)
				if names.Variadic {
					wantlen--
				}
				count := parcount[len(parcount)-1]
				if count != wantlen && (!names.Variadic || count < wantlen) {
					if names.Variadic {
						return p.syntaxErrorWrap(fmt.Errorf("%s: at least %d params, but %d given", names.Name, wantlen, count))
					}
					return p.syntaxErrorWrap(fmt.Errorf("%s: want %d params, but %d given", names.Name, wantlen, count))
				}
				fn.Count = count
				(*prev).Value = fn
				parcount = parcount[:len(parcount)-1]
				bc.push(prev)
			}
			prev = buf.peek()
			if prev.Cmd != CmdCall && prev.Cmd != CmdCallVariadic {
				continue
			}
			objInfo := prev.Object()

			if len(buf) > 1 && buf[len(buf)-2].Lexeme.Type == OPERATOR && objInfo.GetResultsLen() != 1 {
				return p.syntaxError(fmt.Sprintf("function %s must return one value", objInfo.GetName()))
			}

			if (objInfo.Type == ObjFunction && objInfo.GetFunctionInfo().CanWrite) ||
				(objInfo.Type == ObjExtFunc && objInfo.GetExtFuncInfo().CanWrite) {
				setWritable(block)
			}
			if objInfo.Type == ObjFunction && objInfo.GetFunctionInfo().HasTails() {
				if bc.empty() || bc[len(bc)-1].Cmd != CmdFuncTail {
					bc.push(newByteCode(CmdPush, p.lex, make(map[string][]any)))
				}
				if p.nextN(1).Type == DOT && p.nextN(2).Type == IDENTIFIER {
					names := prev.Object().GetFunctionInfo().Tails
					if v, ok := names[p.nextN(2).Value.(string)]; ok {
						buf.push(newByteCode(CmdFuncTail, p.nextN(2), &FuncTailCmd{FuncTail: v}))
						count := 0
						if p.nextN(4).Type != RPAREN {
							count++
						}
						parcount = append(parcount, count)
						p.i += 2
						break
					}
				}
			}
			count := parcount[len(parcount)-1]
			parcount = parcount[:len(parcount)-1]
			if fn := prev.Object().GetFunctionInfo(); fn != nil {
				var y int
				for x := len(buf) - 1; x >= 0; x-- {
					buf := buf[x]
					if buf.Cmd == CmdCallVariadic || buf.Cmd == CmdCall {
						y++
					}
					if y == 2 && p.nextN(1).Type == RPAREN {
						if prev.Object().GetResultsLen() == 0 {
							return p.syntaxErrorWrap(fmt.Errorf("%s used return value, but it has no return value", fn.Name))
						}
						parcount[len(parcount)-1] += len(fn.Results) - 1
						break
					}
				}
			}
			if extFn := prev.Object().GetExtFuncInfo(); extFn != nil {
				wantlen := len(extFn.Params) - extFn.AutoParamsCount()
				if extFn.Variadic {
					wantlen--
				}
				if count != wantlen && (!extFn.Variadic || count < wantlen) && p.conf.IgnoreObj != IgnoreIdent {
					return p.syntaxErrorWrap(fmt.Errorf(eWrongParams, extFn.Name, wantlen))
				}
				var y int
				for x := len(buf) - 1; x >= 0; x-- {
					buf := buf[x]
					if buf.Cmd == CmdCallVariadic || buf.Cmd == CmdCall {
						y++
					}
					if y == 2 && p.nextN(1).Type == RPAREN {
						if prev.Object().GetResultsLen() == 0 {
							return p.syntaxErrorWrap(fmt.Errorf("%s used return value, but it has no return value", extFn.Name))
						}
						parcount[len(parcount)-1] += prev.Object().GetResultsLen() - 1
						break
					}
				}
			}
			if prev.Cmd == CmdCallVariadic {
				bc.push(newByteCode(CmdPush, p.lex, count))
			}
			buf = buf[:len(buf)-1]
			bc.push(prev)
		case RBRACK:
			noMap = true
			for {
				if buf.empty() {
					return p.syntaxError(fmt.Sprintf("there is not %s brack pair", p.lex.Type))
				}
				prev := buf.pop()
				if prev.Cmd == CmdSys && prev.Lexeme.Type == LBRACK {
					break
				}
				bc.push(prev)
			}

			if prev := buf.peek(); prev.want(CmdGetIndex) {
				buf.pop()
				switch v := p.nextN(1).Type; v {
				default:
				case EQ:
					p.i++
					setIndex, noMap = true, false
					indexInfo = prev.IndexInfo()
					continue
				}
				bc.push(prev)
			}
			if p.nextN(1).Type == LBRACK {
				return p.syntaxErrorWrap(errMultiIndex)
			}
		case OPERATOR:
			op, ok := operatorPriority[p.lex.Value.(Token)]
			if !ok {
				return p.syntaxErrorWrap(fmt.Errorf("unknown operator %v", p.lex.Value))
			}
			if op.Cmd == CmdDec || op.Cmd == CmdInc {
				p.ns = true
			}
			if f := bc.peek(); f != nil && (f.Cmd == CmdCall || f.Cmd == CmdCallVariadic) {
				if f.Object().GetResultsLen() != 1 {
					return p.syntaxError(fmt.Sprintf("function %s must return one value", f.Object().GetName()))
				}
			}
			var prevType Token
			if p.i > 0 {
				prevType = p.prevN(1).Type
			}
			if (op.Cmd == CmdSub || op.Cmd == CmdAdd) && (p.i == 0 || (prevType != NUMBER && prevType != IDENTIFIER &&
				prevType != EXTEND && prevType != LITERAL && prevType != RBRACE &&
				prevType != RBRACK && prevType != RPAREN)) {
				if next := p.nextN(1); next.Type != NUMBER && next.Type != IDENTIFIER && next.Type != EXTEND {
					return next.errorPos("expected operand")
				}
				op.Cmd = CmdSign
				op.Priority = uint16(CmdUnary)
			} else if prevLex.Type == OPERATOR && op.Priority != uint16(CmdUnary) {
				return p.syntaxErrorWrap(errOper)
			} else if prevLex.Type == RPAREN && (op.Cmd == CmdDec || op.Cmd == CmdInc) {
				return p.syntaxError(fmt.Sprintf("unexpected %s at end of statement", p.lex.Value))
			}
			byteOper := newByteCode(op.Cmd, p.lex, op.Priority)
			for {
				if buf.empty() {
					buf.push(byteOper)
					break
				}
				prev := buf.peek()
				if prev.Value.(uint16) >= op.Priority && op.Priority != uint16(CmdUnary) && prev.Cmd != CmdSys {
					if prev.Value.(uint16) == uint16(CmdUnary) { // Right to left
						unary := len(buf) - 1
						for unary > 0 && buf[unary-1].Value.(uint16) == uint16(CmdUnary) {
							unary--
						}
						bc = append(bc, buf[unary:]...)
						buf = buf[:unary]
					} else {
						bc.push(prev)
						buf = buf[:len(buf)-1]
					}
				} else {
					buf.push(byteOper)
					break
				}
			}
		case NUMBER, LITERAL:
			if prevLex.Type == RPAREN {
				return p.syntaxErrorWrap(fmt.Errorf("unexpected %v at end of expression", p.lex.Value))
			}
			noMap = true
			cmd = newByteCode(CmdPush, p.lex, p.lex.Value)
		case EXTEND:
			noMap = true
			if p.i < len(p.inputs)-2 {
				if p.nextN(1).Type == LPAREN {
					count := 0
					if p.nextN(2).Type != RPAREN {
						count++
					}
					parcount = append(parcount, count)
					buf.push(newByteCode(CmdCallExtend, p.lex, p.lex.Value.(string)))
					call = true
				}
			}
			if !call {
				cmd = newByteCode(CmdExtend, p.lex, p.lex.Value.(string))
				if p.i < len(p.inputs)-1 && p.nextN(1).Type == LBRACK {
					buf.push(newByteCode(CmdGetIndex, p.lex, &IndexInfo{Extend: p.lex.Value.(string)}))
				}
			}
		case IDENTIFIER:
			noMap = true
			obj, owner := p.findObj(p.lex.Value.(string), block)
			if obj == nil && (p.conf.IgnoreObj != IgnoreIdent || p.i >= len(p.inputs)-2 || p.nextN(1).Type != LPAREN) {
				return p.syntaxErrorWrap(fmt.Errorf(eUnknownIdent, p.lex.Value))
			}
			if p.i < len(p.inputs)-2 {
				if p.nextN(1).Type == LPAREN {
					var (
						isContract  bool
						objContract *ContractInfo
					)
					if p.conf.IgnoreObj == IgnoreIdent && obj == nil {
						obj = &Object{Type: ObjContract}
					}
					if obj == nil || (obj.Type != ObjExtFunc && obj.Type != ObjFunction &&
						obj.Type != ObjContract) {
						return p.syntaxErrorWrap(fmt.Errorf("unknown function or contract %s", p.lex.Value))
					}

					if obj.Type == ObjContract {
						if obj.Value != nil {
							objContract = obj.GetContractInfo()
						}
						obj, owner = p.findObj(`ExecContract`, block)
						if obj == nil {
							return p.syntaxErrorWrap(fmt.Errorf(eUnknownIdent, p.lex.Value))
						}
						isContract = true
					}
					cmd := CmdCall
					if obj.GetVariadic() {
						cmd = CmdCallVariadic
					}
					count := 0
					if p.nextN(2).Type != RPAREN {
						count++
					}
					buf.push(newByteCode(cmd, p.lex, obj))
					if isContract {
						name := StateName(block.ParentOwner().StateId, p.lex.Value.(string))
						for j := len(*block) - 1; j >= 0; j-- {
							topBlock := (*block)[j]
							if topBlock.Type == ObjContract {
								if name == topBlock.GetContractInfo().Name {
									return p.syntaxErrorWrap(errRecursion)
								}
								topBlock.GetContractInfo().Used[name] = true
							}
						}
						if objContract != nil && objContract.CanWrite {
							setWritable(block)
						}
						bc.push(newByteCode(CmdPush, p.lex, name))
						if count == 0 {
							count = 2
							bc.push(newByteCode(CmdPush, p.lex, ""))
							bc.push(newByteCode(CmdPush, p.lex, ""))
						}

						count++
					}
					if p.lex.Value.(string) == "CallContract" {
						count++
						bc.push(newByteCode(CmdPush, p.lex, block.ParentOwner().StateId))
					}
					parcount = append(parcount, count)
					call = true
				}
				if p.nextN(1).Type == LBRACK {
					if obj == nil || obj.Type != ObjVariable {
						return p.syntaxErrorWrap(fmt.Errorf("unknown variable %v", p.lex.Value))
					}
					buf.push(newByteCode(CmdGetIndex, p.lex, &IndexInfo{
						VarOffset: obj.GetVariable().Index,
						Owner:     owner,
					}))
				}
			}
			if !call {
				if obj.Type != ObjVariable {
					return p.syntaxErrorWrap(fmt.Errorf("unknown variable %v", p.lex.Value))
				}
				cmd = newByteCode(CmdVar, p.lex, &VarInfo{Obj: obj, Owner: owner})
			}
		case TAIL:
			cmd = newByteCode(CmdUnwrapArr, p.lex, "")
		}
		if p.lex.Type != NEWLINE {
			prevLex = p.lex
		}
		if cmd != nil {
			bc.push(cmd)
		}
	}
	if prevLex.Type == OPERATOR && (prevLex.Value != Inc && prevLex.Value != Dec) {
		return prevLex.errorWrap(errEndExp)
	}

	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i].Cmd == CmdSys {
			return buf[i].Lexeme.errorPos(fmt.Sprintf("there is not close pair of %s", buf[i].Lexeme.Type))
		}
		bc.push(buf[i])
	}
	if setIndex {
		bc.push(newByteCode(CmdSetIndex, &Lexeme{Line: bc[len(bc)-1].Lexeme.Line}, indexInfo))
	}
	curBlock.Code = append(curBlock.Code, bc...)
	return nil
}

func (p *Parser) parserSliceStmt(buffer, bytecode *ByteCodes) error {
	low, high := -1, -1
	if p.prevN(1).Type == LBRACK {
		low = SliceLow
	}
	if p.nextN(1).Type == RBRACK {
		high = SliceHigh
	}
	if p.prevN(1).Type == NUMBER {
		if (buffer.peek().Cmd == CmdSign && p.prevN(3).Type == LBRACK) || p.prevN(2).Type == LBRACK {
			if _, ok := p.prevN(1).Value.(int64); !ok {
				return p.prevN(1).errorPos("slice index must be integer")
			}
			low = SliceLowNum
		}
	}

	if p.nextN(1).Type == NUMBER && p.nextN(2).Type == RBRACK {
		if _, ok := p.nextN(1).Value.(int64); !ok {
			return p.nextN(1).errorPos("slice index must be integer")
		}
		high = SliceHighNum
	}
	if (p.nextN(1).Value == Sub || p.nextN(1).Value == Add) && p.nextN(2).Type == NUMBER && p.nextN(3).Type == RBRACK {
		if _, ok := p.nextN(2).Value.(int64); !ok {
			return p.nextN(2).errorPos("slice index must be integer")
		}
		high = SliceHighNum
	}
	if low != -1 && high != -1 {
		bytecode.push(newByteCode(CmdSliceColon, p.lex, &SliceItem{Index: [2]int{low, high}}))
		return nil
	}
	return p.syntaxError("invalid colon syntax")
}

func (p *Parser) findObj(name string, block *CodeBlocks) (obj *Object, owner *CodeBlock) {
	statename := StateName(block.ParentOwner().StateId, name)
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

// parseInitValue handles the parsing of a value initialization from the lexemes.
func (p *Parser) parseInitValue(block *CodeBlocks) (value *MapItem, err error) {
	var (
		subArr []*MapItem
		subMap *Map
	)
	switch p.lex.Type {
	case LBRACK:
		subArr, err = p.parseInitArray(block)
		if err == nil {
			value = &MapItem{Type: MapArray, Value: subArr}
		}
	case LBRACE:
		subMap, err = p.parseInitMap(block, false)
		if err == nil {
			value = &MapItem{Type: MapMap, Value: subMap}
		}
	case EXTEND:
		value = &MapItem{Type: MapExtend, Value: p.lex.Value}
	case IDENTIFIER:
		objInfo, tobj := p.findObj(p.lex.Value.(string), block)
		if objInfo == nil {
			err = p.syntaxErrorWrap(fmt.Errorf(eUnknownIdent, p.lex.Value))
		} else {
			value = &MapItem{Type: MapVar, Value: &VarInfo{Obj: objInfo, Owner: tobj}}
		}
	case NUMBER, LITERAL:
		value = &MapItem{Type: MapConst, Value: p.lex.Value}
	default:
		err = p.syntaxErrorExpected("expected string, int value or variable")
	}
	return
}

// parseInitMap handles the parsing of a map initialization from the lexemes.
func (p *Parser) parseInitMap(block *CodeBlocks, oneItem bool) (*Map, error) {
	if !oneItem {
		p.i++
	}
	key := ``
	ret := NewMap()
	state := MustKey
main:
	for ; p.i < len(p.inputs); p.i++ {
		p.set(p.i)
		switch p.lex.Type {
		case NEWLINE:
			continue
		case RBRACE:
			if state == MustColon {
				return nil, p.syntaxErrorExpected("expected colon")
			}
			if state == MustValue {
				return nil, p.syntaxErrorExpected("expected string, int value or variable")
			}
			if state == MustKey {
				return nil, p.syntaxErrorExpected("expected string key")
			}
			break main
		case COMMA, RBRACK:
			if oneItem {
				p.i--
				return ret, nil
			}
		}
		switch state {
		case MustComma:
			if p.lex.Type != COMMA {
				return nil, p.syntaxErrorExpected("expected comma or ]")
			}
			state = MustKey
		case MustColon:
			if p.lex.Type != COLON {
				return nil, p.syntaxErrorExpected("expected colon")
			}
			state = MustValue
		case MustKey:
			switch p.lex.Type & 0xff {
			case IDENTIFIER, LITERAL:
				key = p.lex.Value.(string)
			case EXTEND:
				key = "$" + p.lex.Value.(string)
			case KEYWORD:
				for ikey, v := range KeywordValue {
					if v == p.lex.Type {
						key = ikey
						if v == FUNC && p.i < len(p.inputs)-1 && p.nextN(1).Type&0xff == IDENTIFIER {
							continue main
						}
						break
					}
				}
			default:
				return nil, p.syntaxErrorExpected("expected string key")
			}

			state = MustColon
		case MustValue:
			mapi, err := p.parseInitValue(block)
			if err != nil {
				return nil, err
			}
			ret.Set(key, mapi)
			state = MustComma
		}
	}
	if ret.IsEmpty() && (state == MustKey) {
		return nil, p.syntaxErrorExpected("expected string key")
	}
	if p.i == len(p.inputs) {
		return nil, p.syntaxErrorWrap(errUnclosedMap)
	}
	return ret, nil
}

// parseInitArray handles the parsing of an array initialization from the lexemes.
func (p *Parser) parseInitArray(block *CodeBlocks) ([]*MapItem, error) {
	p.i++
	ret := make([]*MapItem, 0)
	state := MustValue
main:
	for ; p.i < len(p.inputs); p.i++ {
		p.set(p.i)
		switch p.lex.Type {
		case NEWLINE:
			continue
		case RBRACK:
			break main
		}
		switch state {
		case MustComma:
			if p.lex.Type != COMMA {
				return nil, p.syntaxErrorExpected("expected comma or ]")
			}
			state = MustValue
		case MustValue:
			if p.i+1 < len(p.inputs) && p.nextN(1).Type == COLON {
				subMap, err := p.parseInitMap(block, true)
				if err != nil {
					return nil, err
				}
				ret = append(ret, &MapItem{Type: MapMap, Value: subMap})
			} else {
				arri, err := p.parseInitValue(block)
				if err != nil {
					return nil, err
				}
				ret = append(ret, arri)
			}
			state = MustComma
		}
	}
	if len(ret) > 0 && state == MustValue {
		return nil, p.syntaxErrorExpected("expected string, int value or variable")
	}
	if p.i == len(p.inputs) {
		return nil, p.syntaxErrorWrap(errUnclosedArray)
	}
	return ret, nil
}
