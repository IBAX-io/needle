package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Operand struct {
	*Builder
	Id            int32
	Src           SrcPos
	TreeType      TreeType
	IdentifierVar *IdentifierVar
	Literal       *Literal
	Expr          *Expr
}

func NewOperand(b *Builder) *Operand {
	return &Operand{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_Operand,
	}
}

func (o *Operand) Parse(ctx needle.IOperandContext) {
	o.Src = NewSrcPos(ctx)
	if ctx.IdentifierVar() != nil {
		ident := NewIdentifierVar(o.Builder)
		ident.Parse(ctx.IdentifierVar())
		o.IdentifierVar = ident
	}

	if ctx.Literal() != nil {
		literal := NewLiteral(o.Builder)
		literal.Parse(ctx.Literal())
		o.Literal = literal
	}

	if ctx.Expr() != nil {
		expr := NewExpr(o.Builder)
		expr.Parse(ctx.Expr())
		o.Expr = expr
	}
}
