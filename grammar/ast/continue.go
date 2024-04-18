package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type ContinueStmt struct {
	*Builder
	Src SrcPos
}

func NewContinueStmt(b *Builder) *ContinueStmt {
	return &ContinueStmt{
		Builder: b,
	}
}

func (c *ContinueStmt) Parse(ctx needle.IContinueStmtContext) {
	c.Src = NewSrcPos(ctx)
}
