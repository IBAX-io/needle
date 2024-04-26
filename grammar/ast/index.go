package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type IndexExpr struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	BaseExpr *Expr
	Index    *Expr
}

func NewIndexExpr(b *Builder) *IndexExpr {
	return &IndexExpr{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_IndexExpr,
	}
}

func (i *IndexExpr) Parse(exprCtx needle.IExprContext, ctx needle.IIndexExprContext) {
	i.Src = NewSrcPos(exprCtx)
	baseExpr := NewExpr(i.Builder)
	baseExpr.Parse(exprCtx.Expr(0))
	i.BaseExpr = baseExpr
	expr := NewExpr(i.Builder)
	expr.Parse(ctx.GetIndex())
	i.Index = expr
}
