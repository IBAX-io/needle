package ir

import (
	"github.com/IBAX-io/needle/compiler"

	"github.com/IBAX-io/needle/grammar/ast"
)

type Builder struct {
	*ast.Builder
	currentContract *Contract
	CodeBlock       *compiler.CodeBlock
}

type Root struct {
	Owner     *Owner
	Contracts []*Contract
}

type Owner struct {
	StateId uint32
}
