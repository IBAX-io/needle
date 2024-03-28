package compiler

import "testing"

func TestToken_ToString(t *testing.T) {
	var toksBase, delimiters, operators, keyword, typename []Token
	for tok := range EXTEND {
		toksBase = append(toksBase, tok)
	}
	for i := LPAREN; i < RBRACE; i++ {
		delimiters = append(delimiters, i)
	}
	for i := Not; i < Dec; i++ {
		operators = append(operators, i)
	}
	for i := CONTRACT; i < ERROR; i++ {
		keyword = append(keyword, i)
	}
	for i := BOOL; i < FILE; i++ {
		typename = append(typename, i)
	}

	tests := []struct {
		name string
		tok  []Token
		want []string
	}{
		{"case_UNKNOWN_EXTEND", toksBase, []string{
			"UNKNOWN", "DELIMITER", "OPERATOR", "NUMBER",
			"IDENTIFIER", "NEWLINE", "LITERAL", "COMMENT", "KEYWORD", "TYPENAME", "EXTEND",
		}},
		{"case_DELIMITER_LPAREN_RBRACE", delimiters, []string{
			"(", ")", ",", ".", ":", "=", "[", "]", "{", "}",
		}},
		{"case_OPERATOR_NOT_DEC", operators, []string{
			"!", "*", "+", "-", "/", "%", "<", ">", "=", "!=", "&&", "<=", "==", ">=", "||", "&", "|", "^", "<<", ">>",
			"+=", "-=", "*=", "/=", "%=", "<<=", ">>=", "&=", "|=", "^=", "++", "--",
		}},
		{"case_KEYWORD_CONTRACT_ERROR", keyword, []string{
			"contract", "func", "return", "if", "elif", "else", "while", "true", "false", "var", "data", "settings", "break",
			"continue", "warning", "info", "nil", "action", "conditions", "...", "error",
		}},
		{"case_TYPENAME_BOOL_FILE", typename, []string{
			"bool", "bytes", "int", "address", "array", "map", "money", "float", "string", "file",
		}},
		{
			"case_not_found", []Token{Token(20)}, []string{"Token(20)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, tok := range tt.tok {
				if got := tok.ToString(); got != tt.want[i] {
					t.Errorf("ToString() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
