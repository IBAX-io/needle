package compiler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Lexeme is a lexical token of the program.
type Lexeme struct {
	Type Token
	// Types that are assignable to Value:
	//
	//  *LexemeString
	//  *LexemeNumber
	//  *LexemeToken
	//  *LexemeBoolean
	//  *LexemeNil
	Value  isLexemeValue
	Line   int
	Column int
}

type isLexemeValue interface {
	isLexemeValue()
}

type LexemeString struct {
	Value string
}

func (s *LexemeString) String() string {
	return s.Value
}

type LexemeNumber struct {
	Int64     int64
	Float64   float64
	IsInteger bool
}

func (n *LexemeNumber) String() string {
	if n.IsInteger {
		return strconv.FormatInt(n.Int64, 10)
	}
	return strconv.FormatFloat(n.Float64, 'f', -1, 64)
}

type LexemeToken struct {
	Value Token
}

func (t *LexemeToken) String() string {
	return t.Value.String()
}

type LexemeBoolean struct {
	Value bool
}

func (b *LexemeBoolean) String() string {
	return strconv.FormatBool(b.Value)
}

type LexemeNil struct{}

func (n *LexemeNil) String() string {
	return "nil"
}

func (*LexemeString) isLexemeValue()  {}
func (*LexemeNumber) isLexemeValue()  {}
func (*LexemeToken) isLexemeValue()   {}
func (*LexemeBoolean) isLexemeValue() {}
func (*LexemeNil) isLexemeValue()     {}

func NewLexemeString(value string) *LexemeString {
	return &LexemeString{Value: value}
}

func NewLexemeNumber(value any) *LexemeNumber {
	switch value.(type) {
	case int64:
		return &LexemeNumber{Int64: value.(int64), IsInteger: true}
	case float64:
		return &LexemeNumber{Float64: value.(float64)}
	}
	return &LexemeNumber{}
}

func NewLexemeBoolean(value bool) *LexemeBoolean {
	return &LexemeBoolean{Value: value}
}

func NewLexemeToken(value Token) *LexemeToken {
	return &LexemeToken{Value: value}
}

func NewLexemeNil() *LexemeNil {
	return &LexemeNil{}
}

func (l *Lexeme) ToString() string {
	v, ok := l.Value.(*LexemeString)
	if !ok {
		return ""
	}
	return v.Value
}

func (l *Lexeme) Int64() int64 {
	v, ok := l.Value.(*LexemeNumber)
	if !ok {
		return 0
	}
	return v.Int64
}

func (l *Lexeme) IsInteger() bool {
	v, ok := l.Value.(*LexemeNumber)
	if !ok {
		return false
	}
	return v.IsInteger
}

func (l *Lexeme) Float64() float64 {
	v, ok := l.Value.(*LexemeNumber)
	if !ok {
		return 0
	}
	return v.Float64
}

func (l *Lexeme) Boolean() bool {
	v, ok := l.Value.(*LexemeBoolean)
	if !ok {
		return false
	}
	return v.Value
}

func (l *Lexeme) Token() Token {
	v, ok := l.Value.(*LexemeToken)
	if !ok {
		return UNKNOWN
	}
	return v.Value
}

func (l *Lexeme) Interface() any {
	switch v := l.Value.(type) {
	case *LexemeString:
		return v.Value
	case *LexemeNumber:
		if v.IsInteger {
			return v.Int64
		}
		return v.Float64
	case *LexemeBoolean:
		return v.Value
	case *LexemeToken:
		return v.Value
	case *LexemeNil:
		return nil
	}
	return nil
}

func (l *Lexeme) String() string {
	return fmt.Sprintf("%s", l.Value)
}

// NewLexeme creates a new Lexeme with the given parameters.
func NewLexeme(Type Token, value isLexemeValue, line int, column int) *Lexeme {
	return &Lexeme{Type: Type, Value: value, Line: line, Column: column}
}

