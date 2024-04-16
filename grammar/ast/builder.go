package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Builder struct {
	*needle.BaseNeedleParserListener
}

func NewBuilder() *Builder {
	return &Builder{}
}

type SrcPos struct {
	Line   int
	Column int
	Start  int
	End    int
	Length int
}
