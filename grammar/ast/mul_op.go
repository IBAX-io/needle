package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type MulOp struct {
	*Builder

	TreeType  TreeType
	Op        string
	LeftExpr  *Expr
	RightExpr *Expr
}

func NewMulOp(b *Builder) *MulOp {
	return &MulOp{
		Builder:  b,
		TreeType: TreeType_Expr_MulOp,
	}
}

func (m *MulOp) Parse(exprCtx needle.IExprContext, ctx needle.IMul_opContext) {
	m.Op = ctx.GetText()
	leftExpr := NewExpr(m.Builder)
	leftExpr.Parse(exprCtx.Expr(0))
	m.LeftExpr = leftExpr
	rightExpr := NewExpr(m.Builder)
	rightExpr.Parse(exprCtx.Expr(1))
	m.RightExpr = rightExpr
}
