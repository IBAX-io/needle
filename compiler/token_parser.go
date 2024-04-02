package compiler

import (
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
)

// delimiter2Token is a map of delimiters to tokens
var delimiter2Token = map[string]Token{
	"(": LPAREN,
	")": RPAREN,
	",": COMMA,
	".": DOT,
	":": COLON,
	"=": EQ,
	"[": LBRACK,
	"]": RBRACK,
	"{": LBRACE,
	"}": RBRACE,
}

// op2Token is a map of operators to tokens
var op2Token = map[string]Token{
	"!":   Not,
	"*":   Mul,
	"+":   Add,
	"-":   Sub,
	"/":   Quo,
	"<":   Less,
	">":   Great,
	"=":   Assign,
	"!=":  NotEq,
	"&&":  And,
	"<=":  LessEq,
	"==":  EqEq,
	">=":  GrEq,
	"||":  Or,
	"&":   BitAnd,
	"|":   BitOr,
	"^":   BitXor,
	"%":   MOD,
	"<<":  LSHIFT,
	">>":  RSHIFT,
	"+=":  AddEq,
	"-=":  SubEq,
	"*=":  MulEq,
	"/=":  DivEq,
	"%=":  ModEq,
	"<<=": LshEq,
	">>=": RshEq,
	"&=":  AndEq,
	"|=":  OrEq,
	"^=":  XorEq,
	"++":  Inc,
	"--":  Dec,
}

// KeywordValue is a map of keywords to tokens
var KeywordValue = map[string]Token{
	"contract":   CONTRACT,
	"func":       FUNC,
	"return":     RETURN,
	"if":         IF,
	"elif":       ELIF,
	"else":       ELSE,
	"while":      WHILE,
	"true":       TRUE,
	"false":      FALSE,
	"var":        VAR,
	"data":       FIELD,
	"settings":   SETTINGS,
	"break":      BREAK,
	"continue":   CONTINUE,
	"warning":    ERRWARNING,
	"info":       ERRINFO,
	"nil":        NIL,
	"action":     ACTION,
	"conditions": CONDITIONS,
	"...":        TAIL,
	"error":      ERROR,
}

// TypeNameValue is a map of types to tokens
var TypeNameValue = map[string]Token{
	"bool":    BOOL,
	"bytes":   BYTES,
	"int":     INT,
	"address": ADDRESS,
	"array":   ARRAY,
	"map":     MAP,
	"money":   MONEY,
	"float":   FLOAT,
	"string":  STRING,
	"file":    FILE,
}

// TypeNameReflect is a map of types to reflect.Type
var TypeNameReflect = map[Token]reflect.Type{
	BOOL:    reflect.TypeOf(true),
	BYTES:   reflect.TypeOf([]byte{}),
	INT:     reflect.TypeOf(int64(0)),
	ADDRESS: reflect.TypeOf(int64(0)),
	ARRAY:   reflect.TypeOf([]any{}),
	MAP:     reflect.TypeOf(&Map{}),
	MONEY:   reflect.TypeOf(decimal.Zero),
	FLOAT:   reflect.TypeOf(0.0),
	STRING:  reflect.TypeOf(""),
	FILE:    reflect.TypeOf(&Map{}),
}

// GetFieldDefaultValue returns default value for field type
func GetFieldDefaultValue(fieldType Token) any {
	t, ok := TypeNameReflect[fieldType]
	if !ok {
		return nil
	}
	return reflect.Zero(t).Interface()
}

// Lookup maps an identifier to its keyword token
func Lookup(ident string) (Token, bool) {
	tok, ok := KeywordValue[ident]
	return tok, ok
}

func (tok Token) ReflectType() reflect.Type {
	return TypeNameReflect[tok]
}

func (tok Token) EqualsType(t any) bool {
	return tok.ReflectType() == reflect.TypeOf(t)
}

func (tok Token) ToString() string {
	var s string
	var tokenMap map[string]Token
	if 0 <= tok {
		s = tok.String()
		switch tok.Kind() {
		default:
		case TYPENAME:
			tokenMap = TypeNameValue
		case KEYWORD:
			tokenMap = KeywordValue
		case OPERATOR:
			tokenMap = op2Token
		case DELIMITER:
			tokenMap = delimiter2Token
		}
	}
	for op, t := range tokenMap {
		if t == tok {
			s = op
		}
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// Contains checks if a token is in a list of tokens
func (tok Token) Contains(list []Token) bool {
	for _, t := range list {
		if t == tok {
			return true
		}
	}
	return false
}

// Kind returns the kind of token
func (tok Token) Kind() Token {
	if UNKNOWN <= tok && tok <= EXTEND {
		return tok
	}
	return (tok & 0xFF00) >> 8
}
