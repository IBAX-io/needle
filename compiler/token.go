package compiler

//go:generate stringer -type Token -output=token_string.go token.go

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
	LPAREN Token = DELIMITER<<8 | iota
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

// The list of operators for OPERATOR.
const (
	Not Token = OPERATOR<<8 | iota
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

// The list of keyword identifiers for IDENTIFIER.
const (
	CONTRACT Token = KEYWORD<<8 | iota
	FUNC
	RETURN
	IF
	ELIF
	ELSE
	WHILE
	TRUE
	FALSE
	VAR
	FIELD
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

// The list of data types for parameters and variables for TYPENAME.
const (
	BOOL Token = TYPENAME<<8 | iota
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
