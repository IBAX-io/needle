package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type FuncDef struct {
	*Builder
	Id       int32
	Src      SrcPos
	StmtType string

	Name          string
	FuncSignature *FuncSignature
	Body          *Block
}

func NewFuncDef(b *Builder) *FuncDef {
	return &FuncDef{
		Builder:  b,
		Id:       b.GetReferId(),
		StmtType: "FuncDef",
	}
}

func (d *FuncDef) Parse(ctx needle.IFuncDefContext) {
	d.Src = NewSrcPos(ctx)

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
	Id               int32
	Src              SrcPos
	ParameterList    *ParameterList
	FuncTail         []*FuncTail
	ReturnParameters *ReturnParameters
}

func NewFuncSignature(b *Builder) *FuncSignature {
	return &FuncSignature{
		Builder:  b,
		Id:       b.GetReferId(),
		FuncTail: make([]*FuncTail, 0),
	}
}

func (d *FuncSignature) Parse(ctx needle.IFuncSignatureContext) {
	d.Src = NewSrcPos(ctx)
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
	Id            int32
	Src           SrcPos
	Name          string
	ParameterList *ParameterList
}

func NewFuncTail(b *Builder) *FuncTail {
	return &FuncTail{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (d *FuncTail) Parse(ctx needle.IFuncTailContext) {
	d.Src = NewSrcPos(ctx)
	if ctx.ParameterList() != nil {
		d.ParameterList = NewParameterList(d.Builder)
		d.ParameterList.Parse(ctx.ParameterList())
	}
	d.Name = ctx.Identifier().GetText()
}
