package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type ErrorStmt struct {
	*Builder
	Id       int32
	Src      SrcPos
	Kind     TreeType
	TreeType TreeType
	Type     string
	Expr     *Expr
}

func NewErrorStmt(b *Builder) *ErrorStmt {
	return &ErrorStmt{
		Builder: b,
		Id:      b.GetNextId(),
		Kind:    TreeType_Kind_ErrorStmt,
	}
}

func (e *ErrorStmt) Parse(ctx needle.IErrorStmtContext) {
	e.Src = NewSrcPos(ctx)

	if ctx.ERRWARNING() != nil {
		e.TreeType = TreeType_ErrWarningStmt
		e.Type = ctx.ERRWARNING().GetText()
	}

	if ctx.ERRINFO() != nil {
		e.TreeType = TreeType_ErrInfoStmt
		e.Type = ctx.ERRINFO().GetText()
	}

	if ctx.ERROR() != nil {
		e.TreeType = TreeType_ErrorStmt
		e.Type = ctx.ERROR().GetText()
	}

	if ctx.Expr() != nil {
		expr := NewExpr(e.Builder)
		expr.Parse(ctx.Expr())
		e.Expr = expr
	}
}
