package compiler

import "testing"

var toksBase, delimiters, operators, keyword, typename []Token

func init() {
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
}

func TestToken_ToString(t *testing.T) {
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

func TestToken_Kind(t *testing.T) {
	tests := []struct {
		name string
		tok  []Token
		want Token
	}{
		// TODO: Add test cases.
		{"case_UNKNOWN", []Token{UNKNOWN}, UNKNOWN},
		{"case_DELIMITER", append(delimiters, DELIMITER), DELIMITER},
		{"case_OPERATOR", append(operators, OPERATOR), OPERATOR},
		{"case_NUMBER", []Token{NUMBER}, NUMBER},
		{"case_IDENTIFIER", []Token{IDENTIFIER}, IDENTIFIER},
		{"case_NEWLINE", []Token{NEWLINE}, NEWLINE},
		{"case_LITERAL", []Token{LITERAL}, LITERAL},
		{"case_COMMENT", []Token{COMMENT}, COMMENT},
		{"case_KEYWORD", append(keyword, KEYWORD), KEYWORD},
		{"case_TYPENAME", append(typename, TYPENAME), TYPENAME},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tok := range tt.tok {
				if got := tok.Kind(); got != tt.want {
					t.Errorf("Kind() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
