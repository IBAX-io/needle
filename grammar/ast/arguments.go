package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Arguments struct {
	*Builder

	ExprList *ExprList
}

func NewArguments(b *Builder) *Arguments {
	return &Arguments{
		Builder: b,
	}
}

func (a *Arguments) Parse(ctx needle.IArgumentsContext) {
	if ctx.ExprList() != nil {
		exprList := NewExprList(a.Builder)
		exprList.Parse(ctx.ExprList())
		a.ExprList = exprList
	}
}
