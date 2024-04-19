package ast

import "github.com/antlr4-go/antlr/v4"

type SrcPos struct {
	Line   int
	Column int
	Start  int
	End    int
	Length int
}

func NewSrcPos(ctx antlr.ParserRuleContext) SrcPos {
	return SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
}

func NewSrcPosFromSymbol(t antlr.TerminalNode) SrcPos {
	return SrcPos{
		Line:   t.GetSymbol().GetLine(),
		Column: t.GetSymbol().GetColumn(),
		Start:  t.GetSymbol().GetStart(),
		End:    t.GetSymbol().GetStop(),
		Length: t.GetSymbol().GetStop() - t.GetSymbol().GetStart() + 1,
	}
}

func (p SrcPos) GetText(input []byte) string {
	size := len(input)
	start := p.Start
	stop := p.End
	if stop >= size {
		stop = size - 1
	}
	if start >= size {
		return ""
	}
	return string(input[start : stop+1])
}
