package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ContractDef struct {
	*Builder
	Src      SrcPos
	StmtType string

	Name    string
	NamePos SrcPos
	Nodes   []NodeType
}

func NewContractDef(builder *Builder) *ContractDef {
	return &ContractDef{
		Builder:  builder,
		StmtType: "ContractDef",
		Nodes:    make([]NodeType, 0),
	}
}

func (d *ContractDef) Parse(ctx *needle.ContractDefContext, main *SourceMain) {
	main.Src = NewSrcPos(ctx)
	d.Name = ctx.Identifier().GetText()
	d.Src = NewSrcPos(ctx)
	d.Name = ctx.Identifier().GetText()
	d.NamePos = NewSrcPosFromSymbol(ctx.Identifier())
	for _, part := range ctx.AllContractPart() {
		contractPart := NewContractPart(d.Builder)
		contractPart.Parse(part)
		d.Nodes = append(d.Nodes, contractPart.NodeType)
	}

	main.ContractDefs = append(main.ContractDefs, d)
}

type ContractPart struct {
	*Builder
	Src      SrcPos
	StmtType string

	NodeType NodeType
}

func NewContractPart(b *Builder) *ContractPart {
	return &ContractPart{
		Builder:  b,
		StmtType: "ContractPart",
	}
}

func (d *ContractPart) Parse(ctx needle.IContractPartContext) {
	for _, tree := range ctx.GetChildren() {
		switch childCtx := tree.(type) {
		case *needle.DataDefContext:
			def := NewDataDef(d.Builder)
			def.Parse(childCtx)
			d.NodeType = def
		case *needle.SettingsDefContext:
			def := NewSettingsDef(d.Builder)
			def.Parse(childCtx)
			d.NodeType = def
		case *needle.FuncDefContext:
			def := NewFuncDef(d.Builder)
			def.Parse(childCtx)
			d.NodeType = def
		default:
			panic("ContractPart unexpected type")
		}
	}
}
