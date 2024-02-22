package compiler

import (
	"fmt"
)

// CompileBlock compiles a block of input runes into a CodeBlock using the provided CompConfig.
func CompileBlock(input []rune, conf *CompConfig) (*CodeBlock, error) {
	lexer, err := NewLexer(input)
	if err != nil {
		return nil, fmt.Errorf("lexer error: %w", err)
	}
	parser := NewParser(lexer, conf)
	return parser.Parse()
}

// ContractsList parses the given value string and returns a list of contract names found.
func ContractsList(value string) ([]string, error) {
	lex, err := NewLexer([]rune(value))
	if err != nil {
		return []string{}, err
	}
	return lex.nameList(CONTRACT), nil
}

// FunctionsList parses the given value string and returns a list of function names found.
func FunctionsList(value string) ([]string, error) {
	lex, err := NewLexer([]rune(value))
	if err != nil {
		return make([]string, 0), err
	}
	return lex.nameList(FUNC), nil
}
