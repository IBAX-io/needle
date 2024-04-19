package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type UnaryOp struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Op        string
	RightExpr *Expr
}

func NewUnaryOp(b *Builder) *UnaryOp {
	return &UnaryOp{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_UnaryOpExpr,
	}
}

func (u *UnaryOp) Parse(exprCtx needle.IExprContext, ctx needle.IUnary_opContext) {
	u.Src = NewSrcPos(ctx)
	u.Op = ctx.GetText()
	rightExpr := NewExpr(u.Builder)
	rightExpr.Parse(exprCtx.Expr(0))
	u.RightExpr = rightExpr
}
