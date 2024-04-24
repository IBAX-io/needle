package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Expr struct {
	*Builder
	Id           int32
	Src          SrcPos
	Kind         TreeType
	TreeType     TreeType
	PrimaryExpr  *PrimaryExpr
	IndexExpr    *IndexExpr
	SliceExpr    *SliceExpr
	MapExpr      *MapExpr
	ArrayExpr    *ArrayExpr
	ContractCall *ContractCall
	MulOp        *MulOp
	AddOp        *AddOp
	RelOp        *RelOp
	LogicalOp    *LogicalOp
	UnaryOp      *UnaryOp
}

func NewExpr(b *Builder) *Expr {
	return &Expr{
		Builder: b,
		Id:      b.GetReferId(),
		Kind:    TreeType_Kind_Expr,
	}
}

func (e *Expr) Parse(ctx needle.IExprContext) {
	e.Src = NewSrcPos(ctx)
	for _, tree := range ctx.GetChildren() {
		switch child := tree.(type) {
		case needle.IPrimaryExprContext:
			primaryExpr := NewPrimaryExpr(e.Builder)
			primaryExpr.Parse(child)
			e.PrimaryExpr = primaryExpr
			e.TreeType = primaryExpr.TreeType
		case needle.IIndexExprContext:
			indexExpr := NewIndexExpr(e.Builder)
			indexExpr.Parse(ctx, child)
			e.IndexExpr = indexExpr
			e.TreeType = indexExpr.TreeType
		case needle.ISliceExprContext:
			sliceExpr := NewSliceExpr(e.Builder)
			sliceExpr.Parse(ctx, child)
			e.SliceExpr = sliceExpr
			e.TreeType = sliceExpr.TreeType
		case needle.IMapExprContext:
			mapExpr := NewMapExpr(e.Builder)
			mapExpr.Parse(child)
			e.MapExpr = mapExpr
			e.TreeType = mapExpr.TreeType
		case needle.IArrayExprContext:
			arrayExpr := NewArrayExpr(e.Builder)
			arrayExpr.Parse(child)
			e.ArrayExpr = arrayExpr
			e.TreeType = arrayExpr.TreeType
		case needle.IContractCallContext:
			contractCall := NewContractCall(e.Builder)
			contractCall.Parse(child)
			e.TreeType = contractCall.TreeType
			e.ContractCall = contractCall
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

type ExprList struct {
	*Builder
	Id       int32
	Src      SrcPos
	ExprList []*Expr
}

func NewExprList(b *Builder) *ExprList {
	return &ExprList{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (e *ExprList) Parse(ctx needle.IExprListContext) {
	e.Src = NewSrcPos(ctx)
	for _, exprCtx := range ctx.AllExpr() {
		expr := NewExpr(e.Builder)
		expr.Parse(exprCtx)
		e.ExprList = append(e.ExprList, expr)
	}
}

type PrimaryExpr struct {
	*Builder
	Id          int32
	Src         SrcPos
	PrimaryExpr *PrimaryExpr
	TreeType    TreeType
	Operand     *Operand
	Arguments   *Arguments
}

func NewPrimaryExpr(b *Builder) *PrimaryExpr {
	return &PrimaryExpr{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_PrimaryExpr,
	}
}

func (p *PrimaryExpr) Parse(ctx needle.IPrimaryExprContext) {
	p.Src = NewSrcPos(ctx)
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
}
