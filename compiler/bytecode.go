package compiler

import "fmt"

// Bytecode stores a command and an additional parameter.
type Bytecode struct {
	Cmd    Cmd
	Lexeme *Lexeme
	// Types that are assignable to Value:
	//
	//	* FuncTailCmd, assigned to CmdFuncTail.
	//	* IndexInfo, assigned to CmdGetIndex, CmdSetIndex.
	//	* VarInfo, assigned to CmdVar.
	//	* CodeBlock, assigned to CmdIf, CmdElse, CmdWhile
	//	* Object, assigned to CmdCall, CmdCallVariadic both for CodeBlockFunction, ObjExtFunc and CodeBlockContract
	//	* SliceItem, assigned to CmdSliceColon
	//	* Map, assigned to CmdMapInit
	//	* []*VarInfo, assigned to CmdAssignVar
	//	* MapItemList, assigned to CmdArrayInit
	//	* Lexeme.Value, assigned to CmdPush, CmdError
	//	* map[string][]any, assigned to CmdPush for function tail
	//	* string, assigned to CmdPush, CmdCallExtend, CmdExtend
	//	* uint32, assigned to CmdPush for OwnerInfo.StateId
	//	* uint16, assigned to CmdSys, op.Cmd is operatorPriority
	//	* int, assigned to CmdReturn, CmdBreak, CmdContinue, CmdWhile, CmdPush, CmdAssign, CmdUnwrapArr, CmdLabel
	Value any
}

func newBytecode(cmd Cmd, lexeme *Lexeme, value any) *Bytecode {
	return &Bytecode{Cmd: cmd, Lexeme: lexeme, Value: value}
}

func (b *Bytecode) want(w Cmd) bool {
	if b == nil {
		return false
	}
	return b.Cmd == w
}

func (b *Bytecode) FuncTailCmd() *FuncTailCmd {
	if x, ok := b.Value.(*FuncTailCmd); ok {
		return x
	}
	return nil
}

func (b *Bytecode) IndexInfo() *IndexInfo {
	if x, ok := b.Value.(*IndexInfo); ok {
		return x
	}
	return nil
}

func (b *Bytecode) VarInfo() *VarInfo {
	if x, ok := b.Value.(*VarInfo); ok {
		return x
	}
	return nil
}

func (b *Bytecode) CodeBlock() *CodeBlock {
	if x, ok := b.Value.(*CodeBlock); ok {
		return x
	}
	return nil
}

func (b *Bytecode) Map() *Map {
	if x, ok := b.Value.(*Map); ok {
		return x
	}
	return nil
}

func (b *Bytecode) SliceItem() *SliceItem {
	if x, ok := b.Value.(*SliceItem); ok {
		return x
	}
	return nil
}

func (b *Bytecode) Object() *Object {
	if x, ok := b.Value.(*Object); ok {
		return x
	}
	return nil
}

func (b *Bytecode) MapItemList() *MapItemList {
	if x, ok := b.Value.(*MapItemList); ok {
		return x
	}
	return &MapItemList{}
}

func (b *Bytecode) VarInfos() []*VarInfo {
	if x, ok := b.Value.([]*VarInfo); ok {
		return x
	}
	return []*VarInfo{}
}

// Bytecodes is the slice of Bytecode items
type Bytecodes []*Bytecode

func (b *Bytecodes) String() string {
	var ret string
	for _, code := range *b {
		ret += code.Cmd.String() + " " + fmt.Sprint(code.Value) + "\n"
	}
	return ret
}

func (b *Bytecodes) empty() bool {
	return b == nil || len(*b) == 0
}

func (b *Bytecodes) push(x *Bytecode) {
	*b = append(*b, x)
}

func (b *Bytecodes) peek() *Bytecode {
	bsLen := len(*b)
	if bsLen == 0 {
		return nil
	}
	return (*b)[bsLen-1]
}

func (b *Bytecodes) pop() *Bytecode {
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
