package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Arguments struct {
	*Builder

	ArgumentsList *ArgumentsList
}

func NewArguments(b *Builder) *Arguments {
	return &Arguments{
		Builder: b,
	}
}

func (a *Arguments) Parse(ctx needle.IArgumentsContext) {
	if ctx.ArgumentsList() != nil {
		argumentsList := NewArgumentsList(a.Builder)
		argumentsList.Parse(ctx.ArgumentsList())
		a.ArgumentsList = argumentsList
	}
}

type ArgumentsList struct {
	*Builder

	Expressions     []*Expr
	InitMapArrStmts []*InitMapArrStmt
}

func NewArgumentsList(b *Builder) *ArgumentsList {
	return &ArgumentsList{
		Builder: b,
	}
}

func (a *ArgumentsList) Parse(ctx needle.IArgumentsListContext) {
	for _, argument := range ctx.AllExpr() {
		expression := NewExpr(a.Builder)
		expression.Parse(argument)
		a.Expressions = append(a.Expressions, expression)
	}

	for _, context := range ctx.AllInitMapArrStmt() {
		initMapArrStmt := NewInitMapArrStmt(a.Builder)
		initMapArrStmt.Parse(context)
		a.InitMapArrStmts = append(a.InitMapArrStmts, initMapArrStmt)
	}
}
