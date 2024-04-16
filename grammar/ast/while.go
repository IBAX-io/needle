package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type WhileStmt struct {
	*Builder

	Expr  *Expr
	Block *Block
}

func NewWhileStmt(b *Builder) *WhileStmt {
	return &WhileStmt{
		Builder: b,
	}
}

func (w *WhileStmt) Parse(ctx needle.IWhileStmtContext) {
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
