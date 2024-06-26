package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ContractDef struct {
	*Builder
	Id            int32
	Src           SrcPos
	TreeType      TreeType
	Name          string
	NamePos       SrcPos
	ContractParts []*ContractPart
}

func NewContractDef(b *Builder) *ContractDef {
	return &ContractDef{
		Builder:       b,
		Id:            b.GetNextId(),
		TreeType:      TreeType_ContractDef,
		ContractParts: make([]*ContractPart, 0),
	}
}

func (d *ContractDef) Parse(ctx needle.IContractDefContext, main *SourceMain) {
	main.Src = NewSrcPos(ctx)
	d.Name = ctx.Identifier().GetText()
	d.Src = NewSrcPos(ctx)
	d.Name = ctx.Identifier().GetText()
	d.NamePos = NewSrcPosFromSymbol(ctx.Identifier())
	for _, part := range ctx.AllContractPart() {
		contractPart := NewContractPart(d.Builder)
		contractPart.Parse(part)
		d.ContractParts = append(d.ContractParts, contractPart)
	}
}

type ContractPart struct {
	*Builder
	Id          int32
	Src         SrcPos
	TreeType    TreeType
	DataDef     *DataDef
	SettingsDef *SettingsDef
	FuncDef     *FuncDef
}

func NewContractPart(b *Builder) *ContractPart {
	return &ContractPart{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_ContractPart,
	}
}

func (d *ContractPart) Parse(ctx needle.IContractPartContext) {
	d.Src = NewSrcPos(ctx)
	for _, tree := range ctx.GetChildren() {
		switch childCtx := tree.(type) {
		case *needle.DataDefContext:
			def := NewDataDef(d.Builder)
			def.Parse(childCtx)
			d.DataDef = def
		case *needle.SettingsDefContext:
			def := NewSettingsDef(d.Builder)
			def.Parse(childCtx)
			d.SettingsDef = def
		case *needle.FuncDefContext:
			def := NewFuncDef(d.Builder)
			def.Parse(childCtx)
			d.FuncDef = def
		default:
			panic("ContractPart unexpected type")
		}
	}
}

type ContractCall struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Name      string
	Arguments *Arguments
}

func NewContractCall(b *Builder) *ContractCall {
	return &ContractCall{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_ContractCall,
	}
}

func (c *ContractCall) Parse(ctx needle.IContractCallContext) {
	c.Src = NewSrcPos(ctx)
	c.Name = ctx.AtIdentifier().GetText()
	if ctx.Arguments() != nil {
		args := NewArguments(c.Builder)
		args.Parse(ctx.Arguments())
		c.Arguments = args
	}
}
