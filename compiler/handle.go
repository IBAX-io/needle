package compiler

import (
	"fmt"
	"regexp"
	"strings"
)

// identifierRegexp is a regular expression that matches valid identifiers in the language.
// An identifier starts with a unicode letter and can be followed by any number of letters, unicode digits, or underscores.
var identifierRegexp = `^[\p{L}][\p{L}\p{Nd}_]*$`

// canIdent checks if a string is a valid identifier according to the identifierRegexp.
func canIdent(ident string) error {
	if !regexp.MustCompile(identifierRegexp).MatchString(ident) {
		val := ident
		if len(val) > 20 {
			val = val[:20] + "..."
		}
		return fmt.Errorf("invalid identifier: %s", val)
	}
	return nil
}

// compileFunc is a type that represents a function that can be used to compile code blocks.
// It takes a pointer to a CodeBlocks, a stateType, and a pointer to a Lexeme, and returns an error.
type compileFunc func(*CodeBlocks, stateType, *Lexeme) error

// handleNothing is a compileFunc that does nothing and always returns nil.
func handleNothing(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	return nil
}

// handleError is a compileFunc that returns an error based on the current state and lexeme.
// If the lexeme type is NEWLINE, it returns an error indicating an unexpected semicolon or newline.
// Otherwise, it returns an error indicating an unexpected lexeme value.
func handleError(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	err := errTable[int(state)]
	if lex.Type == NEWLINE {
		return fmt.Errorf("%s, unexpected semicolon or newline", err)
	}
	return fmt.Errorf("%s, got %v", err, lex.Value)
}

// StateName checks the name of the contract and modifies it to @[state]name if it is necessary.
// If the name does not start with '@', it returns the name prefixed with the state number.
// If the name starts with '@' and the second character is not a digit, it returns the name prefixed with the state number.
// Otherwise, it returns the name as is.
func StateName(state uint32, name string) string {
	if !strings.HasPrefix(name, `@`) {
		return fmt.Sprintf(`@%d%s`, state, name)
	} else if len(name) > 1 && (name[1] < '0' || name[1] > '9') {
		name = `@1` + name[1:]
	}
	return name
}

// handleBlockDecl is the function for the block declaration.
func handleBlockDecl(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	var info isCodeBlockInfo
	prev := (*buf)[len(*buf)-2]
	block := buf.peek()
	name := lex.GetString()
	switch state {
	case stateBlock:
		if prev.Type != CodeBlockOwner {
			return fmt.Errorf("%s can only be in owner", lex.Value)
		}
		name = StateName(buf.ParentOwner().StateId, name)
		info = &ContractInfo{
			Id:   uint32(len(prev.Children) - 1),
			Name: name, Owner: buf.ParentOwner(), Used: make(map[string]bool),
		}
	default:
		if prev.Type != CodeBlockContract && prev.Type != CodeBlockOwner {
			return fmt.Errorf("%s can only be in contract or owner", lex.Value)
		}
		info = &FunctionInfo{Id: uint32(len(prev.Children) - 1), Name: name}
	}
	block.SetInfo(info)
	if _, ok := prev.Objects[name]; ok {
		return fmt.Errorf("%s '%s' redeclared in this code block", block.Type, name)
	}
	prev.Objects[name] = NewObject(block)
	return nil
}

func handleFuncResult(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	fn := buf.peek().GetFunctionInfo()
	(*fn).Results = append((*fn).Results, lex.GetToken())
	return nil
}

func handleReturn(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdReturn, lex, ""))
	return nil
}

func handleCmdError(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdError, lex, lex.Interface()))
	return nil
}

// handleParamName adds a new parameter name to the function or variable.
func handleParamName(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	name := lex.GetString()
	if err := canIdent(name); err != nil {
		return err
	}
	block := buf.peek()
	if _, ok := block.Objects[name]; ok {
		if state == stateFnParamType {
			return fmt.Errorf("duplicate argument '%s'", name)
		}
		if state == stateVarType {
			return fmt.Errorf("'%s' redeclared in this code block", name)
		}
	}

	block.Objects[name] = NewObject(&ObjInfoVariable{Name: name, Index: len(block.Vars)})
	block.Vars = append(block.Vars, UNKNOWN)

	if block.Type != CodeBlockFunction || (state != stateFnParam && state != stateFnParamType) {
		return nil
	}

	info := block.GetFunctionInfo()
	if !info.HasTails() {
		info.Params = append(info.Params, UNKNOWN)
		return nil
	}
	for name, f := range info.Tails {
		params := append(info.Tails[name].Params, UNKNOWN)
		offset := append(info.Tails[name].Offset, len(block.Vars)-1)
		info.Tails[name] = FuncTail{Name: f.Name, Params: params, Offset: offset}
	}
	return nil
}

