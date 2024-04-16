package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ErrorStmt struct {
	*Builder
	Src SrcPos

	Type string
	Expr *Expr
}

func NewErrorStmt(b *Builder) *ErrorStmt {
	return &ErrorStmt{
		Builder: b,
	}
}

func (e *ErrorStmt) Parse(ctx needle.IErrorStmtContext) {
	e.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}

	if ctx.ERRWARNING() != nil {
		e.Type = ctx.ERRWARNING().GetText()
	}

	if ctx.ERRINFO() != nil {
		e.Type = ctx.ERRINFO().GetText()
	}

	if ctx.ERROR() != nil {
		e.Type = ctx.ERROR().GetText()
	}

	if ctx.Expr() != nil {
		expr := NewExpr(e.Builder)
		expr.Parse(ctx.Expr())
		e.Expr = expr
	}
}
