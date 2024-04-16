package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type InitMapArrStmt struct {
	*Builder
}

func NewInitMapArrStmt(b *Builder) *InitMapArrStmt {
	return &InitMapArrStmt{
		Builder: b,
	}
}

func (i *InitMapArrStmt) Parse(ctx needle.IInitMapArrStmtContext) {
	if ctx.MapStmt() != nil {
		mapStmt := NewMapStmt(i.Builder)
		mapStmt.Parse(ctx.MapStmt())
	}
	if ctx.ArrayStmt() != nil {
	}
}

type PairList struct {
	*Builder
}

func NewPairList(b *Builder) *PairList {
	return &PairList{
		Builder: b,
	}
}

func (l *PairList) Parse(ctx needle.IPairListContext) {
	for _, pair := range ctx.AllPair() {
		p := NewPair(l.Builder)
		p.Parse(pair)
	}
}

type Pair struct {
	*Builder

	KeyName   string
	KeyType   string
	PairValue *PairValue
}

func NewPair(b *Builder) *Pair {
	return &Pair{
		Builder: b,
	}
}

func (p *Pair) Parse(ctx needle.IPairContext) {
	if ctx.IdentifierVar() != nil {
		identifierVar := NewIdentifierVar(p.Builder)
		identifierVar.Parse(ctx.IdentifierVar())
		p.KeyName = identifierVar.Identifier
		p.KeyType = identifierVar.IdentifierType
	}
	if ctx.StringLiteral() != nil {
		p.KeyName = ctx.StringLiteral().GetText()
		p.KeyType = "StringLiteral"
	}

	pairValue := NewPairValue(p.Builder)
	pairValue.Parse(ctx.PairValue())
	p.PairValue = pairValue
}

type PairValue struct {
	*Builder
}

func NewPairValue(b *Builder) *PairValue {
	return &PairValue{
		Builder: b,
	}
}

func (p *PairValue) Parse(ctx needle.IPairValueContext) {
	if ctx.IdentifierVar() != nil {
		identifierVar := NewIdentifierVar(p.Builder)
		identifierVar.Parse(ctx.IdentifierVar())
		if ctx.IndexStmt() != nil {
			indexStmt := NewIndexStmt(p.Builder)
			indexStmt.Parse(ctx.IndexStmt())
		}
		if ctx.SliceStmt() != nil {
			sliceStmt := NewSliceStmt(p.Builder)
			sliceStmt.Parse(ctx.SliceStmt())
		}
	}
	if ctx.Literal() != nil {
		literal := NewLiteral(p.Builder)
		literal.Parse(ctx.Literal())
	}
	if ctx.ArrayStmt() != nil {
		arrayStmt := NewArrayStmt(p.Builder)
		arrayStmt.Parse(ctx.ArrayStmt())
	}
	if ctx.MapStmt() != nil {
		mapStmt := NewMapStmt(p.Builder)
		mapStmt.Parse(ctx.MapStmt())
	}
}
