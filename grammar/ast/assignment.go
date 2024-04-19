package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Assignment struct {
	*Builder
	Id        int32
	Src       SrcPos
	Op        string
	LeftExpr  []*Expr
	RightExpr []*Expr
}

func NewAssignment(b *Builder) *Assignment {
	return &Assignment{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (a *Assignment) Parse(ctx needle.IAssignmentContext) {
	a.Src = NewSrcPos(ctx)
	a.Op = ctx.Assign_op().GetText()

	for _, exprCtx := range ctx.ExprList(0).AllExpr() {
		expr := NewExpr(a.Builder)
		expr.Parse(exprCtx)
		a.LeftExpr = append(a.LeftExpr, expr)
	}

	for _, exprCtx := range ctx.ExprList(1).AllExpr() {
		expr := NewExpr(a.Builder)
		expr.Parse(exprCtx)
		a.RightExpr = append(a.RightExpr, expr)
	}
}
