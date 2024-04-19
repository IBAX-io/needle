package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ArrayExpr struct {
	*Builder

	TreeType  TreeType
	ArrayList *ArrayList
}

func NewArrayExpr(b *Builder) *ArrayExpr {
	return &ArrayExpr{
		Builder:  b,
		TreeType: TreeType_ArrayExpr,
	}
}

func (a *ArrayExpr) Parse(ctx needle.IArrayExprContext) {
	if ctx.ArrayList() != nil {
		arrayList := NewArrayList(a.Builder)
		arrayList.Parse(ctx.ArrayList())
		a.ArrayList = arrayList
	}
}

type ArrayList struct {
	*Builder

	ExprList *ExprList
}

func NewArrayList(b *Builder) *ArrayList {
	return &ArrayList{
		Builder: b,
	}
}

func (a *ArrayList) Parse(ctx needle.IArrayListContext) {
	if ctx.ExprList() != nil {
		exprList := NewExprList(a.Builder)
		exprList.Parse(ctx.ExprList())
		a.ExprList = exprList
	}
}
