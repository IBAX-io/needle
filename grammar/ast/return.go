package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ReturnStmt struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Expr     *Expr
}

func NewReturnStmt(b *Builder) *ReturnStmt {
	return &ReturnStmt{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_ReturnStmt,
	}
}

func (r *ReturnStmt) Parse(ctx needle.IReturnStmtContext) {
	r.Src = NewSrcPos(ctx)
	if ctx.Expr() != nil {
		expr := NewExpr(r.Builder)
		expr.Parse(ctx.Expr())
		r.Expr = expr
	}
}
