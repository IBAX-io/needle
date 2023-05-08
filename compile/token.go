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
	SYSTEM           // Delimiters
	OPERATOR         // Operators
	NUMBER           // integer or float
	IDENTIFIER       // IDENTIFIER, including KEYWORD,TYPENAME,EXTEND,
	NEWLINE          // Line translation
	LITERAL          // string or char
	COMMENT          // Comment

	KEYWORD  // keyword of IDENTIFIER
	TYPENAME // name of the type of IDENTIFIER
	EXTEND   // Referring to an external variable of IDENTIFIER
)

const Error = 0xff

// flags of lexical states
const (
	flagEnd  = 0
	flagNext = 1
	flagPush = 2
	flagPop  = 4
	flagSkip = 8
)

// Delimiters for SYSTEM.
const (
	LPAREN = SYSTEM | ((iota + 1) << 8)
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

var system2Token = map[rune]Token{
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

const (
	OpUnary = iota + 0x0100
)

// Operations for Operator.
const (
	OpNot = OPERATOR | ((iota + 1) << 8)
	OpMul
	OpAdd
	OpSub
	OpNeg
	OpQuo
	OpLess
	OpGreat
	OpAssign
	OpNotEq
	OpAnd
	OpLessEq
	OpEqEq
	OpGrEq
	OpOr
	OpBITAND
	OpBITOR
	OpBITXOR
	OpMOD
	OpLSHIFT
	OpRSHIFT
	OpAddEq
	OpSubEq
	OpMulEq
	OpDivEq
	OpModEq
	OpLshEq
	OpRshEq
	OpAndEq
	OpOrEq
	OpXorEq
	OpInc
	OpDec
)

var op2Token = map[string]Token{
	"!":   OpNot,
	"*":   OpMul,
	"+":   OpAdd,
	"-":   OpSub,
	"/":   OpQuo,
	"<":   OpLess,
	">":   OpGreat,
	"=":   OpAssign,
	"!=":  OpNotEq,
	"&&":  OpAnd,
	"<=":  OpLessEq,
	"==":  OpEqEq,
	">=":  OpGrEq,
	"||":  OpOr,
	"&":   OpBITAND,
	"|":   OpBITOR,
	"^":   OpBITXOR,
	"%":   OpMOD,
	"<<":  OpLSHIFT,
	">>":  OpRSHIFT,
	"+=":  OpAddEq,
	"-=":  OpSubEq,
	"*=":  OpMulEq,
	"/=":  OpDivEq,
	"%=":  OpModEq,
	"<<=": OpLshEq,
	">>=": OpRshEq,
	"&=":  OpAndEq,
	"|=":  OpOrEq,
	"^=":  OpXorEq,
	"++":  OpInc,
	"--":  OpDec,
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
	MAP:     reflect.TypeOf(map[string]any{}),
	MONEY:   reflect.TypeOf(decimal.Zero),
	FLOAT:   reflect.TypeOf(0.0),
	STRING:  reflect.TypeOf(``),
	FILE:    reflect.TypeOf(map[string]any{}),
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

var tokenToString = map[Token]string{
	UNKNOWN: `UNKNOWN`,
	//basic token
	SYSTEM:     `SYSTEM`,
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
	OpNot:    `OpNot`,
	OpMul:    `OpMul`,
	OpAdd:    `OpAdd`,
	OpSub:    `OpSub`,
	OpQuo:    `OpQuo`,
	OpLess:   `OpLess`,
	OpGreat:  `OpGreat`,
	OpAssign: `OpAssign`,
	OpNotEq:  `OpNotEq`,
	OpAnd:    `OpAnd`,
	OpLessEq: `OpLessEq`,
	OpEqEq:   `OpEqEq`,
	OpGrEq:   `OpGrEq`,
	OpOr:     `OpOr`,
	OpBITAND: `OpBITAND`,
	OpBITOR:  `OpBITOR`,
	OpBITXOR: `OpBITXOR`,
	OpMOD:    `OpMOD`,
	OpLSHIFT: `OpLSHIFT`,
	OpRSHIFT: `OpRSHIFT`,
	OpAddEq:  `OpAddEq`,
	OpSubEq:  `OpSubEq`,
	OpMulEq:  `OpMulEq`,
	OpDivEq:  `OpDivEq`,
	OpModEq:  `OpModEq`,
	OpLshEq:  `OpLshEq`,
	OpRshEq:  `OpRshEq`,
	OpAndEq:  `OpAndEq`,
	OpOrEq:   `OpOrEq`,
	OpXorEq:  `OpXorEq`,
	OpInc:    `OpInc`,
	OpDec:    `OpDec`,

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
	TX:         `TX`,
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
