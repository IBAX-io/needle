package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Parameter struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType

	NameList []*IdentifierList
	TypeName []string
}

func NewParameter(b *Builder) *Parameter {
	return &Parameter{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_Parameter,
		NameList: make([]*IdentifierList, 0),
		TypeName: make([]string, 0),
	}
}

func (p *Parameter) Parse(ctx needle.IParameterContext) {
	p.Src = NewSrcPos(ctx)
	for i, context := range ctx.AllIdentifierList() {
		identifierList := NewIdentifierList(p.Builder)
		identifierList.Parse(context)
		p.NameList = append(p.NameList, identifierList)
		p.TypeName = append(p.TypeName, ctx.AllTypeName()[i].GetText())
	}
}

type ParameterList struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType

	Parameter  *Parameter
	Variadic   bool
	Identifier string
}

func NewParameterList(b *Builder) *ParameterList {
	return &ParameterList{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_ParameterList,
	}
}

func (p *ParameterList) Parse(ctx needle.IParameterListContext) {
	p.Src = NewSrcPos(ctx)
	if ctx.Parameter() == nil {
		return
	}
	parameter := NewParameter(p.Builder)
	parameter.Parse(ctx.Parameter())
	p.Parameter = parameter
	if ctx.Identifier() != nil {
		p.Variadic = true
		p.Identifier = ctx.Identifier().GetText()
	}
}

type ReturnParameters struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	TypeName []string
}

func NewReturnParameters(b *Builder) *ReturnParameters {
	return &ReturnParameters{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_ReturnParameters,
		TypeName: make([]string, 0),
	}
}

func (r *ReturnParameters) Parse(ctx needle.IReturnParametersContext) {
	r.Src = NewSrcPos(ctx)
	for _, typeName := range ctx.AllTypeName() {
		r.TypeName = append(r.TypeName, typeName.GetText())
	}
}
