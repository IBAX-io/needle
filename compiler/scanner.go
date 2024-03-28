package compiler

type scanner struct {
	lex *Lexeme
	i   int
	ns  bool // newline and semicolon
}

func (p *Parser) set(i int) {
	p.i = i
	p.lex = p.inputs[i]
}

// nextN returns the next non-blank lexeme.
func (p *Parser) nextN(n int) *Lexeme {
	count := 0
	if n == 0 {
		return p.lex
	}
	for _, lex := range p.inputs[p.i+1:] {
		if lex.Type == NEWLINE && lex.ToString() == "\n" {
			continue
		}
		count++
		if count == n {
			return lex
		}
	}
	return nil
}

// prevN returns the previous non-blank lexeme.
func (p *Parser) prevN(n int) *Lexeme {
	count := 0
	if n == 0 {
		return p.lex
	}
	for i := p.i - 1; i >= 0; i-- {
		lex := p.inputs[i]
		if lex.Type == NEWLINE && lex.ToString() == "\n" {
			continue
		}
		count--
		if count == -n {
			return lex
		}
	}
	return nil
}
