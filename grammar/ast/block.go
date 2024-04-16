package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Block struct {
	*Builder

	Statements []*Statement
}

func NewBlock(b *Builder) *Block {
	return &Block{
		Builder: b,
	}
}

func (b *Block) Parse(ctx needle.IBlockContext) {
	if ctx.StatementList() != nil {
		stmtList := NewStatementList(b.Builder)
		stmtList.Parse(ctx.StatementList())
		b.Statements = stmtList.Statements
	}
}
