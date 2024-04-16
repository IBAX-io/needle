package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type MapStmt struct {
	*Builder
}

func NewMapStmt(b *Builder) *MapStmt {
	return &MapStmt{
		Builder: b,
	}
}

func (o *MapStmt) Parse(ctx needle.IMapStmtContext) {
	if ctx.PairList() != nil {
		pairList := NewPairList(o.Builder)
		pairList.Parse(ctx.PairList())
	}
}
