package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type BreakStmt struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
}

func NewBreakStmt(b *Builder) *BreakStmt {
	return &BreakStmt{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_BreakStmt,
	}
}

func (b *BreakStmt) Parse(ctx needle.IBreakStmtContext) {
	b.Src = NewSrcPos(ctx)
}
