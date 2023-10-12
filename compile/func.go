package compile

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// identifierRegexp = letter { letter | unicode_digit }
// letter        = unicode_letter | "_" .
// unicode_letter = /* a Unicode code point classified as "Letter" */
// unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */
var identifierRegexp = `^[\p{L}][\p{L}\p{Nd}_]*$`

func canIdent(ident string) error {
	if !regexp.MustCompile(identifierRegexp).MatchString(ident) {
		var val = ident
		if len(val) > 20 {
			val = val[:20] + "..."
		}
		return fmt.Errorf("invalid identifier: %s", val)
	}
	return nil
}

type compileFunc func(*CodeBlocks, stateType, *Lexeme) error

func fnNothing(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	return nil
}

func fnError(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	err := errTable[int(state)]
	if lexeme.Type == NEWLINE {
		return fmt.Errorf(`%s (unexpected new line) [Ln:%d]`, err, lexeme.Line-1)
	}
	return fmt.Errorf(`%s %s %v [%s]`, err, lexeme.Type, lexeme.Value, lexeme.Position())
}

// StateName checks the name of the contract and modifies it to @[state]name if it is necessary.
func StateName(state uint32, name string) string {
	if !strings.HasPrefix(name, `@`) {
		return fmt.Sprintf(`@%d%s`, state, name)
	} else if len(name) > 1 && (name[1] < '0' || name[1] > '9') {
		name = `@1` + name[1:]
	}
	return name
}

// fnBlockDecl is the function for the block declaration.
func fnBlockDecl(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	var itype ObjectType
	prev := (*buf)[len(*buf)-2]
	fblock := buf.peek()
	name := lexeme.Value.(string)
	switch state {
	case stateBlock:
		itype = ObjContract
		name = StateName(buf.ParentOwner().StateID, name)
		fblock.Info = &ContractInfo{ID: uint32(len(prev.Children) - 1), Name: name, Owner: buf.ParentOwner(), Used: make(map[string]bool)}
	default:
		itype = ObjFunc
		fblock.Info = &FuncInfo{ID: uint32(len(prev.Children) - 1), Name: name}
	}
	fblock.Type = itype
	if _, ok := prev.Objects[name]; ok {
		return fmt.Errorf("%s '%s' redeclared in this code block", itype, name)
	}
	prev.Objects[name] = NewObjInfo(itype, fblock)
	return nil
}

func fnFuncResult(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	fn := buf.peek().GetFuncInfo()
	(*fn).Results = append((*fn).Results, lexeme.Value.(reflect.Type))
	return nil
}

func fnReturn(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdReturn, lexeme, 0))
	return nil
}

func fnCmdError(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdError, lexeme, lexeme.Value))
	return nil
}

// fnParamName adds a new parameter name to the function or variable.
func fnParamName(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	name := lexeme.Value.(string)
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

	block.Objects[name] = NewObjInfo(ObjVar, &ObjInfoVariable{Name: name, Index: len(block.Vars)})
	block.Vars = append(block.Vars, reflect.TypeOf(nil))

	if block.Type != ObjFunc || (state != stateFnParam && state != stateFnParamType) {
		return nil
	}

	fblock := block.GetFuncInfo()
	if !fblock.HasTails() {
		fblock.Params = append(fblock.Params, reflect.TypeOf(nil))
		return nil
	}
	for name, f := range fblock.Tails {
		if f.Decl {
			continue
		}
		params := append(fblock.Tails[name].Params, reflect.TypeOf(nil))
		offset := append(fblock.Tails[name].Offset, len(block.Vars)-1)
		fblock.Tails[name] = FuncTail{Name: f.Name, Params: params, Offset: offset}
	}
	return nil
}

// fnParamType sets the type of the function parameter or variable.
func fnParamType(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	block := buf.peek()
	rtp := lexeme.Value.(reflect.Type)
	for i, v := range block.Vars {
		if v == reflect.TypeOf(nil) {
			block.Vars[i] = rtp
		}
	}
	if block.Type != ObjFunc || state != stateFnParam {
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

	for key, f := range fblock.Tails {
		if f.Decl {
			continue
		}
		for pkey, param := range fblock.Tails[key].Params {
			if param == reflect.TypeOf(nil) {
				fblock.Tails[key].Params[pkey] = rtp
			}
		}
	}

	return nil
}

// fnDeclTail the name of the tail function.
func fnDeclTail(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	if err := canIdent(lexeme.Value.(string)); err != nil {
		return err
	}

	fblock := buf.peek().GetFuncInfo()
	if !fblock.HasTails() {
		fblock.Tails = make(map[string]FuncTail)
	}
	if _, ok := fblock.Tails[lexeme.Value.(string)]; ok {
		return fmt.Errorf("tail func redeclared '%s'", lexeme.Value)
	}
	for k, name := range fblock.Tails {
		fblock.Tails[k] = FuncTail{
			Name:     name.Name,
			Params:   name.Params,
			Offset:   name.Offset,
			Variadic: name.Variadic,
			Decl:     true,
		}
	}
	fblock.Tails[lexeme.Value.(string)] = FuncTail{
		Name:   lexeme.Value.(string),
		Params: make([]reflect.Type, 0),
		Offset: make([]int, 0),
	}
	return nil
}

// fnTailParamType sets the type of the function final parameter.
func fnTailParamType(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
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
		if f.Decl {
			continue
		}
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

func fnIf(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdIf, lexeme, buf.peek()))
	return nil
}

