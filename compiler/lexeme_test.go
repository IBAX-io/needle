package compiler

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLexeme(t *testing.T) {
	input := []rune("sd //*sds*/ nil shift int $abc true + action func action if elif")
	lexer, err := NewLexer(input)
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range lexer {
		fmt.Printf("lexer: %v %v [%s]\n", l.Type, l.Value, l.Position())
	}
}

func TestNewLexer(t *testing.T) {
	type args struct {
		input []rune
	}
	tests := []struct {
		name    string
		args    args
		want    Lexemes
		wantErr bool
	}{
		{"case1", args{[]rune(`var conditions shift int $abc func action if elif `)}, nil, false},
		{"case2", args{[]rune(`"asd" ...`)}, nil, false},
		{"case3", args{[]rune(`()[]{},:;`)}, nil, false},
		{"case4", args{[]rune(`+= ++ + -= -- - *= * /= / %= % == = != ! <<= <= << < >>= >= >> > || |= | && &= & ^= ^`)}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLexer(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLexer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, l := range got {
				fmt.Printf("%v %v [%d:%d]\n", l.Type, l.Value, l.Line, l.Column)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLexer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
