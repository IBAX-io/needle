package ast

import needle "github.com/IBAX-io/needle/grammar/dist-go"

type ArrayStmt struct {
	*Builder

	ArrayList *ArrayList
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
		a.ArrayList = arrayList
	}
}

type ArrayList struct {
	*Builder

	ArrayValue []*ArrayValue
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
		a.ArrayValue = append(a.ArrayValue, arrayValue)
	}
}

type ArrayValue struct {
	*Builder

	Expr           *Expr
	InitMapArrStmt *InitMapArrStmt
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
		a.Expr = expression
	}
	if ctx.InitMapArrStmt() != nil {
		initMapArrStmt := NewInitMapArrStmt(a.Builder)
		initMapArrStmt.Parse(ctx.InitMapArrStmt())
		a.InitMapArrStmt = initMapArrStmt
	}
}
