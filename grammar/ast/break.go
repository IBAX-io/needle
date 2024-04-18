package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type BreakStmt struct {
	*Builder
	Src      SrcPos
	TreeType TreeType
}

func NewBreakStmt(b *Builder) *BreakStmt {
	return &BreakStmt{
		Builder:  b,
		TreeType: TreeType_BreakStmt,
	}
}

func (b *BreakStmt) Parse(ctx needle.IBreakStmtContext) {
	b.Src = NewSrcPos(ctx)
}