// GetLogger returns a logger with fields containing lexeme information.
func (l *Lexeme) GetLogger() *log.Entry {
	return log.WithFields(log.Fields{"lex_type": l.Type, "lex_line": l.Line, "lex_column": l.Column})
}

// errorPos  creates a new error with the given message and the position of the lexeme.
func (l *Lexeme) errorPos(msg string) error {
	return fmt.Errorf("%s [%s]", msg, l.Position())
}

// errorWrap wraps an existing error with the position of the lexeme.
func (l *Lexeme) errorWrap(err error) error {
	return fmt.Errorf("%w [%s]", err, l.Position())
}

// Position returns a string representation of the position of the lexeme.
func (l *Lexeme) Position() string {
	var s string
	if l.Line != 0 {
		s += fmt.Sprintf("%d", l.Line)
	}
	if l.Column != 0 {
		s += fmt.Sprintf(":%d", l.Column)
	}
	if s == "" {
		s = "-"
	}
	return s
}

// Lexemes is a slice of pointers to Lexeme.
type Lexemes []*Lexeme

// nameList returns a list of names of lexemes of a given token type at a given level.
func (lexemes Lexemes) nameList(tok Token) []string {
	names := make([]string, 0)
	var lvl int
	for i, l := range lexemes {
		switch l.Type {
		default:
		case LBRACE:
			lvl++
		case RBRACE:
			lvl--
		case tok:
			if lvl == 0 && i+1 < len(lexemes) && lexemes[i+1].Type == IDENTIFIER {
				names = append(names, lexemes[i+1].ToString())
			}
		}
	}
	return names
}

// contextLexer is a lexer that maintains context while lexing.
type contextLexer struct {
	input                        []rune
	position                     int
	action                       *action
	skip                         bool
	startPos, endPos, offsetLine int
	line, column                 int
	lexemes                      Lexemes
	ifBuf                        []struct {
		count, pair int
		stop        bool
	}
}

// newContextLexer creates a new contextLexer with the given input.
func newContextLexer(input []rune) *contextLexer {
	return &contextLexer{
		input:    input,
		position: 0, line: 1, column: 0,
		ifBuf: make([]struct {
			count, pair int
			stop        bool
		}, 0),
		lexemes: make(Lexemes, 0),
		action:  &action{},
	}
}

// getNextAction updates the action based on the current character and state.
func (c *contextLexer) getNextAction() {
	if c.position > len(c.input) {
		c.action.state = stateError
		return
	}
	var val int
	if c.position == len(c.input) {
		val = lexTable[c.action.state][charToAlpha(' ')]
	} else {
		val = lexTable[c.action.state][charToAlpha(c.input[c.position])]
	}
	c.action.state = (val >> 16) & 0xffff
	c.action.token = Token(val>>8) & 0xff
	c.action.flag = val & 0xff
}

// NewLexer creates a new Lexemes the given input.
func NewLexer(input []rune) (Lexemes, error) {
	c := newContextLexer(input)
	for {
		if c.position > len(c.input) {
			break
		}
		c.getNextAction()
		a := c.action
		if a.state == stateError {
			return nil, fmt.Errorf("unknown lexeme '%s' [%d:%d]",
				string(c.input[c.position:c.position+1]), c.line, c.position-c.offsetLine+1)
		}
		if hasSkip(a.flag) {
			c.position++
			c.skip = true
			continue
		}
		if a.token > UNKNOWN {
			c.ifBufCheck(a.token)
			startPos := c.position
			if hasPop(a.flag) {
				startPos = c.startPos
			}
			endPos := c.position
			if hasNext(a.flag) {
				endPos++
			}
			lexeme, err := c.getLexeme(startPos, endPos)
			if err != nil {
				return nil, err
			}
			if a.token != COMMENT {
				c.lexemes = append(c.lexemes, lexeme)
			}
		}
		if hasPush(c.action.flag) {
			c.startPos = c.position
		}

		if hasNext(c.action.flag) {
			c.position++
		}
	}
	return c.lexemes, nil
}

