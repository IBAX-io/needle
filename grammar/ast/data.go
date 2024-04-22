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

	Parts []*DataPartList
}

func NewDataDef(b *Builder) *DataDef {
	return &DataDef{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_DataDef,
	}
}

func (d *DataDef) Parse(ctx *needle.DataDefContext) {
	d.Src = NewSrcPos(ctx)
	for _, part := range ctx.AllDataPartList() {
		list := NewDataPartList(d.Builder)
		list.Parse(part)
		d.Parts = append(d.Parts, list)
	}
}

type DataPartList struct {
	*Builder
	Id       int32
	Src      SrcPos
	Name     string
	Typename string
	Tag      string
}

func NewDataPartList(b *Builder) *DataPartList {
	return &DataPartList{
		Builder: b,
		Id:      b.GetReferId(),
	}
}

func (l *DataPartList) Parse(ctx needle.IDataPartListContext) {
	l.Src = NewSrcPos(ctx)
	l.Name = ctx.Identifier().GetText()
	l.Typename = ctx.TypeName().GetText()
	if ctx.GetDataTag() != nil {
		l.Tag = strings.TrimSpace(
			strings.ReplaceAll(ctx.GetDataTag().GetText(), "\"", ""),
		)
	}
}
