package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type MulOp struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Op        string
	LeftExpr  *Expr
	RightExpr *Expr
}

func NewMulOp(b *Builder) *MulOp {
	return &MulOp{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_MulOpExpr,
	}
}

func (m *MulOp) Parse(exprCtx needle.IExprContext, ctx needle.IMul_opContext) {
	m.Src = NewSrcPos(exprCtx)
	m.Op = ctx.GetText()
	leftExpr := NewExpr(m.Builder)
	leftExpr.Parse(exprCtx.Expr(0))
	m.LeftExpr = leftExpr
	rightExpr := NewExpr(m.Builder)
	rightExpr.Parse(exprCtx.Expr(1))
	m.RightExpr = rightExpr
}
