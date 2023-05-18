package compile

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type compileFunc func(*CodeBlocks, stateType, *Lexeme) error

// varRegexp letter { letter | unicode_digit }
var varRegexp = `^[a-zA-Z][a-zA-Z0-9_]*$`

func fnNothing(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	return nil
}

func fnError(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	err := errTable[int(state)]
	logger := lexeme.GetLogger()
	if lexeme.Type == NEWLINE {
		logger.WithFields(log.Fields{"error": err, "lex_value": lexeme.Value, "type": ParseError}).Error("unexpected new line")
		return fmt.Errorf(`%s (unexpected new line) [Ln:%d]`, err, lexeme.Line-1)
	}
	logger.WithFields(log.Fields{"error": err, "lex_value": lexeme.Value, "type": ParseError}).Error("parsing error")
	return fmt.Errorf(`%s %s %v [%d:%d]`, err, lexeme.Type, lexeme.Value, lexeme.Line, lexeme.Column)
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
func fnNameBlock(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	var itype ObjectType
	prev := (*buf)[len(*buf)-2]
	fblock := buf.peek()
	name := lexeme.Value.(string)
	switch state {
	case stateBlock:
		itype = ObjectType_Contract
		name = StateName(buf.ParentOwner().StateID, name)
		fblock.Info = &ContractInfo{ID: uint32(len(prev.Children) - 1), Name: name, Owner: buf.ParentOwner(), Used: make(map[string]bool)}
	default:
		itype = ObjectType_Func
		fblock.Info = &FuncInfo{Name: name}
	}
	fblock.Type = itype
	if obj, ok := prev.Objects[name]; ok {
		if obj.Type == ObjectType_Contract {
			return fmt.Errorf("%s '%s' redeclared in this code block", itype, name)
		}
		return fmt.Errorf("%s '%s' redeclared in this contract '%s'", itype, name, prev.GetContractInfo().Name)
	}
	prev.Objects[name] = &ObjInfo{Type: itype, Value: fblock}
	return nil
}

func fnFuncResult(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	fn := buf.peek().GetFuncInfo()
	(*fn).Results = append((*fn).Results, lexeme.Value.(reflect.Type))
	return nil
}

func fnReturn(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdReturn, lexeme, lexeme.Line, 0))
	return nil
}

func fnCmdError(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdError, lexeme, lexeme.Line, lexeme.Value))
	return nil
}

func fnFparam(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	block := buf.peek()
	if block.Type == ObjectType_Func && (state == stateFParam || state == stateFParamTYPE) {
		fblock := block.GetFuncInfo()
		if fblock.Names == nil {
			fblock.Params = append(fblock.Params, reflect.TypeOf(nil))
		} else {
			for key := range *fblock.Names {
				if key[0] == '_' {
					name := key[1:]
					params := append((*fblock.Names)[name].Params, reflect.TypeOf(nil))
					offset := append((*fblock.Names)[name].Offset, len(block.Vars))
					(*fblock.Names)[name] = FuncName{Params: params, Offset: offset}
					break
				}
			}
		}
	}
	if block.Objects == nil {
		block.Objects = make(map[string]*ObjInfo)
	}
	if !regexp.MustCompile(varRegexp).MatchString(lexeme.Value.(string)) {
		var val = lexeme.Value.(string)
		if len(val) > 20 {
			val = val[:20] + "..."
		}
		return fmt.Errorf("identifier expected, got '%s'", val)
	}
	if _, ok := block.Objects[lexeme.Value.(string)]; ok {
		if state == stateFParamTYPE {
			return fmt.Errorf("duplicate argument '%s'", lexeme.Value.(string))
		} else if state == stateVarType {
			return fmt.Errorf("'%s' redeclared in this code block", lexeme.Value.(string))
		}
	}
	block.Objects[lexeme.Value.(string)] = &ObjInfo{Type: ObjectType_Var, Value: &ObjInfo_Variable{Name: lexeme.Value.(string), Index: len(block.Vars)}}
	block.Vars = append(block.Vars, reflect.TypeOf(nil))
	return nil
}

func fnFtype(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	block := buf.peek()
	if block.Type == ObjectType_Func && state == stateFParam {
		fblock := block.GetFuncInfo()
		if fblock.Names == nil {
			for pkey, param := range fblock.Params {
				if param == reflect.TypeOf(nil) {
					fblock.Params[pkey] = lexeme.Value.(reflect.Type)
				}
			}
		} else {
			for key := range *fblock.Names {
				if key[0] == '_' {
					for pkey, param := range (*fblock.Names)[key[1:]].Params {
						if param == reflect.TypeOf(nil) {
							(*fblock.Names)[key[1:]].Params[pkey] = lexeme.Value.(reflect.Type)
						}
					}
					break
				}
			}
		}
	}
	for vkey, ivar := range block.Vars {
		if ivar == reflect.TypeOf(nil) {
			block.Vars[vkey] = lexeme.Value.(reflect.Type)
		}
	}
	return nil
}