func (c *contextLexer) getLexeme(startPos, endPos int) (*Lexeme, error) {
	var value isLexemeValue

	tk := c.action.token
	switch tk {
	default:
	case NEWLINE:
		if c.input[startPos] == rune(0x0a) {
			c.line++
			c.offsetLine = c.position
		}
		value = NewLexemeString(string(c.input[startPos]))
	case DELIMITER:
		ch := c.input[startPos]
		tk = delimiter2Token[string(ch)]
		value = NewLexemeString(string(ch))
		if len(c.ifBuf) > 0 {
			if tk == LBRACE {
				c.ifBuf[len(c.ifBuf)-1].pair++
			}
			if tk == RBRACE {
				c.ifBuf[len(c.ifBuf)-1].pair--
				if c.ifBuf[len(c.ifBuf)-1].pair == 0 {
					c.ifBuf[len(c.ifBuf)-1].stop = true
				}
			}
		}
	case LITERAL, COMMENT:
		val := string(c.input[startPos+1 : endPos-1])
		if tk == LITERAL && c.skip {
			c.skip = false
			if c.input[startPos] == '"' && c.input[endPos-1] == '"' {
				val = strings.ReplaceAll(val, `\"`, `"`)
				val = strings.ReplaceAll(val, `\t`, "\t")
				val = strings.ReplaceAll(val, `\r`, "\r")
			}
		}
		for _, ch := range val {
			if ch == 0x0a {
				c.line++
			}
		}
		if c.input[startPos] == '"' && c.input[endPos-1] == '"' {
			val = strings.ReplaceAll(val, `\n`, "\n")
		}
		value = NewLexemeString(val)
	case OPERATOR:
		ch := string(c.input[startPos:endPos])
		val, ok := op2Token[ch]
		if !ok {
			return nil, fmt.Errorf("unknown operator '%s' [%d:%d]", ch, c.line, c.startPos-c.offsetLine+1)
		}
		value = NewLexemeToken(val)
	case NUMBER:
		name := string(c.input[startPos:endPos])
		val, err := string2Number(name)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s [%d:%d]", err, c.line, c.startPos-c.offsetLine+1)
		}
		value = NewLexemeNumber(val)
	case IDENTIFIER:
		name := string(c.input[startPos:endPos])
		if name[0] == '$' {
			tk = EXTEND
			value = NewLexemeString(name[1:])
			if err := canIdent(name[1:]); err != nil {
				return nil, err
			}
		} else if keyword, ok := KeywordValue[name]; ok {
			switch keyword {
			case ELIF:
				if len(c.ifBuf) == 0 {
					return nil, fmt.Errorf("expected statement, found '%s' [%d:%d]", name, c.line, startPos-c.offsetLine+1)
				}
				c.lexemes = append(c.lexemes,
					NewLexeme(ELSE, NewLexemeString(ELSE.ToString()), c.line, c.startPos-c.offsetLine+1),
					NewLexeme(LBRACE, NewLexemeString(LBRACE.ToString()), c.line, c.startPos-c.offsetLine+1))
				tk, value = IF, NewLexemeString(IF.ToString())
				c.ifBuf[len(c.ifBuf)-1].count++
			case ACTION, CONDITIONS:
				if len(c.lexemes) == 0 {
					return nil, fmt.Errorf("'%s' can't be the first statement [%d:%d]", name, c.line, startPos-c.offsetLine+1)
				}
				lexf := c.lexemes[len(c.lexemes)-1]
				if lexf.Type&0xff != KEYWORD || lexf.ToString() != FUNC.ToString() {
					c.lexemes = append(c.lexemes, NewLexeme(FUNC, NewLexemeString(FUNC.ToString()), c.line, startPos-c.offsetLine+1))
				}
				value = NewLexemeString(name)
			case TRUE:
				tk, value = NUMBER, NewLexemeBoolean(true)
			case FALSE:
				tk, value = NUMBER, NewLexemeBoolean(false)
			case NIL:
				tk, value = NUMBER, NewLexemeNil()
			default:
				if keyword == IF {
					c.ifBuf = append(c.ifBuf, struct {
						count, pair int
						stop        bool
					}{})
				}
				tk, value = keyword, NewLexemeString(keyword.ToString())
			}
		} else if tInfo, ok := TypeNameValue[name]; ok {
			tk, value = TYPENAME, NewLexemeToken(tInfo)
		} else {
			value = NewLexemeString(name)
		}
	}
	return NewLexeme(tk, value, c.line, startPos-c.offsetLine+1), nil
}

