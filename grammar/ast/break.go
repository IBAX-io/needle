package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type BreakStmt struct {
	*Builder
	Src SrcPos
}

func NewBreakStmt(b *Builder) *BreakStmt {
	return &BreakStmt{
		Builder: b,
	}
}

func (b *BreakStmt) Parse(ctx needle.IBreakStmtContext) {
	b.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
}
