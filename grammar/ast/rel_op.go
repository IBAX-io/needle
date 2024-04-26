package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type RelOp struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Op        string
	LeftExpr  *Expr
	RightExpr *Expr
}

func NewRelOp(b *Builder) *RelOp {
	return &RelOp{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_RelOpExpr,
	}
}

func (r *RelOp) Parse(exprCtx needle.IExprContext, ctx needle.IRel_opContext) {
	r.Src = NewSrcPos(ctx)
	r.Op = ctx.GetText()
	leftExpr := NewExpr(r.Builder)
	leftExpr.Parse(exprCtx.Expr(0))
	r.LeftExpr = leftExpr
	rightExpr := NewExpr(r.Builder)
	rightExpr.Parse(exprCtx.Expr(1))
	r.RightExpr = rightExpr
}
