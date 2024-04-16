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
	v.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
	v.Parameter = NewParameter(v.Builder)
	v.Parameter.Parse(ctx.Parameter())
}
