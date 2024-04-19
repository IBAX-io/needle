package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type LogicalOp struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Op        string
	LeftExpr  *Expr
	RightExpr *Expr
}

func NewLogicalOp(b *Builder) *LogicalOp {
	return &LogicalOp{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_LogicalOpExpr,
	}
}

func (l *LogicalOp) Parse(exprCtx needle.IExprContext, ctx needle.ILogical_opContext) {
	l.Src = NewSrcPos(ctx)
	l.Op = ctx.GetText()
	leftExpr := NewExpr(l.Builder)
	leftExpr.Parse(exprCtx.Expr(0))
	l.LeftExpr = leftExpr
	rightExpr := NewExpr(l.Builder)
	rightExpr.Parse(exprCtx.Expr(1))
	l.RightExpr = rightExpr
}
