package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type IndexExpr struct {
	*Builder

	TreeType TreeType
	Expr     *Expr
}

func NewIndexExpr(b *Builder) *IndexExpr {
	return &IndexExpr{
		Builder:  b,
		TreeType: TreeType_IndexExpr,
	}
}

func (i *IndexExpr) Parse(ctx needle.IIndexExprContext) {
	if ctx.Expr() != nil {
		expr := NewExpr(i.Builder)
		expr.Parse(ctx.Expr())
		i.Expr = expr
	}
}
