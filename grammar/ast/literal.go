package ast

import (
	"encoding/hex"
	"strings"

	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type Literal struct {
	*Builder
	Id              int32
	Src             SrcPos
	Kind            TreeType
	TreeType        TreeType
	Value, HexValue string
}

func NewLiteral(b *Builder) *Literal {
	return &Literal{
		Builder: b,
		Id:      b.GetReferId(),
		Kind:    TreeType_Kind_Literal,
	}
}

func (d *Literal) Parse(ctx needle.ILiteralContext) {
	d.Src = NewSrcPos(ctx)

	if ctx.BooleanLiteral() != nil {
		d.TreeType = TreeType_BooleanLiteral
		d.Value = strings.TrimSpace(
			strings.ReplaceAll(ctx.BooleanLiteral().GetText(), "\"", ""),
		)
		d.HexValue = hex.EncodeToString([]byte(d.Value))
	}

	if ctx.NIL() != nil {
		d.TreeType = TreeType_NIL
		d.Value = ctx.NIL().GetText()
		d.HexValue = hex.EncodeToString([]byte(d.Value))
	}

	if ctx.StringLiteral() != nil {
		lit := NewStringLiteral(d.Builder)
		lit.Parse(ctx.StringLiteral())
		d.TreeType = lit.TreeType
		d.Value = lit.Value
		d.HexValue = lit.HexValue
	}

	if numLit := ctx.NumberLiteral(); numLit != nil {
		lit := NewNumberLiteral(d.Builder)
		lit.Parse(numLit)
		d.TreeType = lit.TreeType
		d.Value = lit.Value
		d.HexValue = lit.HexValue
	}
}

type StringLiteral struct {
	*Builder
	Id              int32
	Src             SrcPos
	TreeType        TreeType
	Value, HexValue string
}

func NewStringLiteral(b *Builder) *StringLiteral {
	return &StringLiteral{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (d *StringLiteral) Parse(ctx needle.IStringLiteralContext) {
	if ctx.InterpretedStringLiteral() != nil {
		d.TreeType = TreeType_InterpretedStringLiteral
		d.Value = strings.TrimSpace(
			strings.ReplaceAll(ctx.InterpretedStringLiteral().GetText(), "\"", ""),
		)
		d.HexValue = hex.EncodeToString([]byte(d.Value))
	}
	if ctx.RawStringLiteral() != nil {
		d.TreeType = TreeType_RawStringLiteral
		d.Value = strings.TrimSpace(
			strings.ReplaceAll(ctx.RawStringLiteral().GetText(), "\"", ""),
		)
		d.HexValue = hex.EncodeToString([]byte(d.Value))
	}
}

type NumberLiteral struct {
	*Builder
	Id              int32
	Src             SrcPos
	TreeType        TreeType
	Value, HexValue string
}

func NewNumberLiteral(b *Builder) *NumberLiteral {
	return &NumberLiteral{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (d *NumberLiteral) Parse(ctx needle.INumberLiteralContext) {
	d.Src = NewSrcPos(ctx)
	if ctx.FloatLiteral() != nil {
		d.TreeType = TreeType_FloatLiteral
		d.Value = ctx.FloatLiteral().GetText()
	}
	if ctx.DecimalLiteral() != nil {
		d.TreeType = TreeType_DecimalLiteral
		d.Value = ctx.DecimalLiteral().GetText()
	}
	if ctx.BinaryLiteral() != nil {
		d.TreeType = TreeType_BinaryLiteral
		d.Value = ctx.BinaryLiteral().GetText()
	}
	if ctx.OctalLiteral() != nil {
		d.TreeType = TreeType_OctalLiteral
		d.Value = ctx.OctalLiteral().GetText()
	}
	if ctx.HexLiteral() != nil {
		d.TreeType = TreeType_HexLiteral
		d.Value = ctx.HexLiteral().GetText()
	}
	d.HexValue = hex.EncodeToString([]byte(d.Value))
}
