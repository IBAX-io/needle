package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type DataDef struct {
	*Builder
	Src      SrcPos
	StmtType string

	Parts []*DataPartList
}

func NewDataDef(b *Builder) *DataDef {
	return &DataDef{
		Builder:  b,
		StmtType: "DataDef",
	}
}

func (d *DataDef) Parse(ctx *needle.DataDefContext) {
	d.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
	for _, part := range ctx.AllDataPartList() {
		list := NewDataPartList()
		list.Parse(part)
		d.Parts = append(d.Parts, list)
	}
}

type DataPartList struct {
	Src      SrcPos
	Name     string
	Typename string
	Tag      string
}

func NewDataPartList() *DataPartList {
	return &DataPartList{}
}

func (l *DataPartList) Parse(ctx needle.IDataPartListContext) {
	l.Src = SrcPos{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
		Start:  ctx.GetStart().GetStart(),
		End:    ctx.GetStop().GetStop(),
		Length: ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1,
	}
	l.Name = ctx.Identifier().GetText()
	l.Typename = ctx.TypeName().GetText()
	if ctx.GetDataTag() != nil {
		l.Tag = ctx.GetDataTag().GetText()
	}
}
