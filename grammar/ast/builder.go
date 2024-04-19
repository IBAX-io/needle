package ast

import (
	"sync/atomic"

	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Builder struct {
	*needle.BaseNeedleParserListener
	parser *needle.NeedleParser
	input  []byte

	referId        int32
	sourceMain     *SourceMain
	commentsParsed bool
	Comments       []*Comment
}

func NewBuilder(parser *needle.NeedleParser, input []byte) *Builder {
	return &Builder{
		parser:  parser,
		input:   input,
		referId: 1,
	}
}

func (b *Builder) GetReferId() int32 {
	return atomic.AddInt32(&b.referId, 1) - 1
}

func (b *Builder) GetSourceMain() *SourceMain {
	return b.sourceMain
}
