package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type IfStmt struct {
	*Builder

	IfBody    []*IfBody
	ElseBlock *Block
}

func NewIfStmt(b *Builder) *IfStmt {
	return &IfStmt{
		Builder: b,
	}
}

func (i *IfStmt) Parse(ctx needle.IIfStmtContext) {
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

	Condition *Expr
	Block     *Block
}

func NewIfBody(b *Builder) *IfBody {
	return &IfBody{
		Builder: b,
	}
}

func (i *IfBody) Parse(ctx needle.IIfBodyContext) {
	expr := NewExpr(i.Builder)
	expr.Parse(ctx.Expr())
	i.Condition = expr

	block := NewBlock(i.Builder)
	block.Parse(ctx.Block())
	i.Block = block
}

type ElseBody struct {
	*Builder

	Block *Block
}

func NewElseBody(b *Builder) *ElseBody {
	return &ElseBody{
		Builder: b,
	}
}

func (e *ElseBody) Parse(ctx needle.IElseBodyContext) {
	block := NewBlock(e.Builder)
	block.Parse(ctx.Block())
	e.Block = block
}
