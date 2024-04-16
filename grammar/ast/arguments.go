package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type Arguments struct {
	*Builder
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
	}
}

type ArgumentsList struct {
	*Builder
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
	}

	for _, context := range ctx.AllInitMapArrStmt() {
		initMapArrStmt := NewInitMapArrStmt(a.Builder)
		initMapArrStmt.Parse(context)
	}
}
