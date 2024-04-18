package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Operand struct {
	*Builder

	IdentifierFull *IdentifierFull
	Literal        *Literal
	Expr           *Expr
	Nil            string
}

func NewOperand(b *Builder) *Operand {
	return &Operand{
		Builder: b,
	}
}

func (o *Operand) Parse(ctx needle.IOperandContext) {
	if ctx.IdentifierFull() != nil {
		ident := NewIdentifierFull(o.Builder)
		ident.Parse(ctx.IdentifierFull())
		o.IdentifierFull = ident
	}

	if ctx.Literal() != nil {
		literal := NewLiteral(o.Builder)
		literal.Parse(ctx.Literal())
		o.Literal = literal
	}
	if ctx.NIL() != nil {
		o.Nil = ctx.NIL().GetText()
	}
	if ctx.Expr() != nil {
		expr := NewExpr(o.Builder)
		expr.Parse(ctx.Expr())
		o.Expr = expr
	}
}
