package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type LogicalOp struct {
	*Builder

	TreeType  TreeType
	Op        string
	LeftExpr  *Expr
	RightExpr *Expr
}

func NewLogicalOp(b *Builder) *LogicalOp {
	return &LogicalOp{
		Builder:  b,
		TreeType: TreeType_LogicalOpExpr,
	}
}

func (l *LogicalOp) Parse(exprCtx needle.IExprContext, ctx needle.ILogical_opContext) {
	l.Op = ctx.GetText()
	leftExpr := NewExpr(l.Builder)
	leftExpr.Parse(exprCtx.Expr(0))
	l.LeftExpr = leftExpr
	rightExpr := NewExpr(l.Builder)
	rightExpr.Parse(exprCtx.Expr(1))
	l.RightExpr = rightExpr
}