func fnFtail(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	var used bool
	block := buf.peek()

	fblock := block.GetFuncInfo()
	if fblock.Names == nil {
		for pkey, param := range fblock.Params {
			if param == reflect.TypeOf(nil) {
				if used {
					return fmt.Errorf(`... parameter must be one`)
				}
				fblock.Params[pkey] = reflect.TypeOf([]any{})
				used = true
			}
		}
		block.GetFuncInfo().Variadic = true
	} else {
		for key := range *fblock.Names {
			if key[0] == '_' {
				name := key[1:]
				for pkey, param := range (*fblock.Names)[name].Params {
					if param == reflect.TypeOf(nil) {
						if used {
							return fmt.Errorf(`... parameter must be one`)
						}
						(*fblock.Names)[name].Params[pkey] = reflect.TypeOf([]any{})
						used = true
					}
				}
				offset := append((*fblock.Names)[name].Offset, len(block.Vars))
				(*fblock.Names)[name] = FuncName{Params: (*fblock.Names)[name].Params,
					Offset: offset, Variadic: true}
				break
			}
		}
	}
	for vkey, ivar := range block.Vars {
		if ivar == reflect.TypeOf(nil) {
			block.Vars[vkey] = reflect.TypeOf([]any{})
		}
	}
	return nil
}

func fnFNameParam(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	block := buf.peek()

	fblock := block.GetFuncInfo()
	if fblock.Names == nil {
		names := make(map[string]FuncName)
		fblock.Names = &names
	}
	for key := range *fblock.Names {
		if key[0] == '_' {
			delete(*fblock.Names, key)
		}
	}
	(*fblock.Names)[`_`+lexeme.Value.(string)] = FuncName{}
	return nil
}

func fnIf(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdIf, lexeme, lexeme.Line, buf.peek()))
	return nil
}

func fnWhile(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdWhile, lexeme, lexeme.Line, buf.peek()))
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdContinue, lexeme, lexeme.Line, 0))
	return nil
}

func fnContinue(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdContinue, lexeme, lexeme.Line, 0))
	return nil
}

func fnBreak(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdBreak, lexeme, lexeme.Line, 0))
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
			lexeme.GetLogger().WithFields(log.Fields{"type": ParseError, "lex_value": lexeme.Value}).Error("modifying system variable")
			return fmt.Errorf(eSysVar, lexeme.Value.(string))
		}
		ivar = VarInfo{Obj: &ObjInfo{Type: ObjectType_ExtVar, Value: &ObjInfo_ExtendVariable{Name: lexeme.Value.(string)}}, Owner: nil}
	} else {
		objInfo, tobj := findVar(lexeme.Value.(string), buf)
		if objInfo == nil || objInfo.Type != ObjectType_Var {
			return fmt.Errorf(`unknown variable '%s'`, lexeme.Value.(string))
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
		block.Code.push(newByteCode(CmdAssignVar, lexeme, lexeme.Line, prev))
	} else {
		block.Code[len(block.Code)-1] = newByteCode(CmdAssignVar, lexeme, lexeme.Line, prev)
	}
	return nil
}

func fnAssign(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	buf.peek().Code.push(newByteCode(CmdAssign, lexeme, lexeme.Line, 0))
	return nil
}

func fnTx(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	contract := buf.peek()
	if contract.Type != ObjectType_Contract {
		return fmt.Errorf(`data can only be in contract`)
	}
	(*contract).GetContractInfo().Tx = new([]*FieldInfo)
	return nil
}

func fnSettings(buf *CodeBlocks, state stateType, lexeme *Lexeme) error {
	contract := buf.peek()
	if contract.Type != ObjectType_Contract {
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

	if !regexp.MustCompile(varRegexp).MatchString(lexeme.Value.(string)) {
		var val = lexeme.Value.(string)
		if len(val) > 20 {
			val = val[:20] + "..."
		}
		return fmt.Errorf("identifier expected, got '%s'", val)
	}
	if buf.get(0).AssertVar(lexeme.Value.(string)) {
		lexeme.GetLogger().WithFields(log.Fields{"type": ParseError, "contract": info.Name, "lex_value": lexeme.Value.(string)}).Error("param variable in the data section of the contract collides with the 'builtin' variable")
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
			(*tx)[i].Original = lexeme.Type
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
		return fmt.Errorf(`there is not if before %v [%d:%d]`, lexeme.Type, lexeme.Line, lexeme.Column)
	}
	buf.get(len(*buf) - 2).Code.push(newByteCode(CmdElse, lexeme, lexeme.Line, buf.peek()))
	return nil
}
