package compiler

import (
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
)

// Token is the incoming program is implemented in this file. It is the first phase of compilation
// where the incoming text is divided into a sequence of lexemes.
type Token uint

// represents the lexical token type of the program.
const (
	UNKNOWN    Token = iota
	DELIMITER        // Delimiters
	OPERATOR         // Operators
	NUMBER           // integer or float
	IDENTIFIER       // IDENTIFIER, including KEYWORD, TYPENAME, EXTEND
	NEWLINE          // Line translation
	LITERAL          // string or char
	COMMENT          // Comment
	KEYWORD          // keyword of IDENTIFIER
	TYPENAME         // name of the type of IDENTIFIER
	EXTEND           // Referring to an external variable of IDENTIFIER
)

// The list of delimiters for DELIMITER.
const (
	LPAREN = DELIMITER | (iota+1)<<8
	RPAREN
	COMMA
	DOT
	COLON
	EQ
	LBRACK
	RBRACK
	LBRACE
	RBRACE
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

// The list of operators for OPERATOR.
const (
	Not = OPERATOR | (iota+1)<<8
	Mul
	Add
	Sub
	Quo
	MOD
	Less
	Great
	Assign
	NotEq
	And
	LessEq
	EqEq
	GrEq
	Or
	BitAnd
	BitOr
	BitXor
	LSHIFT
	RSHIFT
	AddEq
	SubEq
	MulEq
	DivEq
	ModEq
	LshEq
	RshEq
	AndEq
	OrEq
	XorEq
	Inc
	Dec
)

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

// The list of keyword identifiers for IDENTIFIER.
const (
	CONTRACT = KEYWORD | (iota+1)<<8
	FUNC
	RETURN
	IF
	ELIF
	ELSE
	WHILE
	TRUE
	FALSE
	VAR
	TX
	SETTINGS
	BREAK
	CONTINUE
	ERRWARNING
	ERRINFO
	NIL
	ACTION
	CONDITIONS
	TAIL
	ERROR
)

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
	"data":       TX,
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

// data types for parameters and variables for Type.
const (
	BOOL = TYPENAME | (iota+1)<<8
	BYTES
	INT
	ADDRESS
	ARRAY
	MAP
	MONEY
	FLOAT
	STRING
	FILE
)

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
	defaultValue := reflect.Zero(t).Interface()
	return defaultValue
}

// d
var tokenToString = map[Token]string{
	UNKNOWN:    "UNKNOWN",
	DELIMITER:  "DELIMITER",
	OPERATOR:   "OPERATOR",
	NUMBER:     "NUMBER",
	IDENTIFIER: "IDENTIFIER",
	NEWLINE:    "NEWLINE",
	LITERAL:    "LITERAL",
	COMMENT:    "COMMENT",
	KEYWORD:    "KEYWORD",
	TYPENAME:   "TYPENAME",
	EXTEND:     "EXTEND",
}

// Lookup maps an identifier to its keyword token
func Lookup(ident string) (Token, bool) {
	tok, ok := KeywordValue[ident]
	return tok, ok
}

// contains checks if a token is in a list of tokens
func (tok Token) contains(list []Token) bool {
	for _, t := range list {
		if t == tok {
			return true
		}
	}
	return false
}

func (tok Token) String() string {
	s := ""
	f := func(m map[string]Token) {
		for op, t := range m {
			if t == tok {
				s = op
			}
		}
	}
	if 0 <= tok {
		s = tokenToString[tok]
		switch tok & 0xff {
		case TYPENAME:
			f(TypeNameValue)
		case KEYWORD:
			f(KeywordValue)
		case OPERATOR:
			f(op2Token)
		case DELIMITER:
			f(delimiter2Token)
		}
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
