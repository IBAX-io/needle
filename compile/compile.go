package compile

import (
	"fmt"
)

func CompileBlock(input []rune, conf *CompConfig) (*CodeBlock, error) {
	lexer, err := NewLexer(input)
	if err != nil {
		return nil, fmt.Errorf("lexer error: %w", err)
	}
	parser := NewParser(lexer, conf)
	block, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return block, nil
}

// ContractsList returns list of contracts names from source of code
func ContractsList(value string) ([]string, error) {
	lex, err := NewLexer([]rune(value))
	if err != nil {
		return make([]string, 0), err
	}
	return lex.nameList(CONTRACT, 0), nil
}

// FunctionsList returns list of functions names from source of code
func FunctionsList(value string) ([]string, error) {
	lex, err := NewLexer([]rune(value))
	if err != nil {
		return make([]string, 0), err
	}
	return lex.nameList(FUNC, 0), nil
}
