package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type FuncDef struct {
	*Builder
	Src      SrcPos
	StmtType string

	Name          string
	FuncSignature *FuncSignature
	Body          *Block
}

func NewFuncDef(b *Builder) *FuncDef {
	return &FuncDef{
		Builder:  b,
		StmtType: "FuncDef",
	}
}

func (d *FuncDef) Parse(ctx needle.IFuncDefContext) {
	d.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}

	if descriptor := ctx.FuncDescriptor(); descriptor != nil {
		d.Name = descriptor.Identifier().GetText()
	}

	if deft := ctx.DefaultFuncDef(); deft != nil {
		if deft.ACTION() != nil {
			d.Name = deft.ACTION().GetText()
		}
		if deft.CONDITIONS() != nil {
			d.Name = deft.CONDITIONS().GetText()
		}
	}

	sign := NewFuncSignature(d.Builder)
	sign.Parse(ctx.FuncSignature())
	d.FuncSignature = sign

	block := NewBlock(d.Builder)
	block.Parse(ctx.Block())
	d.Body = block
}

type FuncSignature struct {
	*Builder

	ParameterList    *ParameterList
	FuncTail         []*FuncTail
	ReturnParameters *ReturnParameters
}

func NewFuncSignature(b *Builder) *FuncSignature {
	return &FuncSignature{
		Builder:  b,
		FuncTail: make([]*FuncTail, 0),
	}
}

func (d *FuncSignature) Parse(ctx needle.IFuncSignatureContext) {
	if ctx.ParameterList() != nil {
		d.ParameterList = NewParameterList(d.Builder)
		d.ParameterList.Parse(ctx.ParameterList())
	}
	for _, tail := range ctx.AllFuncTail() {
		t := NewFuncTail(d.Builder)
		t.Parse(tail)
		d.FuncTail = append(d.FuncTail, t)
	}
	if ctx.ReturnParameters() != nil {
		returnParameters := NewReturnParameters(d.Builder)
		returnParameters.Parse(ctx.ReturnParameters())
		d.ReturnParameters = returnParameters
	}
}

func (d *FuncSignature) HasTail() bool {
	return len(d.FuncTail) > 0
}

type FuncTail struct {
	*Builder

	Name          string
	ParameterList *ParameterList
}

func NewFuncTail(b *Builder) *FuncTail {
	return &FuncTail{
		Builder: b,
	}
}

func (d *FuncTail) Parse(ctx needle.IFuncTailContext) {
	if ctx.ParameterList() != nil {
		d.ParameterList = NewParameterList(d.Builder)
		d.ParameterList.Parse(ctx.ParameterList())
	}
	d.Name = ctx.Identifier().GetText()
}
