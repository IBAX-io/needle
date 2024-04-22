package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type SettingsDef struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType

	SettingValue []*Literal
	SettingsName []string
}

func NewSettingsDef(b *Builder) *SettingsDef {
	return &SettingsDef{
		Builder:      b,
		Id:           b.GetReferId(),
		TreeType:     TreeType_SettingsDef,
		SettingValue: make([]*Literal, 0),
		SettingsName: make([]string, 0),
	}
}

func (d *SettingsDef) Parse(ctx *needle.SettingsDefContext) {
	d.Src = NewSrcPos(ctx)

	for i, part := range ctx.AllIdentifier() {
		valueCtx := ctx.AllLiteral()[i]
		lit := NewLiteral(d.Builder)
		lit.Parse(valueCtx)
		d.SettingValue = append(d.SettingValue, lit)
		d.SettingsName = append(d.SettingsName, part.GetText())
	}
}
