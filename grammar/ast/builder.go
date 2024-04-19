package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Builder struct {
	*needle.BaseNeedleParserListener
	parser *needle.NeedleParser
	input  []byte

	sourceMain     *SourceMain
	commentsParsed bool
	Comments       []*Comment
}

func NewBuilder(parser *needle.NeedleParser, input []byte) *Builder {
	return &Builder{parser: parser, input: input}
}

func (b *Builder) GetSourceMain() *SourceMain {
	return b.sourceMain
}
