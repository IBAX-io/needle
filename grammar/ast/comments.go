package ast

import (
	"strings"

	needle "github.com/IBAX-io/needle/grammar/dist-go"
	"github.com/antlr4-go/antlr/v4"
)

type Comment struct {
	Src      SrcPos
	NodeType string
	Text     string
}

func (b *Builder) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if !b.commentsParsed {
		tokens := b.parser.GetTokenStream().(*antlr.CommonTokenStream).GetAllTokens()
		for _, token := range tokens {
			if token.GetTokenType() == needle.NeedleLexerLINE_COMMENT {
				comment := &Comment{
					Src: SrcPos{
						Line:   token.GetLine(),
						Column: token.GetColumn(),
						Start:  token.GetStart(),
						End:    token.GetStop(),
						Length: token.GetStop() - token.GetStart() + 1,
					},
					NodeType: "LINE_COMMENT",
					Text:     strings.TrimSpace(token.GetText()),
				}
				b.Comments = append(b.Comments, comment)
			}
			if token.GetTokenType() == needle.NeedleLexerCOMMENT {
				comment := &Comment{
					Src: SrcPos{
						Line:   token.GetLine(),
						Column: token.GetColumn(),
						Start:  token.GetStart(),
						End:    token.GetStop(),
						Length: token.GetStop() - token.GetStart() + 1,
					},
					NodeType: "MULTILINE_COMMENT",
					Text:     strings.TrimSpace(token.GetText()),
				}
				b.Comments = append(b.Comments, comment)
			}
		}
		b.commentsParsed = true
	}
}
