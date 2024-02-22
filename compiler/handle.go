package compiler

import (
	"fmt"
	"reflect"
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
	var itype ObjectType
	prev := (*buf)[len(*buf)-2]
	fblock := buf.peek()
	name := lex.Value.(string)
	switch state {
	case stateBlock:
		if prev.Type != ObjOwner {
			return fmt.Errorf("%s can only be in owner", lex.Value)
		}
		itype = ObjContract
		name = StateName(buf.ParentOwner().StateId, name)
		fblock.Info = &ContractInfo{
			Id:   uint32(len(prev.Children) - 1),
			Name: name, Owner: buf.ParentOwner(), Used: make(map[string]bool),
		}
	default:
		if prev.Type != ObjContract && prev.Type != ObjOwner {
			return fmt.Errorf("%s can only be in contract or owner", lex.Value)
		}
		itype = ObjFunction
		fblock.Info = &FunctionInfo{Id: uint32(len(prev.Children) - 1), Name: name}
	}
	fblock.Type = itype
	if _, ok := prev.Objects[name]; ok {
		return fmt.Errorf("%s '%s' redeclared in this code block", itype, name)
	}
	prev.Objects[name] = NewObject(fblock)
	return nil
}

func handleFuncResult(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	fn := buf.peek().GetFuncInfo()
	(*fn).Results = append((*fn).Results, TypeNameReflect[lex.Value.(Token)])
	return nil
}

func handleReturn(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdReturn, lex, 0))
	return nil
}

func handleCmdError(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdError, lex, lex.Value))
	return nil
}

// handleParamName adds a new parameter name to the function or variable.
func handleParamName(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	name := lex.Value.(string)
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
	block.Vars = append(block.Vars, reflect.TypeOf(nil))

	if block.Type != ObjFunction || (state != stateFnParam && state != stateFnParamType) {
		return nil
	}

	fblock := block.GetFuncInfo()
	if !fblock.HasTails() {
		fblock.Params = append(fblock.Params, reflect.TypeOf(nil))
		return nil
	}
	for name, f := range fblock.Tails {
		params := append(fblock.Tails[name].Params, reflect.TypeOf(nil))
		offset := append(fblock.Tails[name].Offset, len(block.Vars)-1)
		fblock.Tails[name] = FuncTail{Name: f.Name, Params: params, Offset: offset}
	}
	return nil
}

// handleParamType sets the type of the function parameter or variable.
func handleParamType(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	block := buf.peek()
	rtp := TypeNameReflect[lex.Value.(Token)]
	for i, v := range block.Vars {
		if v == reflect.TypeOf(nil) {
			block.Vars[i] = rtp
		}
	}
	if block.Type != ObjFunction || state != stateFnParam {
		return nil
	}
	fblock := block.GetFuncInfo()
	if !fblock.HasTails() {
		for pkey, param := range fblock.Params {
			if param == reflect.TypeOf(nil) {
				fblock.Params[pkey] = rtp
			}
		}
		return nil
	}

	for key := range fblock.Tails {
		for pkey, param := range fblock.Tails[key].Params {
			if param == reflect.TypeOf(nil) {
				fblock.Tails[key].Params[pkey] = rtp
			}
		}
	}

	return nil
}

// handleDeclTail the name of the tail function.
func handleDeclTail(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	if err := canIdent(lex.Value.(string)); err != nil {
		return err
	}

	fblock := buf.peek().GetFuncInfo()
	if !fblock.HasTails() {
		fblock.Tails = make(map[string]FuncTail)
	}
	if _, ok := fblock.Tails[lex.Value.(string)]; ok {
		return fmt.Errorf("tail func redeclared '%s'", lex.Value)
	}
	for k, name := range fblock.Tails {
		fblock.Tails[k] = FuncTail{
			Name:     name.Name,
			Params:   name.Params,
			Offset:   name.Offset,
			Variadic: name.Variadic,
		}
	}
	fblock.Tails[lex.Value.(string)] = FuncTail{
		Name:   lex.Value.(string),
		Params: make([]reflect.Type, 0),
		Offset: make([]int, 0),
	}
	return nil
}

// handleTailParamType sets the type of the function final parameter.
func handleTailParamType(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	block := buf.peek()
	fblock := block.GetFuncInfo()
	for vkey, ivar := range block.Vars {
		if ivar == reflect.TypeOf(nil) {
			block.Vars[vkey] = reflect.TypeOf([]any{})
		}
	}
	var used bool
	if !fblock.HasTails() {
		for pkey, param := range fblock.Params {
			if param == reflect.TypeOf(nil) {
				if used {
					return fmt.Errorf(`... parameter must be one`)
				}
				fblock.Params[pkey] = reflect.TypeOf([]any{})
				used = true
			}
		}
		fblock.Variadic = true
		return nil
	}

	for name, f := range fblock.Tails {
		for pkey, param := range fblock.Tails[name].Params {
			if param == reflect.TypeOf(nil) {
				if used {
					return fmt.Errorf("... parameter must be one")
				}
				fblock.Tails[name].Params[pkey] = reflect.TypeOf([]any{})
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
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdIf, lex, buf.peek()))
	return nil
}

func handleWhile(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdWhile, lex, buf.peek()))
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdContinue, lex, 0))
	return nil
}

