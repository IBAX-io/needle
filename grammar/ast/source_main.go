package ast

import (
	"fmt"

	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type SourceMain struct {
	Id       int32
	Src      SrcPos
	TreeType TreeType

	ContractDefs []*ContractDef
	FuncDefs     []*FuncDef
}

func NewSourceMain() *SourceMain {
	return &SourceMain{
		TreeType:     TreeType_SourceMain,
		ContractDefs: make([]*ContractDef, 0),
		FuncDefs:     make([]*FuncDef, 0),
	}
}

func (b *Builder) ExitSourceMain(ctx *needle.SourceMainContext) {
	for _, err := range b.Errors {
		fmt.Println(err)
	}
}

func (b *Builder) EnterSourceMain(ctx *needle.SourceMainContext) {
	b.sourceMain = NewSourceMain()
	b.sourceMain.Id = b.GetReferId()
	b.sourceMain.Src = NewSrcPos(ctx)
	for _, child := range ctx.GetChildren() {
		if contractCtx, ok := child.(*needle.ContractDefContext); ok {
			def := NewContractDef(b)
			def.Parse(contractCtx, b.sourceMain)
			b.sourceMain.ContractDefs = append(b.sourceMain.ContractDefs, def)
		}
		if funcDef, ok := child.(*needle.FuncDefContext); ok {
			def := NewFuncDef(b)
			def.Parse(funcDef)
			b.sourceMain.FuncDefs = append(b.sourceMain.FuncDefs, def)
		}
	}
}
