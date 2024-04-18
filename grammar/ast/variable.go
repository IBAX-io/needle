package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type VarDef struct {
	*Builder
	Src SrcPos

	Parameter *Parameter
}

func NewVarDef(b *Builder) *VarDef {
	return &VarDef{
		Builder: b,
	}
}

func (v *VarDef) Parse(ctx needle.IVarDefContext) {
	v.Src = NewSrcPos(ctx)
	v.Parameter = NewParameter(v.Builder)
	v.Parameter.Parse(ctx.Parameter())
}
