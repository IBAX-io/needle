package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type WhileStmt struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Expr     *Expr
	Block    *Block
}

func NewWhileStmt(b *Builder) *WhileStmt {
	return &WhileStmt{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_WhileStmt,
	}
}

func (w *WhileStmt) Parse(ctx needle.IWhileStmtContext) {
	w.Src = NewSrcPos(ctx)
	if ctx.Expr() != nil {
		expr := NewExpr(w.Builder)
		expr.Parse(ctx.Expr())
		w.Expr = expr
	}
	if ctx.Block() != nil {
		block := NewBlock(w.Builder)
		block.Parse(ctx.Block())
		w.Block = block
	}
}
