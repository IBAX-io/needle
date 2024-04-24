package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Arguments struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	ExprList *ExprList
}

func NewArguments(b *Builder) *Arguments {
	return &Arguments{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_Arguments,
	}
}

func (a *Arguments) Parse(ctx needle.IArgumentsContext) {
	a.Src = NewSrcPos(ctx)
	if ctx.ExprList() != nil {
		exprList := NewExprList(a.Builder)
		exprList.Parse(ctx.ExprList())
		a.ExprList = exprList
	}
}
