package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type IncDecStmt struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType

	OpPos      SrcPos
	Operator   string
	Expression *Expr
}

func NewIncDecStmt(b *Builder) *IncDecStmt {
	return &IncDecStmt{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_IncDecStmt,
	}
}

func (i *IncDecStmt) Parse(ctx needle.IIncDecStmtContext) {
	i.Src = NewSrcPos(ctx)
	if ctx.IncDec_op().DEC() != nil {
		i.OpPos = NewSrcPosFromSymbol(ctx.IncDec_op().DEC())
		i.Operator = ctx.IncDec_op().DEC().GetText()
	} else if ctx.IncDec_op().INC() != nil {
		i.OpPos = NewSrcPosFromSymbol(ctx.IncDec_op().INC())
		i.Operator = ctx.IncDec_op().INC().GetText()
	}
	if ctx.Expr() != nil {
		expr := NewExpr(i.Builder)
		expr.Parse(ctx.Expr())
		i.Expression = expr
	}
}
