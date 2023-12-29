package compile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Lexeme struct {
	Type Token
	//string
	//int64
	//float64
	//Token
	//bool
	//nil
	Value  any
	Line   int
	Column int
}

func NewLexeme(Type Token, value any, line int, column int) *Lexeme {
	return &Lexeme{Type: Type, Value: value, Line: line, Column: column}
}

// GetLogger returns logger
func (l *Lexeme) GetLogger() *log.Entry {
	return log.WithFields(log.Fields{"lex_type": l.Type, "lex_line": l.Line, "lex_column": l.Column})
}

func (l *Lexeme) error(msg string) error {
	return fmt.Errorf("%s [%s]", msg, l.Position())
}

func (l *Lexeme) errorWrap(err error)error {
	return fmt.Errorf("%s [%s]", err, l.Position())
}

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

type Lexemes []*Lexeme

func (lexemes Lexemes) nameList(tok Token, level int) []string {
	names := make([]string, 0)
	var lvl int
	for i, lexeme := range lexemes {
		switch lexeme.Type {
		default:
		case LBRACE:
			lvl++
		case RBRACE:
			lvl--
		case tok:
			if lvl == level && i+1 < len(lexemes) && lexemes[i+1].Type == IDENTIFIER {
				names = append(names, lexemes[i+1].Value.(string))
			}
		}
	}
	return names
}

type contextLexer struct {
	input                     []rune
	position                  int
	action                    *action
	skip                      bool
	startPos, endPos, offline int
	line, column              int
	lexemes                   Lexemes
	ifBuf                     []struct {
		count int
		pair  int
		stop  bool
	}
}

func newContextLexer(input []rune) *contextLexer {
	return &contextLexer{input: input,
		position: 0, line: 1, column: 0,
		ifBuf: make([]struct {
			count, pair int
			stop        bool
		}, 0),
		lexemes: make(Lexemes, 0),
		action:  &action{},
	}
}

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
				string(c.input[c.startPos:c.position]), c.line, c.position-c.offline+1)
		}
		if hasSkip(a.flag) {
			c.position++
			c.skip = true
			continue
		}
		if a.token > UNKNOWN {
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
	var value any
	tk := c.action.token
	switch tk {
	default:
	case NEWLINE:
		if c.input[startPos] == rune(0x0a) {
			c.line++
			c.offline = c.position
		}
		value = string(c.input[startPos])
	case DELIMITER:
		ch := c.input[startPos]
		tk = delimiter2Token[string(ch)]
		value = string(ch)
		if len(c.ifBuf) > 0 {
			if ch == '{' {
				c.ifBuf[len(c.ifBuf)-1].pair++
			}
			if ch == '}' {
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
				//offline = off + uint32(i) + 1
			}
		}
		if c.input[startPos] == '"' && c.input[endPos-1] == '"' {
			val = strings.ReplaceAll(val, `\n`, "\n")
		}
		value = val
	case OPERATOR:
		val := string(c.input[startPos:endPos])
		var ok bool
		value, ok = op2Token[val]
		if !ok {
			return nil, fmt.Errorf("unknown operator '%s' [%d:%d]", val, c.line, c.position-c.offline+1)
		}
	case NUMBER:
		name := string(c.input[startPos:endPos])
		val, err := string2Number(name)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s [%d:%d]", err, c.line, c.position-c.offline+1)
		}
		value = val
	case IDENTIFIER:
		name := string(c.input[startPos:endPos])
		if name[0] == '$' {
			tk = EXTEND
			value = name[1:]
			if err := canIdent(name[1:]); err != nil {
				return nil, err
			}
			break
		} else if keyID, ok := KeywordValue[name]; ok {
			switch keyID {
			case ELIF:
				if len(c.ifBuf) == 0 {
					return nil, fmt.Errorf(`expected statement, found '%s' [%d:%d]`, name, c.line, startPos-c.offline+1)
				}
				c.lexemes = append(c.lexemes,
					NewLexeme(ELSE, ELSE.String(), c.line, c.position-c.offline+1),
					NewLexeme(LBRACE, LBRACE.String(), c.line, c.position-c.offline+1))
				tk, value = IF, IF.String()
				c.ifBuf[len(c.ifBuf)-1].count++
			case ACTION, CONDITIONS:
				if len(c.lexemes) == 0 {
					return nil, fmt.Errorf(`'%s' can't be the first statement [%d:%d]`, name, c.line, startPos-c.offline+1)
				}
				lexf := c.lexemes[len(c.lexemes)-1]
				if lexf.Type&0xff != KEYWORD || lexf.Value.(string) != FUNC.String() {
					c.lexemes = append(c.lexemes, NewLexeme(FUNC, FUNC.String(), c.line, startPos-c.offline+1))
				}
				value = name

			case TRUE:
				tk, value = NUMBER, true
			case FALSE:
				tk, value = NUMBER, false
			case NIL:
				tk, value = NUMBER, nil
			default:
				if keyID == IF {
					c.ifBuf = append(c.ifBuf, struct {
						count, pair int
						stop        bool
					}{})
				}
				tk, value = keyID, keyID.String()
			}
		} else if tInfo, ok := TypeNameValue[name]; ok {
			tk, value = TYPENAME, tInfo
		} else {
			value = name
		}
	}
	return NewLexeme(tk, value, c.line, startPos-c.offline+1), nil
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
	s = strings.ReplaceAll(s, `_`, ``)
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
