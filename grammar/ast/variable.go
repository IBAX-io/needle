package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type VarDef struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Parameter *Parameter
}

func NewVarDef(b *Builder) *VarDef {
	return &VarDef{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_VarDef,
	}
}

func (v *VarDef) Parse(ctx needle.IVarDefContext) {
	v.Src = NewSrcPos(ctx)
	v.Parameter = NewParameter(v.Builder)
	v.Parameter.Parse(ctx.Parameter())
}