func fnWhile(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdWhile, lexeme, buf.peek()))
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdContinue, lexeme, 0))
	return nil
}

func fnContinue(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdContinue, lexeme, 0))
	return nil
}

func fnBreak(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdBreak, lexeme, 0))
	return nil
}

func fnAssignVar(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	block := buf.peek()
	var (
		prev []*VarInfo
		ivar VarInfo
	)
	if lexeme.Type == EXTEND {
		if buf.get(0).AssertVar(lexeme.Value.(string)) {
			return fmt.Errorf(eSysVar, lexeme.Value.(string))
		}
		obj := NewObjInfo(ObjExtVar, &ObjInfoExtendVariable{Name: lexeme.Value.(string)})
		ivar = VarInfo{Obj: obj, Owner: nil}
	} else {
		objInfo, tobj := findVar(lexeme.Value.(string), buf)
		if objInfo == nil || objInfo.Type != ObjVar {
			return fmt.Errorf(`unknown variable %s`, lexeme.Value.(string))
		}
		ivar = VarInfo{Obj: objInfo, Owner: tobj}
	}
	if len(block.Code) > 0 {
		if block.Code[len(block.Code)-1].Cmd == CmdAssignVar {
			prev = block.Code[len(block.Code)-1].Value.([]*VarInfo)
		}
	}
	prev = append(prev, &ivar)
	if len(prev) == 1 {
		block.Code.push(newByteCode(CmdAssignVar, lexeme, prev))
	} else {
		block.Code[len(block.Code)-1] = newByteCode(CmdAssignVar, lexeme, prev)
	}
	return nil
}

func fnAssign(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	if lexeme.Type != OPERATOR {
		buf.peek().Code.push(newByteCode(CmdAssign, lexeme, 0))
	}
	return nil
}

func fnTx(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	contract := buf.peek()
	if contract.Type != ObjContract {
		return fmt.Errorf(`data can only be in contract`)
	}
	(*contract).GetContractInfo().Tx = new([]*FieldInfo)
	return nil
}

func fnSettings(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	contract := buf.peek()
	if contract.Type != ObjContract {
		return fmt.Errorf(`settings can only be in contract`)
	}
	(*contract).GetContractInfo().Settings = make(map[string]any)
	return nil
}

func fnConstName(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	sets := buf.peek().GetContractInfo().Settings
	sets[lexeme.Value.(string)] = nil
	return nil
}

func fnConstValue(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	sets := buf.peek().GetContractInfo().Settings
	for key, val := range sets {
		if val == nil {
			sets[key] = lexeme.Value
			break
		}
	}
	return nil
}

func fnField(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	info := buf.peek().GetContractInfo()
	tx := info.Tx
	if len(*tx) > 0 && (*tx)[len(*tx)-1].Type == reflect.TypeOf(nil) &&
		(*tx)[len(*tx)-1].Tags != `_` {
		return fmt.Errorf(eDataType, lexeme.Line, lexeme.Column)
	}

	if err := canIdent(lexeme.Value.(string)); err != nil {
		return err
	}
	if buf.get(0).AssertVar(lexeme.Value.(string)) {
		return fmt.Errorf(eDataParamVarCollides, lexeme.Value.(string), info.Name)
	}
	*tx = append(*tx, &FieldInfo{Name: lexeme.Value.(string), Type: reflect.TypeOf(nil)})
	return nil
}

func fnFields(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) > 0 && (*tx)[len(*tx)-1].Type == nil {
		return fmt.Errorf(eDataType, lexeme.Line, lexeme.Column)
	}
	return nil
}

func fnFieldComma(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) == 0 || (*tx)[len(*tx)-1].Type != nil {
		return fmt.Errorf(eDataName, lexeme.Line, lexeme.Column)
	}
	(*tx)[len(*tx)-1].Tags = `_`
	return nil
}

func fnFieldLine(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) > 0 && (*tx)[len(*tx)-1].Type == nil {
		return fmt.Errorf(eDataType, lexeme.Line, lexeme.Column)
	}
	for i, field := range *tx {
		if field.Tags == `_` {
			(*tx)[i].Tags = ``
		}
	}
	return nil
}

func fnFieldType(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) == 0 || (*tx)[len(*tx)-1].Type != nil {
		return fmt.Errorf(eDataName, lexeme.Line, lexeme.Column)
	}
	for i, field := range *tx {
		if field.Type == reflect.TypeOf(nil) {
			(*tx)[i].Type = lexeme.Value.(reflect.Type)
			(*tx)[i].Original = TypeReflect2Token((*tx)[i].Type)
		}
	}
	return nil
}

func fnFieldTag(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	tx := buf.peek().GetContractInfo().Tx
	if len(*tx) == 0 || (*tx)[len(*tx)-1].Type == nil || len((*tx)[len(*tx)-1].Tags) != 0 {
		return fmt.Errorf(eDataTag, lexeme.Line, lexeme.Column)
	}
	for i := len(*tx) - 1; i >= 0; i-- {
		if i == len(*tx)-1 || (*tx)[i].Tags == `_` {
			(*tx)[i].Tags = lexeme.Value.(string)
			continue
		}
		break
	}
	return nil
}

func fnElse(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	if buf.get(len(*buf)-2).Code.peek().Cmd != CmdIf {
		return fmt.Errorf(`there is not if before %v [%s]`, lexeme.Type, lexeme.Position())
	}
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdElse, lexeme, buf.peek()))
	return nil
}
