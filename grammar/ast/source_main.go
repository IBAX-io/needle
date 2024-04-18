package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type SourceMain struct {
	Src      SrcPos
	StmtType string

	ContractDefs []*ContractDef
	FuncDefs     []*FuncDef
}

func NewSourceMain() *SourceMain {
	return &SourceMain{
		StmtType:     "SourceMain",
		ContractDefs: make([]*ContractDef, 0),
		FuncDefs:     make([]*FuncDef, 0),
	}
}

func (b *Builder) EnterSourceMain(ctx *needle.SourceMainContext) {
	b.sourceMain = NewSourceMain()
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
