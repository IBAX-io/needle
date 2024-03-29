package compiler

import (
	"reflect"
	"strconv"
	"strings"
)

/*
	ByteCode could be described as a tree where functions and contracts are on the top level and nesting goes further
	according to nesting of bracketed. Tree nodes are structures of 'CodeBlock' type. For instance,
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

type CodeBlockType int32

const (
	CodeBlockDefault CodeBlockType = iota + 1
	CodeBlockOwner
	CodeBlockContract
	CodeBlockFunction
	CodeBlockIf
	CodeBlockElse
	CodeBlockWhile
)

var CodeBlockTypeName = map[int32]string{
	1: "CodeBlockDefault",
	2: "CodeBlockOwner",
	3: "CodeBlockContract",
	4: "CodeBlockFunction",
	5: "CodeBlockIf",
	6: "CodeBlockElse",
	7: "CodeBlockWhile",
}

func (c CodeBlockType) String() string {
	s, ok := CodeBlockTypeName[int32(c)]
	if ok {
		return s
	}
	return strconv.Itoa(int(c))
}

// CodeBlock contains all information about compiled block {...} and its children
type CodeBlock struct {
	Objects map[string]*Object
	Type    CodeBlockType
	// Types that are assignable to Info:
	//
	//  *FunctionInfo
	//  *ContractInfo
	//  *OwnerInfo
	//  *CodeBlockIfInfo
	//  *CodeBlockElseInfo
	//  *CodeBlockWhileInfo
	Info   isCodeBlockInfo
	Parent *CodeBlock
	Vars   []Token
	Code   ByteCodes
	// PredeclaredVar is a list of variables that are declared in the block
	PredeclaredVar []string
	Children       CodeBlocks
}

func NewCodeBlock(conf *CompConfig) *CodeBlock {
	if conf == nil {
		conf = new(CompConfig)
	}
	return &CodeBlock{
		Objects: conf.MakeExtFunc(),
		// Reserved 256 indexes for system purposes
		Children:       make(CodeBlocks, 256, 1024),
		Type:           CodeBlockOwner,
		Info:           conf.Owner,
		PredeclaredVar: conf.PreVar,
	}
}

type isCodeBlockInfo interface {
	isCodeBlockInfo()
}

func (*OwnerInfo) isCodeBlockInfo()          {}
func (*ContractInfo) isCodeBlockInfo()       {}
func (*FunctionInfo) isCodeBlockInfo()       {}
func (*CodeBlockIfInfo) isCodeBlockInfo()    {}
func (*CodeBlockElseInfo) isCodeBlockInfo()  {}
func (*CodeBlockWhileInfo) isCodeBlockInfo() {}

func (bc *CodeBlock) GetInfo() isCodeBlockInfo {
	if bc != nil {
		return bc.Info
	}
	return nil
}

func (bc *CodeBlock) GetFunctionInfo() *FunctionInfo {
	if x, ok := bc.GetInfo().(*FunctionInfo); ok {
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

// AssertVar checks if the variable is declared in the block
func (bc *CodeBlock) AssertVar(name string) bool {
	for _, s := range bc.PredeclaredVar {
		if s == name {
			return true
		}
	}
	return false
}

// SetInfo sets the type of the block. It could be a function, contract, if, else or while.
func (bc *CodeBlock) SetInfo(info isCodeBlockInfo) {
	var t CodeBlockType
	switch info.(type) {
	case *FunctionInfo:
		t = CodeBlockFunction
	case *ContractInfo:
		t = CodeBlockContract
	case *OwnerInfo:
		t = CodeBlockOwner
	case *CodeBlockIfInfo:
		t = CodeBlockIf
	case *CodeBlockElseInfo:
		t = CodeBlockElse
	case *CodeBlockWhileInfo:
		t = CodeBlockWhile
	}
	bc.Type = t
	bc.Info = info
}

// CheckLoop checks if the current block or any of its parent blocks is of
// type 'ObjWhile', which represents a loop in the code. If a loop is found,
// it returns true, otherwise false.
func (bc *CodeBlock) CheckLoop() bool {
	if bc.Type == CodeBlockWhile {
		return true
	}
	if bc.Parent != nil {
		return bc.Parent.CheckLoop()
	}
	return false
}

// GetName returns the name of the block.
func (bc *CodeBlock) GetName() string {
	var name string
	switch bc.Type {
	case CodeBlockContract:
		name = bc.GetContractInfo().Name
	case CodeBlockFunction:
		name = bc.GetFunctionInfo().Name
	}
	return name
}

// GetType returns the type of the block.
func (bc *CodeBlock) GetType() CodeBlockType {
	return bc.Type
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
		if !ret.IsCodeBlockContract() && !ret.IsCodeBlockFunction() {
			return nil
		}
		bc = ret.GetCodeBlock()
	}
	return
}

func (bc *CodeBlock) IsParentContract() bool {
	if bc.Parent != nil && bc.Parent.Type == CodeBlockContract {
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
				CanWrite: item.CanWrite,
			}
			for i := 0; i < fobj.NumIn(); i++ {
				if isauto, ok := item.AutoPars[fobj.In(i).String()]; ok {
					data.Auto[i] = isauto
				}
				data.Params[i] = fobj.In(i)
			}
			for i := 0; i < fobj.NumOut(); i++ {
				data.Results[i] = fobj.Out(i)
			}
			bc.Objects[item.Name] = NewObject(data)
		}
	}
}

func (bc *CodeBlock) CodeToString() string {
	var ret string
	ret = bc.Code.String()
	for _, child := range bc.Children {
		ret += child.CodeToString()
	}
	return ret
}

func setWritable(cbs *CodeBlocks) {
	for i := len(*cbs) - 1; i >= 0; i-- {
		cb := (*cbs)[i]
		if cb.Type == CodeBlockFunction {
			cb.GetFunctionInfo().CanWrite = true
		}
		if cb.Type == CodeBlockContract {
			cb.GetContractInfo().CanWrite = true
		}
	}
}

func findVar(name string, cbs *CodeBlocks) (*Object, *CodeBlock) {
	if len(name) == 0 || cbs.peek() == nil {
		return nil, nil
	}
	resolve, ok := cbs.peek().resolve(name)
	if ok {
		return resolve.Objects[name], resolve
	}
	return nil, nil
}
