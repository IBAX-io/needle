package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type MapExpr struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	PairList *PairList
}

func NewMapExpr(b *Builder) *MapExpr {
	return &MapExpr{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_MapExpr,
	}
}

func (o *MapExpr) Parse(ctx needle.IMapExprContext) {
	o.Src = NewSrcPos(ctx)
	if ctx.PairList() != nil {
		pairList := NewPairList(o.Builder)
		pairList.Parse(ctx.PairList())
		o.PairList = pairList
	}
}

type PairList struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Pairs    []*Pair
}

func NewPairList(b *Builder) *PairList {
	return &PairList{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_PairList,
	}
}

func (l *PairList) Parse(ctx needle.IPairListContext) {
	l.Src = NewSrcPos(ctx)
	for _, pair := range ctx.AllPair() {
		p := NewPair(l.Builder)
		p.Parse(pair)
		l.Pairs = append(l.Pairs, p)
	}
}

type Pair struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Key      string
	Expr     *Expr
}

func NewPair(b *Builder) *Pair {
	return &Pair{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_Pair,
	}
}

func (p *Pair) Parse(ctx needle.IPairContext) {
	p.Src = NewSrcPos(ctx)
	if ctx.IdentifierVar() != nil {
		identifierVar := NewIdentifierVar(p.Builder)
		identifierVar.Parse(ctx.IdentifierVar())
		p.Key = identifierVar.Name
		p.TreeType = identifierVar.TreeType
	}
	if ctx.StringLiteral() != nil {
		stringLiteral := NewStringLiteral(p.Builder)
		stringLiteral.Parse(ctx.StringLiteral())
		p.Key = stringLiteral.Value
		p.TreeType = stringLiteral.TreeType
	}

	expr := NewExpr(p.Builder)
	expr.Parse(ctx.Expr())
	p.Expr = expr
}
