package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

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
	c.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
}
