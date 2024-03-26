package compiler

import "fmt"

// ByteCode stores a command and an additional parameter.
type ByteCode struct {
	Cmd    Cmd
	Lexeme *Lexeme
	// Types that are assignable to Value:
	//
	//  *FuncTailCmd, assigned to CmdFuncTail.
	//  *IndexInfo, assigned to CmdGetIndex, CmdSetIndex.
	//	*VarInfo, assigned to CmdVar.
	//	*CodeBlock, assigned to CmdIf, CmdElse, CmdWhile
	//	*Object, assigned to CmdCall, CmdCallVariadic both for ObjFunction, ObjExtFunc and ObjContract
	//	*SliceItem, assigned to CmdSliceColon
	//	*Map, assigned to CmdMapInit
	//	*[]*VarInfo, assigned to CmdAssignVar
	//	*[]*MapItem, assigned to CmdArrayInit
	//	*Lexeme.Value, assigned to CmdPush, CmdError
	//	*map[string][]any, assigned to CmdPush for function tail
	//	*string, assigned to CmdPush, CmdCallExtend, CmdExtend
	//	*uint32, assigned to CmdPush for OwnerInfo.StateId
	//	*uint16, assigned to CmdSys, op.Cmd is operatorPriority
	//	*int, assigned to CmdReturn, CmdBreak, CmdContinue, CmdWhile, CmdPush, CmdAssign, CmdUnwrapArr, CmdLabel
	Value any
}

func newByteCode(cmd Cmd, lexeme *Lexeme, value any) *ByteCode {
	return &ByteCode{Cmd: cmd, Lexeme: lexeme, Value: value}
}

func (b *ByteCode) want(w Cmd) bool {
	if b == nil {
		return false
	}
	return b.Cmd == w
}

func (b *ByteCode) FuncTailCmd() *FuncTailCmd {
	if x, ok := b.Value.(*FuncTailCmd); ok {
		return x
	}
	return nil
}

func (b *ByteCode) IndexInfo() *IndexInfo {
	if x, ok := b.Value.(*IndexInfo); ok {
		return x
	}
	return nil
}

func (b *ByteCode) VarInfo() *VarInfo {
	if x, ok := b.Value.(*VarInfo); ok {
		return x
	}
	return nil
}

func (b *ByteCode) CodeBlock() *CodeBlock {
	if x, ok := b.Value.(*CodeBlock); ok {
		return x
	}
	return nil
}

func (b *ByteCode) Map() *Map {
	if x, ok := b.Value.(*Map); ok {
		return x
	}
	return nil
}

func (b *ByteCode) SliceItem() *SliceItem {
	if x, ok := b.Value.(*SliceItem); ok {
		return x
	}
	return nil
}

func (b *ByteCode) Object() *Object {
	if x, ok := b.Value.(*Object); ok {
		return x
	}
	return nil
}

func (b *ByteCode) MapItems() []*MapItem {
	if x, ok := b.Value.([]*MapItem); ok {
		return x
	}
	return []*MapItem{}
}

func (b *ByteCode) VarInfos() []*VarInfo {
	if x, ok := b.Value.([]*VarInfo); ok {
		return x
	}
	return []*VarInfo{}
}

// ByteCodes is the slice of ByteCode items
type ByteCodes []*ByteCode

func (b *ByteCodes) String() string {
	var ret string
	for _, code := range *b {
		ret = ret + code.Cmd.String() + " " + fmt.Sprint(code.Value) + "\n"
	}
	return ret
}

func (b *ByteCodes) empty() bool {
	return b == nil || len(*b) == 0
}

func (b *ByteCodes) push(x *ByteCode) {
	*b = append(*b, x)
}

func (b *ByteCodes) peek() *ByteCode {
	bsLen := len(*b)
	if bsLen == 0 {
		return nil
	}
	return (*b)[bsLen-1]
}

func (b *ByteCodes) pop() *ByteCode {
	bsLen := len(*b)
	if bsLen == 0 {
		return nil
	}
	ret := (*b)[bsLen-1]
	*b = (*b)[:bsLen-1]
	return ret
}

type (
	// FuncTailCmd contains the function tail and the count of parameters.
	FuncTailCmd struct {
		Count    int
		FuncTail FuncTail
	}

	// VarInfo contains the variable or extended variable information.
	VarInfo struct {
		Obj   *Object
		Owner *CodeBlock // is nil if the variable is global
	}

	// IndexInfo contains the information for SetIndex.
	IndexInfo struct {
		VarOffset int
		Owner     *CodeBlock
		Extend    string
	}
)
