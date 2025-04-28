package ir

import "github.com/IBAX-io/needle/grammar/ast"

type Literal struct {
	Type     string
	Value    string
	HexValue string
}

func (b *Builder) processLiteral(ast *ast.Literal) *Literal {
	return &Literal{
		Type:     ast.TreeType.String(),
		Value:    ast.Value,
		HexValue: ast.HexValue,
	}
}
