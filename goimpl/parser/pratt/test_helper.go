package pratt

import (
	"fmt"
	"goimpl/token"
)

// stubLexer is a stub lexer for unit testing, conforms to the lexer.LexerType interface.
type stubLexer struct {
	tokens []token.Token
}

func (l *stubLexer) NextToken() token.Token {
	if len(l.tokens) == 0 {
		return token.Token{Type: token.EOF, Literal: ""}
	}
	tok := l.tokens[0]
	l.tokens = l.tokens[1:]
	return tok
}

func printErrs(p *PrattParser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	fmt.Printf("parser has %d error(s)\n", len(errors))
	for _, msg := range errors {
		fmt.Printf("parser error: %q\n", msg)
	}
}
