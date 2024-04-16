package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type SliceStmt struct {
	*Builder

	LeftExpr  *IndexNumber
	RightExpr *IndexNumber
}

func NewSliceStmt(b *Builder) *SliceStmt {
	return &SliceStmt{
		Builder: b,
	}
}

func (s *SliceStmt) Parse(ctx needle.ISliceStmtContext) {
	if ctx.IndexNumber(0) != nil {
		leftExpr := NewIndexNumber(s.Builder)
		leftExpr.Parse(ctx.IndexNumber(0))
		s.LeftExpr = leftExpr
	}

	if ctx.IndexNumber(1) != nil {
		rightExpr := NewIndexNumber(s.Builder)
		rightExpr.Parse(ctx.IndexNumber(1))
		s.RightExpr = rightExpr
	}
}
