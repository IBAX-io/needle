package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type IndexExpr struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Expr     *Expr
}

func NewIndexExpr(b *Builder) *IndexExpr {
	return &IndexExpr{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_IndexExpr,
	}
}

func (i *IndexExpr) Parse(ctx needle.IIndexExprContext) {
	i.Src = NewSrcPos(ctx)
	if ctx.Expr() != nil {
		expr := NewExpr(i.Builder)
		expr.Parse(ctx.Expr())
		i.Expr = expr
	}
}
