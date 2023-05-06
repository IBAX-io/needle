package compile

import (
	"fmt"
	"testing"
)

func TestNewLexeme(t *testing.T) {
	lexer, err := NewLexer([]rune(`$shift  = 23`))
	t.Log(err)
	for _, l := range lexer {
		fmt.Printf(" %v  %v [%d:%d]\n", l.Type, l.Value, l.Line, l.Column)
	}
}
