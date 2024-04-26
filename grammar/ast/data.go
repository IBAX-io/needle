package ast

import (
	"strings"

	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type DataDef struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType

	DataParts []*DataPart
}

func NewDataDef(b *Builder) *DataDef {
	return &DataDef{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_DataDef,
	}
}

func (d *DataDef) Parse(ctx *needle.DataDefContext) {
	d.Src = NewSrcPos(ctx)
	for _, part := range ctx.AllDataPart() {
		list := NewDataPart(d.Builder)
		list.Parse(part)
		d.DataParts = append(d.DataParts, list)
	}
}

type DataPart struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Name     string
	Typename string
	Tag      string
}

func NewDataPart(b *Builder) *DataPart {
	return &DataPart{
		Builder:  b,
		Id:       b.GetNextId(),
		TreeType: TreeType_DataPart,
	}
}

func (l *DataPart) Parse(ctx needle.IDataPartContext) {
	l.Src = NewSrcPos(ctx)
	l.Name = ctx.Identifier().GetText()
	l.Typename = ctx.TypeName().GetText()
	if ctx.GetDataTag() != nil {
		l.Tag = strings.TrimSpace(
			strings.ReplaceAll(ctx.GetDataTag().GetText(), "\"", ""),
		)
	}
}
