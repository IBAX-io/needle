package compile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Lexeme struct {
	Type   Token
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

func NewLexer(input []rune) (Lexemes, error) {
	type ifBuf struct {
		count int
		pair  int
		stop  bool
	}

	var (
		curState, offline,
		flag, start, off int
		tk      Token
		line    = 1
		skip    bool
		length  = len(input) + 1
		lexemes = make(Lexemes, 0, len(input)/4)
		ifbuf   = make([]ifBuf, 0)
	)
	todo := func(r rune) {
		val := lexTable[curState][charToAlpha(r)]
		curState = (val >> 16) & 0xffff
		tk = Token(val>>8) & 0xff
		flag = val & 0xff
	}
	for off < length {
		var end bool
		if off == length-1 {
			todo(' ')
			end = true
		} else {
			todo(input[off])
		}
		if curState == stateError {
			l, h := off, off+1
			if end || off != length-1 {
				l, h = start, off
			}
			return nil, fmt.Errorf("unknown lexeme '%s' [%d:%d]",
				string(input[l:h]), line, off-offline+1)
		}
		if hasSkip(flag) {
			off++
			skip = true
			continue
		}
		if tk > UNKNOWN {
			lexOffset := off
			if hasPop(flag) {
				lexOffset = start
			}
			right := off
			if hasNext(flag) {
				right++
			}
			if len(ifbuf) > 0 && ifbuf[len(ifbuf)-1].stop && tk != NEWLINE {
				name := string(input[lexOffset:right])
				if name != `else` && name != `elif` {
					for i := 0; i < ifbuf[len(ifbuf)-1].count; i++ {
						lexemes = append(lexemes, NewLexeme(RBRACE, RBRACE, line, lexOffset-offline+1))
					}
					ifbuf = ifbuf[:len(ifbuf)-1]
				} else {
					ifbuf[len(ifbuf)-1].stop = false
				}
			}
			var value any
			switch tk {
			default:
			case NEWLINE:
				if input[lexOffset] == rune(0x0a) {
					line++
					offline = off
				}
			case DELIMITER:
				ch := input[lexOffset]
				tk = delimiter2Token[ch]
				value = string(ch)
				if len(ifbuf) > 0 {
					if ch == '{' {
						ifbuf[len(ifbuf)-1].pair++
					}
					if ch == '}' {
						ifbuf[len(ifbuf)-1].pair--
						if ifbuf[len(ifbuf)-1].pair == 0 {
							ifbuf[len(ifbuf)-1].stop = true
						}
					}
				}
			case LITERAL, COMMENT:
				val := string(input[lexOffset+1 : right-1])
				if tk == LITERAL && skip {
					skip = false
					if input[lexOffset] == '"' && input[right-1] == '"' {
						val = strings.ReplaceAll(val, `\"`, `"`)
						val = strings.ReplaceAll(val, `\t`, "\t")
						val = strings.ReplaceAll(val, `\r`, "\r")
					}
				}
				for _, ch := range val {
					if ch == 0x0a {
						line++
						//offline = off + uint32(i) + 1
					}
				}
				if input[lexOffset] == '"' && input[right-1] == '"' {
					val = strings.ReplaceAll(val, `\n`, "\n")
				}
				value = val
			case OPERATOR:
				val := string(input[lexOffset:right])
				var ok bool
				value, ok = op2Token[val]
				if !ok {
					return nil, fmt.Errorf("unknown operator '%s' [%d:%d]", val, line, off-offline+1)
				}
			case NUMBER:
				name := string(input[lexOffset:right])
				val, err := string2Number(name)
				if err != nil {
					return nil, fmt.Errorf("invalid number: %s [%d:%d]", err, line, off-offline+1)
				}
				value = val
			case IDENTIFIER:
				name := string(input[lexOffset:right])
				if name[0] == '$' {
					tk = EXTEND
					value = name[1:]
					if err := canIdent(name[1:]); err != nil {
						return nil, err
					}
				} else if keyID, ok := KeywordValue[name]; ok {
					switch keyID {
					case ELIF:
						if len(ifbuf) == 0 {
							return nil, fmt.Errorf(`expected statement, found '%s' [%d:%d]`, name, line, lexOffset-offline+1)
						}
						lexemes = append(lexemes,
							NewLexeme(ELSE, Keyword2Str(ELSE), line, lexOffset-offline+1),
							NewLexeme(LBRACE, Keyword2Str(LBRACE), line, lexOffset-offline+1))
						tk, value = IF, Keyword2Str(IF)
						ifbuf[len(ifbuf)-1].count++
					case ACTION, CONDITIONS:
						if len(lexemes) == 0 {
							return nil, fmt.Errorf(`'%s' can't be the first statement [%d:%d]`, name, line, lexOffset-offline+1)
						}
						lexf := lexemes[len(lexemes)-1]
						if lexf.Type&0xff != KEYWORD || lexf.Value.(string) != Keyword2Str(FUNC) {
							lexemes = append(lexemes, NewLexeme(FUNC, Keyword2Str(FUNC), line, lexOffset-offline+1))
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
							ifbuf = append(ifbuf, ifBuf{})
						}
						tk, value = keyID, Keyword2Str(keyID)
					}
				} else if tInfo, ok := TypeNameValue[name]; ok {
					tk, value = TYPENAME, tInfo
				} else {
					value = name
				}
			}
			if tk != COMMENT {
				lexemes = append(lexemes, NewLexeme(tk, value, line, lexOffset-offline+1))
			}
		}
		if hasPush(flag) {
			start = off
		}
		if hasNext(flag) {
			off++
		}
	}
	return lexemes, nil
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
