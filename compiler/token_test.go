package compiler

import (
	"fmt"
	"testing"
)

func TestGetFieldDefaultValue(t *testing.T) {
	for i := 1; i < 11; i++ {
		lexeme := TYPENAME<<8 | Token(i)
		defaultValue := GetFieldDefaultValue(lexeme)
		fmt.Printf("%d,Type: %v, DefaultValue: %v\n", i, lexeme.ToString(), defaultValue)
	}
	name := "address"

	fmt.Println(NewLexeme(TYPENAME, NewLexemeToken(TypeNameValue[name]), 1, 1))
}

func TestLookup(t *testing.T) {
	type args struct {
		ident string
	}
	tests := []struct {
		name  string
		args  args
		want  Token
		want1 bool
	}{
		{"TestLookup", args{"int"}, INT, true},
		{"TestLookup", args{"data"}, FIELD, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Lookup(tt.args.ident)
			if got != tt.want {
				t.Errorf("Lookup() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Lookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
