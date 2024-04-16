package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type IdentifierList struct {
	*Builder
	Src      SrcPos
	StmtType string

	Name []string
}

func NewIdentifierList(b *Builder) *IdentifierList {
	return &IdentifierList{
		Builder:  b,
		StmtType: "Identifier",
		Name:     make([]string, 0),
	}
}

func (i *IdentifierList) Parse(ctx *needle.IdentifierListContext) {
	i.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}

	for _, terminal := range ctx.AllIdentifier() {
		i.Name = append(i.Name, terminal.GetText())
	}
}

type IdentifierFull struct {
	*Builder

	Identifier     string
	IdentifierType string
}

func NewIdentifierFull(b *Builder) *IdentifierFull {
	return &IdentifierFull{
		Builder: b,
	}
}

func (i *IdentifierFull) Parse(ctx needle.IIdentifierFullContext) {
	if ident := ctx.Identifier(); ident != nil {
		i.Identifier = ident.GetText()
		i.IdentifierType = "Identifier"
	}
	if dollar := ctx.DollarIdentifier(); dollar != nil {
		i.Identifier = dollar.GetText()
		i.IdentifierType = "DollarIdentifier"
	}
	if at := ctx.AtIdentifier(); at != nil {
		i.Identifier = at.GetText()
		i.IdentifierType = "AtIdentifier"
	}
}

type IdentifierVar struct {
	*Builder
	Identifier     string
	IdentifierType string
}

func NewIdentifierVar(b *Builder) *IdentifierVar {
	return &IdentifierVar{
		Builder: b,
	}
}

func (i *IdentifierVar) Parse(ctx needle.IIdentifierVarContext) {
	if ident := ctx.Identifier(); ident != nil {
		i.Identifier = ident.GetText()
		i.IdentifierType = "Identifier"
	}
	if dollar := ctx.DollarIdentifier(); dollar != nil {
		i.Identifier = dollar.GetText()
		i.IdentifierType = "DollarIdentifier"
	}
}
