package compile

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

	KEYWORD  // keyword of IDENTIFIER
	TYPENAME // name of the type of IDENTIFIER
	EXTEND   // Referring to an external variable of IDENTIFIER
)

// Delimiters for Delimiter.
const (
	LPAREN = DELIMITER | ((iota + 1) << 8)
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

var delimiter2Token = map[rune]Token{
	'(': LPAREN,
	')': RPAREN,
	',': COMMA,
	'.': DOT,
	':': COLON,
	'=': EQ,
	'[': LBRACK,
	']': RBRACK,
	'{': LBRACE,
	'}': RBRACE,
}

// Operations for Operator.
const (
	Not = OPERATOR | ((iota + 1) << 8)
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
	CONTRACT = KEYWORD | ((iota + 1) << 8)
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
	`contract`:   CONTRACT,
	`func`:       FUNC,
	`return`:     RETURN,
	`if`:         IF,
	`elif`:       ELIF,
	`else`:       ELSE,
	`while`:      WHILE,
	`true`:       TRUE,
	`false`:      FALSE,
	`var`:        VAR,
	`data`:       TX,
	`settings`:   SETTINGS,
	`break`:      BREAK,
	`continue`:   CONTINUE,
	`warning`:    ERRWARNING,
	`info`:       ERRINFO,
	`nil`:        NIL,
	`action`:     ACTION,
	`conditions`: CONDITIONS,
	`...`:        TAIL,
	`error`:      ERROR,
}
var KeywordStr = map[Token]string{
	CONTRACT:   `contract`,
	FUNC:       `func`,
	RETURN:     `return`,
	IF:         `if`,
	ELIF:       `elif`,
	ELSE:       `else`,
	WHILE:      `while`,
	TRUE:       `true`,
	FALSE:      `false`,
	VAR:        `var`,
	TX:         `data`,
	SETTINGS:   `settings`,
	BREAK:      `break`,
	CONTINUE:   `continue`,
	ERRWARNING: `warning`,
	ERRINFO:    `info`,
	NIL:        `nil`,
	ACTION:     `action`,
	CONDITIONS: `conditions`,
	TAIL:       `...`,
	ERROR:      `error`,
}

func Keyword2Str(t Token) string {
	return KeywordStr[t]
}

// data types for parameters and variables for Type.
const (
	BOOL = TYPENAME | ((iota + 1) << 8)
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
	`bool`:    BOOL,
	`bytes`:   BYTES,
	`int`:     INT,
	`address`: ADDRESS,
	`array`:   ARRAY,
	`map`:     MAP,
	`money`:   MONEY,
	`float`:   FLOAT,
	`string`:  STRING,
	`file`:    FILE,
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
	STRING:  reflect.TypeOf(``),
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

func SupportedType(fieldType reflect.Type) bool {
	for _, t := range TypeNameReflect {
		if t == fieldType {
			return true
		}
	}
	return false
}

var tokenToString = map[Token]string{
	UNKNOWN: `UNKNOWN`,
	//basic token
	DELIMITER:  `DELIMITER`,
	OPERATOR:   `OPERATOR`,
	NUMBER:     `NUMBER`,
	IDENTIFIER: `IDENTIFIER`,
	NEWLINE:    `NEWLINE`,
	LITERAL:    `LITERAL`,
	COMMENT:    `COMMENT`,
	KEYWORD:    `KEYWORD`,
	TYPENAME:   `TYPENAME`,
	EXTEND:     `EXTEND`,

	//system
	LPAREN: `LPAREN`,
	RPAREN: `RPAREN`,
	COMMA:  `COMMA`,
	DOT:    `DOT`,
	COLON:  `COLON`,
	EQ:     `EQ`,
	LBRACK: `LBRACK`,
	RBRACK: `RBRACK`,
	LBRACE: `LBRACE`,
	RBRACE: `RBRACE`,

	//operator
	Not:    `Not`,
	Mul:    `Mul`,
	Add:    `Add`,
	Sub:    `Sub`,
	Quo:    `Quo`,
	MOD:    `MOD`,
	Less:   `Less`,
	Great:  `Great`,
	Assign: `Assign`,
	NotEq:  `NotEq`,
	And:    `And`,
	LessEq: `LessEq`,
	EqEq:   `EqEq`,
	GrEq:   `GrEq`,
	Or:     `Or`,
	BitAnd: `BitAnd`,
	BitOr:  `BitOr`,
	BitXor: `BitXor`,
	LSHIFT: `LSHIFT`,
	RSHIFT: `RSHIFT`,
	AddEq:  `AddEq`,
	SubEq:  `SubEq`,
	MulEq:  `MulEq`,
	DivEq:  `DivEq`,
	ModEq:  `ModEq`,
	LshEq:  `LshEq`,
	RshEq:  `RshEq`,
	AndEq:  `AndEq`,
	OrEq:   `OrEq`,
	XorEq:  `XorEq`,
	Inc:    `Inc`,
	Dec:    `Dec`,

	//keyword
	CONTRACT:   `CONTRACT`,
	FUNC:       `FUNC`,
	RETURN:     `RETURN`,
	IF:         `IF`,
	ELIF:       `ELIF`,
	ELSE:       `ELSE`,
	WHILE:      `WHILE`,
	TRUE:       `TRUE`,
	FALSE:      `FALSE`,
	VAR:        `VAR`,
	TX:         `DATA`,
	SETTINGS:   `SETTINGS`,
	BREAK:      `BREAK`,
	CONTINUE:   `CONTINUE`,
	ERRWARNING: `ERRWARNING`,
	ERRINFO:    `ERRINFO`,
	NIL:        `NIL`,
	ACTION:     `ACTION`,
	CONDITIONS: `CONDITIONS`,
	TAIL:       `TAIL`,
	ERROR:      `ERROR`,

	//typename
	BOOL:    `BOOL`,
	BYTES:   `BYTES`,
	INT:     `INT`,
	ADDRESS: `ADDRESS`,
	ARRAY:   `ARRAY`,
	MAP:     `MAP`,
	MONEY:   `MONEY`,
	FLOAT:   `FLOAT`,
	STRING:  `STRING`,
	FILE:    `FILE`,
}

// Lookup maps an identifier to its keyword token
func Lookup(ident string) (Token, bool) {
	tok, ok := KeywordValue[ident]
	return tok, ok
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok {
		s = tokenToString[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
