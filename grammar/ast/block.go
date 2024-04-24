package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Block struct {
	*Builder
	Id            int32
	Src           SrcPos
	TreeType      TreeType
	StatementList *StatementList
}

func NewBlock(b *Builder) *Block {
	return &Block{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_Block,
	}
}

func (b *Block) Parse(ctx needle.IBlockContext) {
	b.Src = NewSrcPos(ctx)
	if ctx.StatementList() != nil {
		stmtList := NewStatementList(b.Builder)
		stmtList.Parse(ctx.StatementList())
		b.StatementList = stmtList
	}
}
