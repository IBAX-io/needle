package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type InitMapArrStmt struct {
	*Builder

	MapStmt   *MapStmt
	ArrayStmt *ArrayStmt
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
		i.MapStmt = mapStmt
	}
	if ctx.ArrayStmt() != nil {
		arrayStmt := NewArrayStmt(i.Builder)
		arrayStmt.Parse(ctx.ArrayStmt())
		i.ArrayStmt = arrayStmt
	}
}

type PairList struct {
	*Builder

	Pairs []*Pair
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
		l.Pairs = append(l.Pairs, p)
	}
}

type Pair struct {
	*Builder

	KeyName   string
	TreeType  TreeType
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
		p.KeyName = identifierVar.Name
		p.TreeType = identifierVar.TreeType
	}
	if ctx.StringLiteral() != nil {
		stringLiteral := NewStringLiteral(p.Builder)
		stringLiteral.Parse(ctx.StringLiteral())
		p.KeyName = stringLiteral.Value
		p.TreeType = stringLiteral.TreeType
	}

	pairValue := NewPairValue(p.Builder)
	pairValue.Parse(ctx.PairValue())
	p.PairValue = pairValue
}

type PairValue struct {
	*Builder

	IdentifierVar  *IdentifierVar
	Literal        *Literal
	InitMapArrStmt *InitMapArrStmt

	IndexStmt *IndexStmt
	SliceStmt *SliceStmt
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
		p.IdentifierVar = identifierVar
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
	if ctx.Literal() != nil {
		literal := NewLiteral(p.Builder)
		literal.Parse(ctx.Literal())
		p.Literal = literal
	}
	if ctx.InitMapArrStmt() != nil {
		initMapArrStmt := NewInitMapArrStmt(p.Builder)
		initMapArrStmt.Parse(ctx.InitMapArrStmt())
		p.InitMapArrStmt = initMapArrStmt
	}
}
