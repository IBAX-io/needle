package ast

import (
	"encoding/hex"
	"strings"

	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Literal struct {
	*Builder
	Src      SrcPos
	StmtType string

	Kind            string
	Value, HexValue string
}

func NewLiteral(b *Builder) *Literal {
	return &Literal{
		Builder:  b,
		StmtType: "Literal",
	}
}

func (d *Literal) Parse(ctx needle.ILiteralContext) {
	d.Src = NewSrcPos(ctx)

	if ctx.BooleanLiteral() != nil {
		d.Kind = "boolean"
		d.Value = strings.TrimSpace(
			strings.ReplaceAll(ctx.BooleanLiteral().GetText(), "\"", ""),
		)
		d.HexValue = hex.EncodeToString([]byte(d.Value))
	}

	if ctx.StringLiteral() != nil {
		d.Kind = "string"
		d.Value = strings.TrimSpace(
			strings.ReplaceAll(ctx.StringLiteral().GetText(), "\"", ""),
		)
		d.HexValue = hex.EncodeToString([]byte(d.Value))

	}

	if numLit := ctx.NumberLiteral(); numLit != nil {
		lit := NewNumberLiteral(d.Builder)
		lit.Parse(numLit)
		d.Kind = lit.Kind
		d.Value = lit.Value
		d.HexValue = lit.HexValue
	}
}

type NumberLiteral struct {
	*Builder

	Kind            string
	Value, HexValue string
}

func NewNumberLiteral(b *Builder) *NumberLiteral {
	return &NumberLiteral{
		Builder: b,
	}
}

func (d *NumberLiteral) Parse(ctx needle.INumberLiteralContext) {
	if ctx.FloatLiteral() != nil {
		d.Kind = "float"
		d.Value = ctx.FloatLiteral().GetText()
	}
	if ctx.DecimalLiteral() != nil {
		d.Kind = "decimal"
		d.Value = ctx.DecimalLiteral().GetText()
	}
	if ctx.BinaryLiteral() != nil {
		d.Kind = "binary"
		d.Value = ctx.BinaryLiteral().GetText()
	}
	if ctx.OctalLiteral() != nil {
		d.Kind = "octal"
		d.Value = ctx.OctalLiteral().GetText()
	}
	if ctx.HexLiteral() != nil {
		d.Kind = "hex"
		d.Value = ctx.HexLiteral().GetText()
	}
	d.HexValue = hex.EncodeToString([]byte(d.Value))
}
