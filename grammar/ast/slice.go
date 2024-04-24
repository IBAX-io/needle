package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type SliceExpr struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	BaseExpr *Expr
	LowExpr  *Expr
	HighExpr *Expr
}

func NewSliceExpr(b *Builder) *SliceExpr {
	return &SliceExpr{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_SliceExpr,
	}
}

func (s *SliceExpr) Parse(exprCtx needle.IExprContext, ctx needle.ISliceExprContext) {
	s.Src = NewSrcPos(ctx)
	baseExpr := NewExpr(s.Builder)
	baseExpr.Parse(exprCtx.Expr(0))
	s.BaseExpr = baseExpr
	if ctx.GetLow() != nil {
		lowExpr := NewExpr(s.Builder)
		lowExpr.Parse(ctx.GetLow())
		s.LowExpr = lowExpr
	}

	if ctx.GetHigh() != nil {
		highExpr := NewExpr(s.Builder)
		highExpr.Parse(ctx.GetHigh())
		s.HighExpr = highExpr
	}
}
