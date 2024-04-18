package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Assignment struct {
	*Builder

	Op        string
	LeftExpr  []*Expr
	RightExpr []*Expr
}

func NewAssignment(b *Builder) *Assignment {
	return &Assignment{
		Builder: b,
	}
}

func (a *Assignment) Parse(ctx needle.IAssignmentContext) {
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

type AssignMapArrStmt struct {
	*Builder

	IdentifierVar  string
	InitMapArrStmt *InitMapArrStmt
}

func NewAssignMapArrStmt(b *Builder) *AssignMapArrStmt {
	return &AssignMapArrStmt{
		Builder: b,
	}
}

func (a *AssignMapArrStmt) Parse(ctx needle.IAssignMapArrStmtContext) {
	a.IdentifierVar = ctx.IdentifierVar().GetText()
	initMapArrStmt := NewInitMapArrStmt(a.Builder)
	initMapArrStmt.Parse(ctx.InitMapArrStmt())
	a.InitMapArrStmt = initMapArrStmt
}
