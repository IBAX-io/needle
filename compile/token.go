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

func (tok Token) IsBasic() bool { return UNKNOWN < tok && tok < EXTEND }

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
	LPAREN = SYSTEM | ('(' << 8)
	RPAREN = SYSTEM | (')' << 8)
	COMMA  = SYSTEM | (',' << 8)
	DOT    = SYSTEM | ('.' << 8)
	COLON  = SYSTEM | (':' << 8)
	EQ     = SYSTEM | ('=' << 8)
	LBRACK = SYSTEM | ('[' << 8)
	RBRACK = SYSTEM | (']' << 8)
	LBRACE = SYSTEM | ('{' << 8)
	RBRACE = SYSTEM | ('}' << 8)
)

// Operations for Operator.
const (
	Not    Token = 0x0021 // !
	Mul          = 0x002a // *
	Add          = 0x002b // +
	Sub          = 0x002d // -
	Neg          = 0x012d // - unary
	Quo          = 0x002f // / quotient
	Less         = 0x003c // <
	Great        = 0x003e // >
	Assign       = 0x003d // =

	NotEq  = 0x213d // !=
	And    = 0x2626 // &&
	LessEq = 0x3c3d // <=
	EqEq   = 0x3d3d // ==
	GrEq   = 0x3e3d // >=
	Or     = 0x7c7c // ||

	BITAND = 0x0026   // &
	BITOR  = 0x007c   // |
	BITXOR = 0x005e   // ^
	MOD    = 0x0025   // %
	LSHIFT = 0x3c3c   // <<
	RSHIFT = 0x3e3e   // >>
	AddEq  = 0x2b3d   // +=
	SubEq  = 0x2d3d   // -=
	MulEq  = 0x2a3d   // *=
	DivEq  = 0x2f3d   // /=
	ModEq  = 0x253d   // %=
	LshEq  = 0x3c3c3d // <<=
	RshEq  = 0x3e3e3d // >>=
	AndEq  = 0x263d   // &=
	OrEq   = 0x7c3d   // |=
	XorEq  = 0x5e3d   // ^=
	Inc    = 0x2b2b   // ++
	Dec    = 0x2d2d   // --
)

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

const keyFunc = FUNC >> 8
const keyIf = IF >> 8
const keyElse = ELSE >> 8
const keyLBRACE = LBRACE >> 8
const keyRBRACE = RBRACE >> 8

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

var tokenToString = map[Token]string{
	UNKNOWN: `unknown`,
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

	//keyword
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

	//name of data type
	BOOL:    `bool`,
	BYTES:   `bytes`,
	INT:     `int`,
	ADDRESS: `address`,
	ARRAY:   `array`,
	MAP:     `map`,
	MONEY:   `money`,
	FLOAT:   `float`,
	STRING:  `string`,
	FILE:    `file`,

	//delimiters
	LPAREN: `(`,
	RPAREN: `)`,
	COMMA:  `,`,
	DOT:    `.`,
	COLON:  `:`,
	EQ:     `=`,
	LBRACK: `[`,
	RBRACK: `]`,
	LBRACE: `{`,
	RBRACE: `}`,

	//operators
	Not:    `!`,
	Mul:    `*`,
	Add:    `+`,
	Sub:    `-`,
	Neg:    `-`,
	Quo:    `/`,
	Less:   `<`,
	Great:  `>`,
	NotEq:  `!=`,
	And:    `&&`,
	LessEq: `<=`,
	EqEq:   `==`,
	GrEq:   `>=`,
	Or:     `||`,

	MOD:    `% `,
	ModEq:  `%=`,
	BITAND: `&`,
	BITOR:  `|`,
	BITXOR: `^`,

	AddEq: `+=`,
	SubEq: `-=`,
	MulEq: `*=`,
	DivEq: `/=`,
	AndEq: `&=`,
	OrEq:  `|=`,
	XorEq: `^=`,

	Inc:    `++`,
	Dec:    `--`,
	LSHIFT: `<<`,
	RSHIFT: `>>`,
	LshEq:  `<<=`,
	RshEq:  `>>=`,
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
