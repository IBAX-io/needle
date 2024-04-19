package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type IdentifierList struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Name     []string
}

func NewIdentifierList(b *Builder) *IdentifierList {
	return &IdentifierList{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_Identifier,
		Name:     make([]string, 0),
	}
}

func (i *IdentifierList) Parse(ctx needle.IIdentifierListContext) {
	i.Src = NewSrcPos(ctx)

	for _, terminal := range ctx.AllIdentifier() {
		i.Name = append(i.Name, terminal.GetText())
	}
}

type IdentifierFull struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Name     string
}

func NewIdentifierFull(b *Builder) *IdentifierFull {
	return &IdentifierFull{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (i *IdentifierFull) Parse(ctx needle.IIdentifierFullContext) {
	i.Src = NewSrcPos(ctx)
	if ident := ctx.Identifier(); ident != nil {
		i.Name = ident.GetText()
		i.TreeType = TreeType_Identifier
	}
	if dollar := ctx.DollarIdentifier(); dollar != nil {
		i.Name = dollar.GetText()
		i.TreeType = TreeType_DollarIdentifier
	}
	if at := ctx.AtIdentifier(); at != nil {
		i.Name = at.GetText()
		i.TreeType = TreeType_AtIdentifier
	}
}

type IdentifierVar struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Name     string
}

func NewIdentifierVar(b *Builder) *IdentifierVar {
	return &IdentifierVar{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (i *IdentifierVar) Parse(ctx needle.IIdentifierVarContext) {
	i.Src = NewSrcPos(ctx)
	if ident := ctx.Identifier(); ident != nil {
		i.Name = ident.GetText()
		i.TreeType = TreeType_Identifier
	}
	if dollar := ctx.DollarIdentifier(); dollar != nil {
		i.Name = dollar.GetText()
		i.TreeType = TreeType_DollarIdentifier
	}
}
