package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Arguments struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	ExprList *ExprList
	Ellipsis bool
}

func NewArguments(b *Builder) *Arguments {
	return &Arguments{
		Builder:  b,
		Id:       b.GetNextId(),
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
	if ctx.TAIL() != nil {
		a.Ellipsis = true
	}
}
