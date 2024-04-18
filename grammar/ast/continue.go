package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type ContinueStmt struct {
	*Builder
	Src      SrcPos
	TreeType TreeType
}

func NewContinueStmt(b *Builder) *ContinueStmt {
	return &ContinueStmt{
		Builder:  b,
		TreeType: TreeType_ContinueStmt,
	}
}

func (c *ContinueStmt) Parse(ctx needle.IContinueStmtContext) {
	c.Src = NewSrcPos(ctx)
}
