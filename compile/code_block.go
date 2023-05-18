package compile

import (
	"fmt"
	"reflect"
	"strings"
)

/* Byte code could be described as a tree where functions and contracts are on the top level and
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
	Objects map[string]*ObjInfo
	Type    ObjectType
	//Types that are valid to be assigned to Info:
	//*FuncInfo
	//*ContractInfo
	//*OwnerInfo
	Info           isCodeBlockInfo
	Parent         *CodeBlock
	Vars           []reflect.Type
	Code           ByteCodes
	PredeclaredVar []string
	Children       CodeBlocks
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
	for _, s := range bc.PredeclaredVar {
		if s == name {
			return true
		}
	}
	return false
}

// ByteCode stores a command and an additional parameter.
type ByteCode struct {
	Cmd    CmdT
	Lexeme *Lexeme
	Line   int
	//FuncNameCmd
	//*ObjInfo
	//uint16
	//*IndexInfo
	//[]*VarInfo
	//*VarInfo
	//*CodeBlock
	//*Map
	//[]mapItem
	//string
	//uint32
	//lexeme.Value
	//nil
	//int
	Value any
}

func newByteCode(cmd CmdT, Lexeme *Lexeme, line int, value any) *ByteCode {
	return &ByteCode{Cmd: cmd, Lexeme: Lexeme, Line: line, Value: value}
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

// ByteCodes is the slice of ByteCode items
type ByteCodes []*ByteCode

func (b ByteCodes) String() string {
	var ret string
	for _, code := range b {
		ret = ret + cmdName[code.Cmd] + " "
		ret = ret + fmt.Sprint(code.Value) + "\n"
	}
	return ret
}

func (bs *ByteCodes) push(x any) {
	*bs = append(*bs, x.(*ByteCode))
}

func (bs *ByteCodes) peek() *ByteCode {
	bsLen := len(*bs)
	if bsLen == 0 {
		return nil
	}
	return (*bs)[bsLen-1]
}

// ObjInfo is the common object type
type ObjInfo struct {
	Type ObjectType
	// Types that are valid to be assigned to Value:
	//	*CodeBlock
	//	*ExtFuncInfo
	//	*ObjInfo_Variable
	//	*ObjInfo_ExtendVariable
	Value isObjInfoValue
}

func NewCodeBlock(ext *ExtendData) *CodeBlock {
	return &CodeBlock{
		Objects: ext.MakeExtFunc(),
		// Reserved 256 indexes for system purposes
		Children:       make(CodeBlocks, 256, 1024),
		Info:           ext.Info,
		PredeclaredVar: ext.MakePreVar(),
	}
}

func (bc *CodeBlock) GetObjByName(name string) (ret *ObjInfo) {
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
		if ret.Type != ObjectType_Contract && ret.Type != ObjectType_Func {
			return nil
		}
		bc = ret.GetCodeBlock()
	}
	return
}

func (bc *CodeBlock) IsParentContract() bool {
	if bc.Parent != nil && bc.Parent.Type == ObjectType_Contract {
		return true
	}
	return false
}

func setWritable(block *CodeBlocks) {
	for i := len(*block) - 1; i >= 0; i-- {
		blockItem := (*block)[i]
		if blockItem.Type == ObjectType_Func {
			blockItem.GetFuncInfo().CanWrite = true
		}
		if blockItem.Type == ObjectType_Contract {
			blockItem.GetContractInfo().CanWrite = true
		}
	}
}
func (ret *ObjInfo) GetParamsLen() int {
	if ret.Type == ObjectType_ExtFunc {
		return len(ret.GetExtFuncInfo().Params)
	}
	return len(ret.GetCodeBlock().GetFuncInfo().Params)
}

func findVar(name string, block *CodeBlocks) (*ObjInfo, *CodeBlock) {
	if len(name) == 0 || block.peek() == nil {
		return nil, nil
	}
	resolve, ok := block.peek().resolve(name)
	if ok {
		return resolve.Objects[name], resolve
	}

	//for i := len(*block) - 1; i >= 0; i-- {
	//	if obj, ok := (*block)[i].Objects[name]; ok {
	//		return obj, (*block)[i]
	//	}
	//}
	return nil, nil
}
