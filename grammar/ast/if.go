package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type IfStmt struct {
	*Builder
	Id        int32
	Src       SrcPos
	Kind      TreeType
	IfBody    []*IfBody
	ElseBlock *Block
}

func NewIfStmt(b *Builder) *IfStmt {
	return &IfStmt{
		Builder: b,
		Id:      b.GetReferId(),
		Kind:    TreeType_Kind_ControlStmt,
	}
}

func (i *IfStmt) Parse(ctx needle.IIfStmtContext) {
	i.Src = NewSrcPos(ctx)
	for _, body := range ctx.AllIfBody() {
		ifBody := NewIfBody(i.Builder)
		ifBody.Parse(body)
		i.IfBody = append(i.IfBody, ifBody)
	}

	if ctx.ElseBody() != nil {
		elseBody := NewElseBody(i.Builder)
		elseBody.Parse(ctx.ElseBody())
		i.ElseBlock = elseBody.Block
	}
}

type IfBody struct {
	*Builder
	Id        int32
	Src       SrcPos
	TreeType  TreeType
	Condition *Expr
	Block     *Block
}

func NewIfBody(b *Builder) *IfBody {
	return &IfBody{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_IfStmt,
	}
}

func (i *IfBody) Parse(ctx needle.IIfBodyContext) {
	i.Src = NewSrcPos(ctx)
	expr := NewExpr(i.Builder)
	expr.Parse(ctx.Expr())
	i.Condition = expr

	block := NewBlock(i.Builder)
	block.Parse(ctx.Block())
	i.Block = block
}

type ElseBody struct {
	*Builder
	Id       int32
	Src      SrcPos
	TreeType TreeType
	Block    *Block
}

func NewElseBody(b *Builder) *ElseBody {
	return &ElseBody{
		Builder:  b,
		Id:       b.GetReferId(),
		TreeType: TreeType_ElseStmt,
	}
}

func (e *ElseBody) Parse(ctx needle.IElseBodyContext) {
	e.Src = NewSrcPos(ctx)
	block := NewBlock(e.Builder)
	block.Parse(ctx.Block())
	e.Block = block
}
