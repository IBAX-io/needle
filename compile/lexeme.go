package compile

import (
	"encoding/binary"
	"fmt"
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

type ifBuf struct {
	count int
	pair  int
	stop  bool
}

type Lexemes []*Lexeme

func NewLexer(input []rune) (Lexemes, error) {
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
		if off == length-1 {
			todo(' ')
		} else {
			todo(input[off])
		}

		if curState == Error {
			return nil, fmt.Errorf(`unknown lexeme '%s' [%d:%d]`,
				string(input[off:off+1]), line, off-offline+1)
		}
		if hasSkip(flag) {
			off++
			skip = true
			continue
		}
		if tk > 0 {
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
						lexemes = append(lexemes, NewLexeme(RBRACE, uint32(keyRBRACE), line, lexOffset-offline+1))
					}
					ifbuf = ifbuf[:len(ifbuf)-1]
				} else {
					ifbuf[len(ifbuf)-1].stop = false
				}
			}
			var value any
			switch tk {
			case NEWLINE:
				if input[lexOffset] == rune(0x0a) {
					line++
					offline = off
				}
			case SYSTEM:
				ch := Token(input[lexOffset])
				tk |= ch << 8
				value = ch
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
					val = strings.Replace(strings.Replace(val, `\"`, `"`, -1), `\t`, "\t", -1)
					val = strings.Replace(val, `\r`, "\r", -1)
					//val = strings.Replace(strings.Replace(val, `\r`, "\r", -1), `\n`, "\n", -1)
				}
				//value = val
				for _, ch := range val {
					if ch == 0x0a {
						line++
						//offline = off + uint32(i) + 1
					}
				}
				val = strings.Replace(val, `\n`, "\n", -1)
				value = val
			case OPERATOR:
				oper := []byte(string(input[lexOffset:right]))
				value = Token(binary.BigEndian.Uint32(append(make([]byte, 4-len(oper)), oper...)))
			case NUMBER:
				name := string(input[lexOffset:right])
				if strings.ContainsAny(name, `.`) {
					val, err := strconv.ParseFloat(name, 64)
					if err != nil {
						return nil, fmt.Errorf(`%v %s [%d:%d]`, err, name, line, off-offline+1)
					}
					value = val
				} else if val, err := strconv.ParseInt(name, 10, 64); err == nil {
					value = val
				} else {
					return nil, fmt.Errorf(`%v %s [%d:%d]`, err, name, line, off-offline+1)
				}
			case IDENTIFIER:
				name := string(input[lexOffset:right])
				if name[0] == '$' {
					tk = EXTEND
					value = name[1:]
				} else if keyID, ok := KeywordValue[name]; ok {
					switch keyID {
					case ELIF:
						if len(ifbuf) > 0 {
							lexemes = append(lexemes,
								NewLexeme(ELSE, uint32(keyElse), line, lexOffset-offline+1),
								NewLexeme(LBRACE, uint32(keyLBRACE), line, lexOffset-offline+1))
							tk = IF
							value = uint32(keyIf)
							ifbuf[len(ifbuf)-1].count++
						}
					case ACTION, CONDITIONS:
						if len(lexemes) > 0 {
							lexf := lexemes[len(lexemes)-1]
							if lexf.Type&0xff != KEYWORD || lexf.Value.(Token) != keyFunc {
								lexemes = append(lexemes, NewLexeme(FUNC, uint32(keyFunc), line, lexOffset-offline+1))
							}
						}
						value = name
					case TRUE:
						tk = NUMBER
						value = true
					case FALSE:
						tk = NUMBER
						value = false
					case NIL:
						tk = NUMBER
						value = nil
					default:
						if keyID == IF {
							ifbuf = append(ifbuf, ifBuf{})
						}
						tk = keyID
						value = keyID
					}
				} else if tInfo, ok := TypeNameReflect[TypeNameValue[name]]; ok {
					tk = TYPENAME
					value = tInfo
				} else {
					value = name
				}
			default:
				//fmt.Println("error---", off, tk, start, off, string(input[start:off]))
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