func (c *contextLexer) ifBufCheck(token Token) {
	if len(c.ifBuf) > 0 && c.ifBuf[len(c.ifBuf)-1].stop && token != NEWLINE {
		name := string(c.input[c.startPos:c.position])
		if name != ELSE.ToString() && name != ELIF.ToString() {
			for i := 0; i < c.ifBuf[len(c.ifBuf)-1].count; i++ {
				c.lexemes = append(c.lexemes, NewLexeme(RBRACE, NewLexemeString(RBRACE.ToString()), c.line-1, c.startPos-c.offsetLine+1))
			}
			c.ifBuf = c.ifBuf[:len(c.ifBuf)-1]
		} else {
			c.ifBuf[len(c.ifBuf)-1].stop = false
		}
	}
}

func hasNext(flag int) bool {
	return (flag & flagNext) != 0
}

func hasPush(flag int) bool {
	return (flag & flagPush) != 0
}

func hasPop(flag int) bool {
	return (flag & flagPop) != 0
}

func hasSkip(flag int) bool {
	return (flag & flagSkip) != 0
}

var (
	hexRegex     = regexp.MustCompile("^0[xX][0-9a-fA-F]+$")
	octalRegex   = regexp.MustCompile("^0[oO][0-7]+$")
	binaryRegex  = regexp.MustCompile("^0[bB][01]+$")
	decimalRegex = regexp.MustCompile("^[0-9]+$")
	floatRegex   = regexp.MustCompile("^([0-9]+([.][0-9]+)?|[.][0-9]+)([eE][+-]?[0-9]+)?$")
)

func isHex(s string) bool     { return hexRegex.MatchString(s) }
func isOctal(s string) bool   { return octalRegex.MatchString(s) }
func isBinary(s string) bool  { return binaryRegex.MatchString(s) }
func isDecimal(s string) bool { return decimalRegex.MatchString(s) }
func isFloat(s string) bool   { return floatRegex.MatchString(s) }

// string2Number convert number string literal to int64 or float64, include '._+-bBoOxXeE', support decimal_lit | binary_lit | octal_lit | hex_lit | float_lit.
func string2Number(s string) (interface{}, error) {
	for i, ch := range s {
		if ch == '_' {
			if i == 0 || s[i-1] == '_' || i == len(s)-1 {
				return nil, fmt.Errorf("underscore '%s' ", s)
			}
		}
	}
	s = strings.ReplaceAll(s, "_", "")
	if isHex(s) {
		if strings.HasPrefix(s, "0x") ||
			strings.HasPrefix(s, "0X") {
			s = s[2:]
		}
		return strconv.ParseInt(s, 16, 64)
	}

	if isOctal(s) {
		if strings.HasPrefix(s, "0o") ||
			strings.HasPrefix(s, "0O") {
			s = s[2:]
		}
		return strconv.ParseInt(s, 8, 64)
	}

	if isBinary(s) {
		if strings.HasPrefix(s, "0b") ||
			strings.HasPrefix(s, "0B") {
			s = s[2:]
		}
		return strconv.ParseInt(s, 2, 64)
	}

	if isDecimal(s) {
		return strconv.ParseInt(s, 10, 64)
	}

	if isFloat(s) {
		return strconv.ParseFloat(s, 64)
	}

	return nil, fmt.Errorf("unsupported to convert '%s' to number", s)
}
