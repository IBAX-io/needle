package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ArrayStmt struct {
	*Builder
}

func NewArrayStmt(b *Builder) *ArrayStmt {
	return &ArrayStmt{
		Builder: b,
	}
}

func (a *ArrayStmt) Parse(ctx needle.IArrayStmtContext) {
	if ctx.ArrayList() != nil {
		arrayList := NewArrayList(a.Builder)
		arrayList.Parse(ctx.ArrayList())
	}
}

type ArrayList struct {
	*Builder
}

func NewArrayList(b *Builder) *ArrayList {
	return &ArrayList{
		Builder: b,
	}
}

func (a *ArrayList) Parse(ctx needle.IArrayListContext) {
	for _, value := range ctx.AllArrayValue() {
		arrayValue := NewArrayValue(a.Builder)
		arrayValue.Parse(value)
	}
}

type ArrayValue struct {
	*Builder
}

func NewArrayValue(b *Builder) *ArrayValue {
	return &ArrayValue{
		Builder: b,
	}
}

func (a *ArrayValue) Parse(ctx needle.IArrayValueContext) {
	if ctx.Expr() != nil {
		expression := NewExpr(a.Builder)
		expression.Parse(ctx.Expr())
	}
	if ctx.ArrayStmt() != nil {
		arrayStmt := NewArrayStmt(a.Builder)
		arrayStmt.Parse(ctx.ArrayStmt())
	}
	if ctx.MapStmt() != nil {
		mapStmt := NewMapStmt(a.Builder)
		mapStmt.Parse(ctx.MapStmt())
	}
}
