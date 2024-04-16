package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Parameter struct {
	*Builder
	Src      SrcPos
	StmtType string

	NameList []IdentifierList
	TypeName []string
}

func NewParameter(b *Builder) *Parameter {
	return &Parameter{
		Builder:  b,
		StmtType: "Parameter",
		NameList: make([]IdentifierList, 0),
		TypeName: make([]string, 0),
	}
}

func (p *Parameter) Parse(ctx needle.IParameterContext) {
	p.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
	for i, context := range ctx.AllIdentifierList() {
		identifierList := NewIdentifierList(p.Builder)
		identifierList.Parse(context.(*needle.IdentifierListContext))
		p.NameList = append(p.NameList, *identifierList)
		p.TypeName = append(p.TypeName, ctx.AllTypeName()[i].GetText())
	}
}

type ParameterList struct {
	*Builder
	Src      SrcPos
	StmtType string

	Parameter *Parameter
}

func NewParameterList(b *Builder) *ParameterList {
	return &ParameterList{
		Builder:  b,
		StmtType: "ParameterList",
	}
}

func (p *ParameterList) Parse(ctx needle.IParameterListContext) {
	p.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
	if ctx.Parameter() == nil {
		return
	}
	parameter := NewParameter(p.Builder)
	parameter.Parse(ctx.Parameter())
	p.Parameter = parameter
}

type ReturnParameters struct {
	*Builder
	TypeName []string
}

func NewReturnParameters(b *Builder) *ReturnParameters {
	return &ReturnParameters{
		Builder:  b,
		TypeName: make([]string, 0),
	}
}

func (r *ReturnParameters) Parse(ctx needle.IReturnParametersContext) {
	for _, typeName := range ctx.AllTypeName() {
		r.TypeName = append(r.TypeName, typeName.GetText())
	}
}