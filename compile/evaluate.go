package compile

import "fmt"

type scanner struct {
	lex *Lexeme
	i   int
	ns  bool
}

func (p *Parser) set(i int) {
	p.i = i
	p.lex = p.inputs[i]
}

// nextNonBlank returns the next non-blank lexeme.
func (p *Parser) nextNonBlankN(n int) *Lexeme {
	count := 0
	for _, lex := range p.inputs[p.i+1:] {
		if lex.Type == NEWLINE && lex.Value == "\n" {
			continue
		}
		count++
		if count == n {
			return lex
		}
	}
	return nil
}

// prevNonBlankN returns the previous non-blank lexeme.
func (p *Parser) prevNonBlankN(n int) *Lexeme {
	count := 0
	for i := p.i - 1; i >= 0; i-- {
		lex := p.inputs[i]
		if lex.Type == NEWLINE && lex.Value == "\n" {
			continue
		}
		count--
		if count == -n {
			return lex
		}
	}
	return nil
}

func (p *Parser) syntaxError(msg string) error {
	return p.lex.error(msg)
}

func (p *Parser) syntaxErrorWrap(err error) error {
	return p.lex.errorWrap(err)
}

func (p *Parser) syntaxErrorExpected(msg string) error {
	return p.lex.error(fmt.Sprintf("unexpected %s, %s", p.lex.Value, msg))
}