// handleParamType sets the type of the function parameter or variable.
func handleParamType(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	block := buf.peek()
	rtp := lex.GetToken()
	for i, v := range block.Vars {
		if v == UNKNOWN {
			block.Vars[i] = rtp
		}
	}
	if block.Type != CodeBlockFunction || state != stateFnParam {
		return nil
	}
	fblock := block.GetFunctionInfo()
	if !fblock.HasTails() {
		for pkey, param := range fblock.Params {
			if param == UNKNOWN {
				fblock.Params[pkey] = rtp
			}
		}
		return nil
	}

	for key := range fblock.Tails {
		for pkey, param := range fblock.Tails[key].Params {
			if param == UNKNOWN {
				fblock.Tails[key].Params[pkey] = rtp
			}
		}
	}

	return nil
}

// handleDeclTail the name of the tail function.
func handleDeclTail(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	if err := canIdent(lex.GetString()); err != nil {
		return err
	}

	info := buf.peek().GetFunctionInfo()
	if !info.HasTails() {
		info.Tails = make(map[string]FuncTail)
	}
	if _, ok := info.Tails[lex.GetString()]; ok {
		return fmt.Errorf("tail func redeclared '%s'", lex.Value)
	}
	for k, name := range info.Tails {
		info.Tails[k] = FuncTail{
			Name:     name.Name,
			Params:   name.Params,
			Offset:   name.Offset,
			Variadic: name.Variadic,
		}
	}
	info.Tails[lex.GetString()] = FuncTail{
		Name:   lex.GetString(),
		Params: make([]Token, 0),
		Offset: make([]int, 0),
	}
	return nil
}

// handleTailParamType sets the type of the function final parameter.
func handleTailParamType(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	block := buf.peek()
	fblock := block.GetFunctionInfo()
	for vkey, ivar := range block.Vars {
		if ivar == UNKNOWN {
			block.Vars[vkey] = ARRAY
		}
	}
	var used bool
	if !fblock.HasTails() {
		for pkey, param := range fblock.Params {
			if param == UNKNOWN {
				if used {
					return fmt.Errorf(`... parameter must be one`)
				}
				fblock.Params[pkey] = ARRAY
				used = true
			}
		}
		fblock.Variadic = true
		return nil
	}

	for name, f := range fblock.Tails {
		for pkey, param := range fblock.Tails[name].Params {
			if param == UNKNOWN {
				if used {
					return fmt.Errorf("... parameter must be one")
				}
				fblock.Tails[name].Params[pkey] = ARRAY
				used = true
			}
		}
		offset := append(fblock.Tails[name].Offset, len(block.Vars))
		fblock.Tails[name] = FuncTail{
			Name:     f.Name,
			Params:   fblock.Tails[name].Params,
			Offset:   offset,
			Variadic: true,
		}
	}
	return nil
}

func handleIf(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().SetInfo(&CodeBlockIfInfo{})
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdIf, lex, buf.peek()))
	return nil
}

func handleWhile(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().SetInfo(&CodeBlockWhileInfo{})
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdWhile, lex, buf.peek()))
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdContinue, lex, ""))
	return nil
}

func handleContinue(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdContinue, lex, ""))
	return nil
}

func handleBreak(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdBreak, lex, ""))
	return nil
}

func handleAssignVar(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	block := buf.peek()
	var (
		prev []*VarInfo
		ivar VarInfo
	)
	if lex.Type == EXTEND {
		if buf.get(0).AssertVar(lex.GetString()) {
			return fmt.Errorf(eSysVar, lex.GetString())
		}
		obj := NewObject(&ObjInfoExtendVariable{Name: lex.GetString()})
		ivar = VarInfo{Obj: obj, Owner: nil}
	} else {
		obj, owner := findVar(lex.GetString(), buf)
		if obj == nil || obj.Type != ObjVariable {
			return fmt.Errorf(`unknown variable %s`, lex.GetString())
		}
		ivar = VarInfo{Obj: obj, Owner: owner}
	}
	if len(block.Code) > 0 {
		if block.Code[len(block.Code)-1].Cmd == CmdAssignVar {
			prev = block.Code[len(block.Code)-1].VarInfos()
		}
	}
	prev = append(prev, &ivar)
	if len(prev) == 1 {
		block.Code.push(newByteCode(CmdAssignVar, lex, prev))
	} else {
		block.Code[len(block.Code)-1] = newByteCode(CmdAssignVar, lex, prev)
	}
	return nil
}

