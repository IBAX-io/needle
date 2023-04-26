package compile

import (
	"fmt"
	"testing"
)

func TestNewLexeme(t *testing.T) {
	lexer, err := NewLexer([]rune(`contract;action;var s string;s = "abc"`))
	fmt.Println(lexer, err)
}
