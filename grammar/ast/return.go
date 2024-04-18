package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ReturnStmt struct {
	*Builder

	TreeType TreeType
	Expr     *Expr
}

func NewReturnStmt(b *Builder) *ReturnStmt {
	return &ReturnStmt{
		Builder:  b,
		TreeType: TreeType_ReturnStmt,
	}
}

func (r *ReturnStmt) Parse(ctx needle.IReturnStmtContext) {
	if ctx.Expr() != nil {
		expr := NewExpr(r.Builder)
		expr.Parse(ctx.Expr())
		r.Expr = expr
	}
}
