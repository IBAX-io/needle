package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ArrayExpr struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	ExprList *ExprList
	Comma    bool
}

func NewArrayExpr(b *Builder) *ArrayExpr {
	return &ArrayExpr{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_ArrayExpr,
	}
}

func (a *ArrayExpr) Parse(ctx needle.IArrayExprContext) {
	a.Src = NewSrcPos(ctx)
	if ctx.ExprList() != nil {
		exprList := NewExprList(a.Builder)
		exprList.Parse(ctx.ExprList())
		a.ExprList = exprList
	}
	if ctx.COMMA() != nil {
		a.Comma = true
	}
}
