package compile

import (
	"fmt"
	"testing"
)

func TestToken_String(t *testing.T) {
	fmt.Println(FUNC, RBRACE, Add, BOOL, EXTEND, KEYWORD, LITERAL)
}

func TestGetFieldDefaultValue(t *testing.T) {
	for i := 0; i < 11; i++ {
		lexeme := TYPENAME | Token((i+1)<<8)
		defaultValue := GetFieldDefaultValue(lexeme)
		fmt.Printf("%d,Type: %v, DefaultValue: %v\n", i, lexeme.String(), defaultValue)
	}
	name := "address"
	r, _ := TypeNameReflect[TypeNameValue[name]]

	fmt.Println(NewLexeme(TypeNameValue[name], r, 1, 1))
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
		{"TestLookup", args{"data"}, TX, true},
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
