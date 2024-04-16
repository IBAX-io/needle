package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type IncDecStmt struct {
	*Builder
	Src      SrcPos
	StmtType string

	Operator   string
	Expression *Expr
}

func NewIncDecStmt(b *Builder) *IncDecStmt {
	return &IncDecStmt{
		Builder:  b,
		StmtType: "IncDecStmt",
	}
}

func (i *IncDecStmt) Parse(ctx needle.IIncDecStmtContext) {
	i.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
	if ctx.IncDec_op().DEC() != nil {
		i.Operator = ctx.IncDec_op().DEC().GetText()
	} else if ctx.IncDec_op().INC() != nil {
		i.Operator = ctx.IncDec_op().INC().GetText()
	}
	if ctx.Expr() != nil {
		expr := NewExpr(i.Builder)
		expr.Parse(ctx.Expr())
		i.Expression = expr
	}
}
