package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type SourceMain struct {
	Src      SrcPos
	StmtType string

	ContractDefs []*ContractDef
	FuncDefs     []*FuncDef
}

func NewSourceMain() *SourceMain {
	return &SourceMain{StmtType: "SourceMain"}
}

func (b *Builder) EnterSourceMain(ctx *needle.SourceMainContext) {
	main := NewSourceMain()
	for _, child := range ctx.GetChildren() {
		if contractCtx, ok := child.(*needle.ContractDefContext); ok {
			def := NewContractDef(b)
			def.Parse(contractCtx, main)
		}
		if funcDef, ok := child.(*needle.FuncDefContext); ok {
			def := NewFuncDef(b)
			def.Parse(funcDef)
		}
	}
}
