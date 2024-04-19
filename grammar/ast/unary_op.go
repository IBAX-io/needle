package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type UnaryOp struct {
	*Builder

	TreeType  TreeType
	Op        string
	RightExpr *Expr
}

func NewUnaryOp(b *Builder) *UnaryOp {
	return &UnaryOp{
		Builder:  b,
		TreeType: TreeType_UnaryOpExpr,
	}
}

func (u *UnaryOp) Parse(exprCtx needle.IExprContext, ctx needle.IUnary_opContext) {
	u.Op = ctx.GetText()
	rightExpr := NewExpr(u.Builder)
	rightExpr.Parse(exprCtx.Expr(0))
	u.RightExpr = rightExpr
}
