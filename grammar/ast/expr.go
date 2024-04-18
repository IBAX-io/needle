package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Expr struct {
	*Builder

	Kind     TreeType
	TreeType TreeType

	PrimaryExpr *PrimaryExpr
	MulOp       *MulOp
	AddOp       *AddOp
	RelOp       *RelOp
	LogicalOp   *LogicalOp
	UnaryOp     *UnaryOp
}

func NewExpr(b *Builder) *Expr {
	return &Expr{
		Builder: b,
		Kind:    TreeType_Kind_Expr,
	}
}

func (e *Expr) Parse(ctx needle.IExprContext) {
	for _, tree := range ctx.GetChildren() {
		switch child := tree.(type) {
		case needle.IPrimaryExprContext:
			primaryExpr := NewPrimaryExpr(e.Builder)
			primaryExpr.Parse(child)
			e.PrimaryExpr = primaryExpr
			e.TreeType = primaryExpr.TreeType
		case needle.IMul_opContext:
			mulOp := NewMulOp(e.Builder)
			mulOp.Parse(ctx, child)
			e.MulOp = mulOp
			e.TreeType = mulOp.TreeType
		case needle.IAdd_opContext:
			addOp := NewAddOp(e.Builder)
			addOp.Parse(ctx, child)
			e.AddOp = addOp
			e.TreeType = addOp.TreeType
		case needle.IRel_opContext:
			relOp := NewRelOp(e.Builder)
			relOp.Parse(ctx, child)
			e.RelOp = relOp
			e.TreeType = relOp.TreeType
		case needle.ILogical_opContext:
			logicalOp := NewLogicalOp(e.Builder)
			logicalOp.Parse(ctx, child)
			e.LogicalOp = logicalOp
			e.TreeType = logicalOp.TreeType
		case needle.IUnary_opContext:
			unaryOp := NewUnaryOp(e.Builder)
			unaryOp.Parse(ctx, child)
			e.UnaryOp = unaryOp
			e.TreeType = unaryOp.TreeType
		}
	}
}

type PrimaryExpr struct {
	*Builder

	TreeType    TreeType
	Operand     *Operand
	PrimaryExpr *PrimaryExpr

	Arguments *Arguments
	IndexStmt *IndexStmt
	SliceStmt *SliceStmt
}

func NewPrimaryExpr(b *Builder) *PrimaryExpr {
	return &PrimaryExpr{
		Builder:  b,
		TreeType: TreeType_PrimaryExpr,
	}
}

func (p *PrimaryExpr) Parse(ctx needle.IPrimaryExprContext) {
	if ctx.Operand() != nil {
		operand := NewOperand(p.Builder)
		operand.Parse(ctx.Operand())
		p.Operand = operand
	}

	if ctx.PrimaryExpr() != nil {
		primaryExpr := NewPrimaryExpr(p.Builder)
		primaryExpr.Parse(ctx.PrimaryExpr())
		p.PrimaryExpr = primaryExpr
	}

	if ctx.Arguments() != nil {
		arguments := NewArguments(p.Builder)
		arguments.Parse(ctx.Arguments())
		p.Arguments = arguments
	}

	if ctx.IndexStmt() != nil {
		indexStmt := NewIndexStmt(p.Builder)
		indexStmt.Parse(ctx.IndexStmt())
		p.IndexStmt = indexStmt
	}

	if ctx.SliceStmt() != nil {
		sliceStmt := NewSliceStmt(p.Builder)
		sliceStmt.Parse(ctx.SliceStmt())
		p.SliceStmt = sliceStmt
	}
}