func handleAssign(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	if lex.Type != OPERATOR {
		buf.peek().Code.push(newByteCode(CmdAssign, lex, ""))
	} else {
		if !lex.GetToken().Contains([]Token{
			AddEq, SubEq, MulEq, DivEq,
			ModEq, LshEq, RshEq, AndEq, OrEq, XorEq, Inc, Dec,
		}) {
			return fmt.Errorf("evaluated but not used: %s", lex.Value)
		}
	}
	return nil
}

func handleNewField(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	contract := buf.peek()
	if contract.Type != CodeBlockContract {
		return fmt.Errorf(`data can only be in contract`)
	}
	(*contract).GetContractInfo().Field = new([]*FieldInfo)
	return nil
}

func handleSettings(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	contract := buf.peek()
	if contract.Type != CodeBlockContract {
		return fmt.Errorf(`settings can only be in contract`)
	}
	(*contract).GetContractInfo().Settings = make(map[string]any)
	return nil
}

func handleConstName(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	sets := buf.peek().GetContractInfo().Settings
	sets[lex.GetString()] = nil
	return nil
}

func handleConstValue(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	sets := buf.peek().GetContractInfo().Settings
	for key, val := range sets {
		if val == nil {
			sets[key] = lex.Interface()
			break
		}
	}
	return nil
}

func handleField(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	if err := canIdent(lex.GetString()); err != nil {
		return err
	}
	info := buf.peek().GetContractInfo()
	field := info.Field
	if len(*field) > 0 && (*field)[len(*field)-1].Type == UNKNOWN &&
		(*field)[len(*field)-1].Tags != `_` {
		return fmt.Errorf(eDataType)
	}

	if buf.get(0).AssertVar(lex.GetString()) {
		return fmt.Errorf(eDataParamVarCollides, lex.GetString(), info.Name)
	}
	*field = append(*field, &FieldInfo{Name: lex.GetString(), Type: UNKNOWN})
	return nil
}

func handleFields(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	field := buf.peek().GetContractInfo().Field
	if len(*field) > 0 && (*field)[len(*field)-1].Type == UNKNOWN {
		return fmt.Errorf(eDataType)
	}
	return nil
}

func handleFieldComma(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	field := buf.peek().GetContractInfo().Field
	if len(*field) == 0 || (*field)[len(*field)-1].Type != UNKNOWN {
		return fmt.Errorf(eDataName)
	}
	(*field)[len(*field)-1].Tags = `_`
	return nil
}

func handleFieldLine(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	field := buf.peek().GetContractInfo().Field
	if len(*field) > 0 && (*field)[len(*field)-1].Type == UNKNOWN {
		return fmt.Errorf(eDataType)
	}
	for i, fie := range *field {
		if fie.Tags == `_` {
			(*field)[i].Tags = ``
		}
	}
	return nil
}

func handleFieldType(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	field := buf.peek().GetContractInfo().Field
	if len(*field) == 0 || (*field)[len(*field)-1].Type != UNKNOWN {
		return fmt.Errorf(eDataName)
	}
	for i, fie := range *field {
		if fie.Type == UNKNOWN {
			(*field)[i].Type = lex.GetToken()
			(*field)[i].Original = lex.GetToken()
		}
	}
	return nil
}

func handleFieldTag(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	field := buf.peek().GetContractInfo().Field
	if len(*field) == 0 || (*field)[len(*field)-1].Type == UNKNOWN || len((*field)[len(*field)-1].Tags) != 0 {
		return fmt.Errorf(eDataTag)
	}
	for i := len(*field) - 1; i >= 0; i-- {
		if i == len(*field)-1 || (*field)[i].Tags == `_` {
			(*field)[i].Tags = lex.GetString()
			continue
		}
		break
	}
	return nil
}

func handleElse(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	if buf.get(len(*buf)-2).Code.peek().Cmd != CmdIf {
		return fmt.Errorf("there is not if before %v", lex.Type)
	}
	buf.peek().SetInfo(&CodeBlockElseInfo{})
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdElse, lex, buf.peek()))
	return nil
}
