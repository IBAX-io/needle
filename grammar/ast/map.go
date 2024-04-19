package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type MapExpr struct {
	*Builder

	TreeType TreeType
	PairList *PairList
}

func NewMapExpr(b *Builder) *MapExpr {
	return &MapExpr{
		Builder:  b,
		TreeType: TreeType_MapExpr,
	}
}

func (o *MapExpr) Parse(ctx needle.IMapExprContext) {
	if ctx.PairList() != nil {
		pairList := NewPairList(o.Builder)
		pairList.Parse(ctx.PairList())
		o.PairList = pairList
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

	IdentifierVar *IdentifierVar
	Literal       *Literal

	IndexExpr *IndexExpr
	SliceExpr *SliceExpr

	MapExpr   *MapExpr
	ArrayExpr *ArrayExpr
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
		if ctx.IndexExpr() != nil {
			indexExpr := NewIndexExpr(p.Builder)
			indexExpr.Parse(ctx.IndexExpr())
			p.IndexExpr = indexExpr
		}
		if ctx.SliceExpr() != nil {
			sliceExpr := NewSliceExpr(p.Builder)
			sliceExpr.Parse(ctx.SliceExpr())
			p.SliceExpr = sliceExpr
		}
	}
	if ctx.Literal() != nil {
		literal := NewLiteral(p.Builder)
		literal.Parse(ctx.Literal())
		p.Literal = literal
	}
	if ctx.MapExpr() != nil {
		mapExpr := NewMapExpr(p.Builder)
		mapExpr.Parse(ctx.MapExpr())
		p.MapExpr = mapExpr
	}
	if ctx.ArrayExpr() != nil {
		arrayExpr := NewArrayExpr(p.Builder)
		arrayExpr.Parse(ctx.ArrayExpr())
		p.ArrayExpr = arrayExpr
	}
}
