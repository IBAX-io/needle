package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ArrayExpr struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	ArrayList *ArrayList
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
	if ctx.ArrayList() != nil {
		arrayList := NewArrayList(a.Builder)
		arrayList.Parse(ctx.ArrayList())
		a.ArrayList = arrayList
	}
}

type ArrayList struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	ExprList *ExprList
}

func NewArrayList(b *Builder) *ArrayList {
	return &ArrayList{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_ArrayList,
	}
}

func (a *ArrayList) Parse(ctx needle.IArrayListContext) {
	a.Src = NewSrcPos(ctx)
	if ctx.ExprList() != nil {
		exprList := NewExprList(a.Builder)
		exprList.Parse(ctx.ExprList())
		a.ExprList = exprList
	}
}
