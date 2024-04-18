package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type SliceStmt struct {
	*Builder

	LowExpr  *IndexNumber
	HighExpr *IndexNumber
}

func NewSliceStmt(b *Builder) *SliceStmt {
	return &SliceStmt{
		Builder: b,
	}
}

func (s *SliceStmt) Parse(ctx needle.ISliceStmtContext) {
	if ctx.IndexNumber(0) != nil {
		lowExpr := NewIndexNumber(s.Builder)
		lowExpr.Parse(ctx.IndexNumber(0))
		s.LowExpr = lowExpr
	}

	if ctx.IndexNumber(1) != nil {
		highExpr := NewIndexNumber(s.Builder)
		highExpr.Parse(ctx.IndexNumber(1))
		s.HighExpr = highExpr
	}
}
