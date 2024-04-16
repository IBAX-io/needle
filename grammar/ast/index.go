package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type IndexStmt struct {
	*Builder

	Expr *Expr
}

func NewIndexStmt(b *Builder) *IndexStmt {
	return &IndexStmt{
		Builder: b,
	}
}

func (i *IndexStmt) Parse(ctx needle.IIndexStmtContext) {
	if ctx.Expr() != nil {
		expr := NewExpr(i.Builder)
		expr.Parse(ctx.Expr())
		i.Expr = expr
	}
}

type IndexNumber struct {
	*Builder

	Number        *NumberLiteral
	IdentifierVar *IdentifierVar
}

func NewIndexNumber(b *Builder) *IndexNumber {
	return &IndexNumber{
		Builder: b,
	}
}

func (i *IndexNumber) Parse(ctx needle.IIndexNumberContext) {
	if ctx.NumberLiteral() != nil {
		number := NewNumberLiteral(i.Builder)
		number.Parse(ctx.NumberLiteral())
		i.Number = number
	}

	if ctx.IdentifierVar() != nil {
		ident := NewIdentifierVar(i.Builder)
		ident.Parse(ctx.IdentifierVar())
		i.IdentifierVar = ident
	}
}
