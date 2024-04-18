package ast

import (
	needle "github.com/IBAX-io/needle/grammar/dist-go"
)

type StatementList struct {
	*Builder
	Statements []*Statement
}

func NewStatementList(b *Builder) *StatementList {
	return &StatementList{
		Builder: b,
	}
}

func (s *StatementList) Parse(ctx needle.IStatementListContext) {
	for _, context := range ctx.AllStatement() {
		stmt := NewStatement(s.Builder)
		stmt.Parse(context)
		s.Statements = append(s.Statements, stmt)
	}
}

type Statement struct {
	*Builder
	Statement NodeType
}

func NewStatement(b *Builder) *Statement {
	return &Statement{
		Builder: b,
	}
}

func (s *Statement) Parse(ctx needle.IStatementContext) {
	if ctx.Block() != nil {
		block := NewBlock(s.Builder)
		block.Parse(ctx.Block())
		s.Statement = block
	}
	if ctx.SimpleStmt() != nil {
		simpleStmt := NewSimpleStmt(s.Builder)
		simpleStmt.Parse(ctx.SimpleStmt())
		s.Statement = simpleStmt.Statement
	}

	if ctx.VarDef() != nil {
		varDef := NewVarDef(s.Builder)
		varDef.Parse(ctx.VarDef())
		s.Statement = varDef
	}
	if ctx.IfStmt() != nil {
		ifStmt := NewIfStmt(s.Builder)
		ifStmt.Parse(ctx.IfStmt())
		s.Statement = ifStmt
	}
	if ctx.WhileStmt() != nil {
		whileStmt := NewWhileStmt(s.Builder)
		whileStmt.Parse(ctx.WhileStmt())
		s.Statement = whileStmt
	}
	if ctx.BreakStmt() != nil {
		breakStmt := NewBreakStmt(s.Builder)
		breakStmt.Parse(ctx.BreakStmt())
		s.Statement = breakStmt
	}
	if ctx.ContinueStmt() != nil {
		continueStmt := NewContinueStmt(s.Builder)
		continueStmt.Parse(ctx.ContinueStmt())
		s.Statement = continueStmt
	}
	if ctx.ReturnStmt() != nil {
		returnStmt := NewReturnStmt(s.Builder)
		returnStmt.Parse(ctx.ReturnStmt())
		s.Statement = returnStmt
	}
	if ctx.ErrorStmt() != nil {
		errorStmt := NewErrorStmt(s.Builder)
		errorStmt.Parse(ctx.ErrorStmt())
		s.Statement = errorStmt
	}
}

type SimpleStmt struct {
	*Builder

	Statement NodeType
}

func NewSimpleStmt(b *Builder) *SimpleStmt {
	return &SimpleStmt{
		Builder: b,
	}
}

func (s *SimpleStmt) Parse(ctx needle.ISimpleStmtContext) {
	if ctx.Expr() != nil {
		expr := NewExpr(s.Builder)
		expr.Parse(ctx.Expr())
		s.Statement = expr
	}
	if ctx.IncDecStmt() != nil {
		incDecStmt := NewIncDecStmt(s.Builder)
		incDecStmt.Parse(ctx.IncDecStmt())
		s.Statement = incDecStmt
	}
	if ctx.Assignment() != nil {
		assignment := NewAssignment(s.Builder)
		assignment.Parse(ctx.Assignment())
		s.Statement = assignment
	}
	if ctx.AssignMapArrStmt() != nil {
		assignMapArrStmt := NewAssignMapArrStmt(s.Builder)
		assignMapArrStmt.Parse(ctx.AssignMapArrStmt())
		s.Statement = assignMapArrStmt
	}
}