func handleContinue(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdContinue, lex, 0))
	return nil
}

func handleBreak(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdBreak, lex, 0))
	return nil
}

func handleAssignVar(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	block := buf.peek()
	var (
		prev []*VarInfo
		ivar VarInfo
	)
	if lex.Type == EXTEND {
		if buf.get(0).AssertVar(lex.Value.(string)) {
			return fmt.Errorf(eSysVar, lex.Value.(string))
		}
		obj := NewObject(&ObjInfoExtendVariable{Name: lex.Value.(string)})
		ivar = VarInfo{Obj: obj, Owner: nil}
	} else {
		objInfo, tobj := findVar(lex.Value.(string), buf)
		if objInfo == nil || objInfo.Type != ObjVariable {
			return fmt.Errorf(`unknown variable %s`, lex.Value.(string))
		}
		ivar = VarInfo{Obj: objInfo, Owner: tobj}
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
		buf.peek().Code.push(newByteCode(CmdAssign, lex, 0))
	} else {
		if !lex.Value.(Token).contains([]Token{
			AddEq, SubEq, MulEq, DivEq,
			ModEq, LshEq, RshEq, AndEq, OrEq, XorEq, Inc, Dec,
		}) {
			return fmt.Errorf("evaluated but not used: %s", lex.Value)
		}
	}
	return nil
}

func handleTx(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	contract := buf.peek()
	if contract.Type != ObjContract {
		return fmt.Errorf(`data can only be in contract`)
	}
	(*contract).GetContractInfo().Tx = new([]*FieldInfo)
	return nil
}

func handleSettings(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	contract := buf.peek()
	if contract.Type != ObjContract {
		return fmt.Errorf(`settings can only be in contract`)
	}
	(*contract).GetContractInfo().Settings = make(map[string]any)
	return nil
}

func handleConstName(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	sets := buf.peek().GetContractInfo().Settings
	sets[lex.Value.(string)] = nil
	return nil
}

func handleConstValue(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	sets := buf.peek().GetContractInfo().Settings
	for key, val := range sets {
		if val == nil {
			sets[key] = lex.Value
			break
		}
	}
	return nil
}

func handleField(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	info := buf.peek().GetContractInfo()
	tx := info.Tx
	if len(*tx) > 0 && (*tx)[len(*tx)-1].Type == reflect.TypeOf(nil) &&
		(*tx)[len(*tx)-1].Tags != `_` {
		return fmt.Errorf(eDataType)
	}

	if err := canIdent(lex.Value.(string)); err != nil {
		return err
	}
	if buf.get(0).AssertVar(lex.Value.(string)) {
		return fmt.Errorf(eDataParamVarCollides, lex.Value.(string), info.Name)
	}
	*tx = append(*tx, &FieldInfo{Name: lex.Value.(string), Type: reflect.TypeOf(nil)})
	return nil
}

func handleFields(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) > 0 && (*tx)[len(*tx)-1].Type == nil {
		return fmt.Errorf(eDataType)
	}
	return nil
}

func handleFieldComma(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) == 0 || (*tx)[len(*tx)-1].Type != nil {
		return fmt.Errorf(eDataName)
	}
	(*tx)[len(*tx)-1].Tags = `_`
	return nil
}

func handleFieldLine(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) > 0 && (*tx)[len(*tx)-1].Type == nil {
		return fmt.Errorf(eDataType)
	}
	for i, field := range *tx {
		if field.Tags == `_` {
			(*tx)[i].Tags = ``
		}
	}
	return nil
}

func handleFieldType(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) == 0 || (*tx)[len(*tx)-1].Type != nil {
		return fmt.Errorf(eDataName)
	}
	for i, field := range *tx {
		if field.Type == reflect.TypeOf(nil) {
			(*tx)[i].Type = TypeNameReflect[lex.Value.(Token)]
			(*tx)[i].Original = lex.Value.(Token)
		}
	}
	return nil
}

func handleFieldTag(buf *CodeBlocks, state stateType, lex *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) == 0 || (*tx)[len(*tx)-1].Type == nil || len((*tx)[len(*tx)-1].Tags) != 0 {
		return fmt.Errorf(eDataTag)
	}
	for i := len(*tx) - 1; i >= 0; i-- {
		if i == len(*tx)-1 || (*tx)[i].Tags == `_` {
			(*tx)[i].Tags = lex.Value.(string)
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
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdElse, lex, buf.peek()))
	return nil
}
