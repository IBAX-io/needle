package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type SliceExpr struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
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

func (s *SliceExpr) Parse(ctx needle.ISliceExprContext) {
	s.Src = NewSrcPos(ctx)
	if ctx.Expr(0) != nil {
		lowExpr := NewExpr(s.Builder)
		lowExpr.Parse(ctx.Expr(0))
		s.LowExpr = lowExpr
	}

	if ctx.Expr(1) != nil {
		highExpr := NewExpr(s.Builder)
		highExpr.Parse(ctx.Expr(1))
		s.HighExpr = highExpr
	}
}
