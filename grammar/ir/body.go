package ir

import (
	"github.com/IBAX-io/needle/grammar/ast"
)

type Body struct {
	Statements []Statement
}

type Statement interface {
	GetReferId() int32
	GetTreeType() ast.TreeType
	GetSrc() ast.SrcPos
	Parse()
}

func (b *Builder) processBody(ast *ast.Block) *Body {
	body := &Body{}
	for _, stmt := range ast.StatementList.Statements {
		body.Statements = append(body.Statements, b.processStatement(stmt))
	}
	return body
}

func (b *Builder) processStatement(ast *ast.Statement) Statement {
	if ast.VarDef != nil {
		return b.processVarDef(ast.VarDef)
	}
	return nil
}

func (b *Builder) processVarDef(ast *ast.VarDef) Statement {
	v := &Variable{
		ast:  ast,
		Name: make([]string, 0),
		Type: make([]string, 0),
	}
	for i, list := range ast.Parameter.NameList {
		for _, s := range list.Name {
			v.Name = append(v.Name, s)
			v.Type = append(v.Type, ast.Parameter.TypeName[i])
		}
	}
	return v
}

type Variable struct {
	ast  *ast.VarDef
	Name []string
	Type []string
}

func (v *Variable) GetReferId() int32 {
	return v.ast.Id
}

func (v *Variable) GetTreeType() ast.TreeType {
	return v.ast.TreeType
}

func (v *Variable) GetSrc() ast.SrcPos {
	return v.ast.Src
}

func (v *Variable) Parse() {}

type OperatorExpr struct {
	ast       *ast.Expr
	TreeType  ast.TreeType
	IncAndDec bool
	LeftExpr  *ast.Expr
	RightExpr *ast.Expr
}

func (o OperatorExpr) GetReferId() int32 {
	return o.ast.Id
}

func (o OperatorExpr) GetTreeType() ast.TreeType {
	return o.TreeType
}

func (o OperatorExpr) GetSrc() ast.SrcPos {
	return o.ast.Src
}

func (o OperatorExpr) Parse() {
	if op := o.ast.LogicalOp; op != nil {
		o.LeftExpr = op.LeftExpr
		o.RightExpr = op.RightExpr
	}
	if op := o.ast.MulOp; op != nil {
		o.LeftExpr = op.LeftExpr
		o.RightExpr = op.RightExpr
	}
	if op := o.ast.AddOp; op != nil {
		o.LeftExpr = op.LeftExpr
		o.RightExpr = op.RightExpr
	}
	if op := o.ast.RelOp; op != nil {
		o.LeftExpr = op.LeftExpr
		o.RightExpr = op.RightExpr
	}
	if op := o.ast.UnaryOp; op != nil {
		o.RightExpr = op.RightExpr
	}
}
