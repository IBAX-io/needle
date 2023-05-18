package compile

import "fmt"

func CompileBlock(input []rune, ext *ExtendData) (*CodeBlock, error) {
	lexer, err := NewLexer(input)
	if err != nil {
		return nil, fmt.Errorf("lexer: %w", err)
	}
	return NewParser(lexer, ext)
}

// ContractsList returns list of contracts names from source of code
func ContractsList(value string) ([]string, error) {
	names := make([]string, 0)
	lexemes, err := NewLexer([]rune(value))
	if err != nil {
		return names, err
	}
	var level int
	for i, lexeme := range lexemes {
		switch lexeme.Type {
		case LBRACE:
			level++
		case RBRACE:
			level--
		case CONTRACT, FUNC:
			if level == 0 && i+1 < len(lexemes) && lexemes[i+1].Type == IDENTIFIER {
				names = append(names, lexemes[i+1].Value.(string))
			}
		}
	}
	return names, nil
}
