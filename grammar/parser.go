package parser

import (
	"fmt"
	"os"

	needle "github.com/IBAX-io/needle/grammar/dist-go"

	"github.com/antlr4-go/antlr/v4"
)

// ParserBase is a wrapper for the generated parser that provides some extra functionality.
type ParserBase struct {
	*needle.NeedleParser
	input      []byte
	path       string
	inputSteam *antlr.InputStream
	lexer      *needle.NeedleLexer
	tokenSteam *antlr.CommonTokenStream
	listener   *SyntaxErrorListener
	listeners  map[string]antlr.ParseTreeListener
}

// NewParserBase returns a new parser for the given data and path.
func NewParserBase(data []byte, path string) (*ParserBase, error) {
	input := antlr.NewInputStream(string(data))
	lexer := needle.NewNeedleLexer(input)
	listener := NewSyntaxErrorListener(path)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := needle.NewNeedleParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(listener)
	return &ParserBase{
		NeedleParser: p,
		input:        data,
		path:         path,
		inputSteam:   input,
		lexer:        lexer,
		tokenSteam:   stream,
		listener:     listener,
		listeners:    make(map[string]antlr.ParseTreeListener),
	}, nil
}

// NewParserFile returns a new parser for the given file.
func NewParserFile(file string) (*ParserBase, error) {
	readFile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return NewParserBase(readFile, file)
}

// Parse parses the SourceMain and walks the parse tree.
func (p *ParserBase) Parse() {
	tree := p.SourceMain()
	for _, listener := range p.listeners {
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	}
}

func (p *ParserBase) RegisterListener(name string, listener antlr.ParseTreeListener) {
	if _, ok := p.listeners[name]; ok {
		fmt.Println("listener", name, "already registered")
		return
	}
	p.listeners[name] = listener
}

// PrintlnError prints all syntax errors to stdout.
func (p *ParserBase) PrintlnError() {
	p.listener.PrintlnError()
}

// SyntaxErrorListener is a custom error listener that collects syntax errors.
type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	*needle.BaseNeedleParserListener
	Errors []SyntaxError
	Path   string
}

// SyntaxError represents a syntax error.
type SyntaxError struct {
	Line    int
	Column  int
	Message string
}

// NewSyntaxErrorListener returns a new SyntaxErrorListener.
func NewSyntaxErrorListener(path string) *SyntaxErrorListener {
	return &SyntaxErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]SyntaxError, 0),
		Path:                 path,
	}
}

// PrintlnError prints all syntax errors to stdout.
func (s *SyntaxErrorListener) PrintlnError() {
	for _, e := range s.Errors {
		str := "syntax error "
		if s.Path != "" {
			str += s.Path + ":"
		}
		str += fmt.Sprintf("%d:%d %s\n", e.Line, e.Column, e.Message)
		fmt.Printf(str)
	}
}

// SyntaxError is called when a syntax error is encountered.
func (s *SyntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	s.Errors = append(s.Errors, SyntaxError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}
