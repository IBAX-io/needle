package compile

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	Byte code could be described as a tree where functions and contracts are on the top level and

nesting goes further according to nesting of bracketed. Tree nodes are structures of
'CodeBlock' type. For instance,

	 func a {
		 if b {
			 while d {

			 }
		 }
		 if c {
		 }
	 }
	 will be compiled into CodeBlock(a) which will have two child blocks CodeBlock (b) and CodeBlock (c) that
	 are responsible for executing bytecode inside if. CodeBlock (b) will have a child CodeBlock (d) with
	 a cycle.
*/

// CodeBlock contains all information about compiled block {...} and its children
type CodeBlock struct {
	Objects map[string]*Object
	Type    ObjectType
	//Types that are valid to be assigned to Info:
	//*FuncInfo
	//*ContractInfo
	//*OwnerInfo
	//or nil if the block is braced by {}, for example: if,else,while
	Info   isCodeBlockInfo
	Parent *CodeBlock
	Vars   []reflect.Type
	Code   ByteCodes
	// PredeclaredVar is a list of variables that are declared in the block
	PreVar   []string
	Children CodeBlocks
}

func NewCodeBlock(conf *CompConfig) *CodeBlock {
	return &CodeBlock{
		Objects: conf.MakeExtFunc(),
		// Reserved 256 indexes for system purposes
		Children: make(CodeBlocks, 256, 1024),
		Type:     conf.Owner.ObjectType(),
		Info:     conf.Owner,
		PreVar:   conf.PreVar,
	}
}

type isCodeBlockInfo interface {
	isCodeBlockInfo()
}

func (*OwnerInfo) isCodeBlockInfo()    {}
func (*ContractInfo) isCodeBlockInfo() {}
func (*FuncInfo) isCodeBlockInfo()     {}

func (bc *CodeBlock) GetInfo() isCodeBlockInfo {
	if bc != nil {
		return bc.Info
	}
	return nil
}

func (bc *CodeBlock) GetFuncInfo() *FuncInfo {
	if x, ok := bc.GetInfo().(*FuncInfo); ok {
		return x
	}
	return nil
}

func (bc *CodeBlock) GetContractInfo() *ContractInfo {
	if x, ok := bc.GetInfo().(*ContractInfo); ok {
		return x
	}
	return nil
}

func (bc *CodeBlock) GetOwnerInfo() *OwnerInfo {
	if x, ok := bc.GetInfo().(*OwnerInfo); ok {
		return x
	}
	return nil
}

func (bc *CodeBlock) resolve(name string) (*CodeBlock, bool) {
	_, ok := bc.Objects[name]
	if ok {
		return bc, true
	}
	if bc.Parent != nil {
		return bc.Parent.resolve(name)
	}
	return nil, false
}

func (bc *CodeBlock) AssertVar(name string) bool {
	for _, s := range bc.PreVar {
		if s == name {
			return true
		}
	}
	return false
}

// CodeBlocks is a slice of blocks
type CodeBlocks []*CodeBlock

func (bs *CodeBlocks) ParentOwner() *OwnerInfo {
	bsLen := len(*bs)
	if bsLen == 0 {
		return nil
	}
	return (*bs)[0].GetOwnerInfo()
}

func (bs *CodeBlocks) push(x any) {
	*bs = append(*bs, x.(*CodeBlock))
}

func (bs *CodeBlocks) peek() *CodeBlock {
	bsLen := len(*bs)
	if bsLen == 0 {
		return nil
	}
	return (*bs)[bsLen-1]
}

func (bs *CodeBlocks) get(idx int) *CodeBlock {
	if idx >= 0 && len(*bs) > 0 && len(*bs) > idx {
		return (*bs)[idx]
	}
	return nil
}

func (bc *CodeBlock) GetObjByName(name string) (ret *Object) {
	if bc == nil {
		return nil
	}
	var ok bool
	names := strings.Split(name, `.`)
	for i, name := range names {
		ret, ok = bc.Objects[name]
		if !ok {
			return nil
		}
		if i == len(names)-1 {
			return
		}
		if ret.Type != ObjContract && ret.Type != ObjFunc {
			return nil
		}
		bc = ret.GetCodeBlock()
	}
	return
}

func (bc *CodeBlock) IsParentContract() bool {
	if bc.Parent != nil && bc.Parent.Type == ObjContract {
		return true
	}
	return false
}

func (bc *CodeBlock) SetExtendFunc(ext []ExtendFunc) {
	for _, item := range ext {
		fobj := reflect.ValueOf(item.Func).Type()
		switch fobj.Kind() {
		case reflect.Func:
			data := &ExtFuncInfo{
				Name:     item.Name,
				Params:   make([]reflect.Type, fobj.NumIn()),
				Results:  make([]reflect.Type, fobj.NumOut()),
				Auto:     make([]string, fobj.NumIn()),
				Variadic: fobj.IsVariadic(),
				Func:     item.Func,
				CanWrite: item.CanWrite}
			for i := 0; i < fobj.NumIn(); i++ {
				if isauto, ok := item.AutoPars[fobj.In(i).String()]; ok {
					data.Auto[i] = isauto
				}
				data.Params[i] = fobj.In(i)
			}
			for i := 0; i < fobj.NumOut(); i++ {
				data.Results[i] = fobj.Out(i)
			}
			bc.Objects[item.Name] = NewObject(ObjExtFunc, data)
		}
	}
}

// ByteCode stores a command and an additional parameter.
type ByteCode struct {
	Cmd    CmdT
	Lexeme *Lexeme
	//*FuncTailCmd, assigned to CmdFuncTail.
	//*IndexInfo, assigned to CmdGetIndex, CmdSetIndex
	//*VarInfo, assigned to CmdVar
	//*CodeBlock, assigned to CmdIf, CmdElse, CmdWhile
	//*Object, assigned to CmdCall, CmdCallVariadic both for ObjFunc, ObjExtFunc and ObjContract
	//*SliceItem, assigned to CmdSliceColon
	//*Map, assigned to CmdMapInit
	//[]*VarInfo, assigned to CmdAssignVar
	//[]*MapItem, assigned to CmdArrayInit
	//Lexeme.Value, assigned to CmdPush, CmdError
	//map[string][]any, assigned to CmdPush for function tail
	//string, assigned to CmdPush, CmdCallExtend, CmdExtend
	//uint32, assigned to CmdPush for OwnerInfo.StateID
	//uint16, assigned to CmdSys, op.Cmd is operatorPriority
	//int, assigned to CmdReturn, CmdBreak, CmdContinue, CmdWhile, CmdPush, CmdAssign, CmdUnwrapArr, CmdLabel
	Value any
}

func newByteCode(cmd CmdT, Lexeme *Lexeme, value any) *ByteCode {
	return &ByteCode{Cmd: cmd, Lexeme: Lexeme, Value: value}
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
		ret = ret + code.Cmd.String() + " "
		ret = ret + fmt.Sprint(code.Value) + "\n"
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

func setWritable(block *CodeBlocks) {
	for i := len(*block) - 1; i >= 0; i-- {
		blockItem := (*block)[i]
		if blockItem.Type == ObjFunc {
			blockItem.GetFuncInfo().CanWrite = true
		}
		if blockItem.Type == ObjContract {
			blockItem.GetContractInfo().CanWrite = true
		}
	}
}

func findVar(name string, block *CodeBlocks) (*Object, *CodeBlock) {
	if len(name) == 0 || block.peek() == nil {
		return nil, nil
	}
	resolve, ok := block.peek().resolve(name)
	if ok {
		return resolve.Objects[name], resolve
	}
	return nil, nil
}
