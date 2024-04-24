package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type AddOp struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Op        string
	LeftExpr  *Expr
	RightExpr *Expr
}

func NewAddOp(b *Builder) *AddOp {
	return &AddOp{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_AddOpExpr,
	}
}

func (a *AddOp) Parse(exprCtx needle.IExprContext, ctx needle.IAdd_opContext) {
	a.Src = NewSrcPos(exprCtx)
	a.Op = ctx.GetText()
	leftExpr := NewExpr(a.Builder)
	leftExpr.Parse(exprCtx.Expr(0))
	a.LeftExpr = leftExpr
	rightExpr := NewExpr(a.Builder)
	rightExpr.Parse(exprCtx.Expr(1))
	a.RightExpr = rightExpr
}
