package parser

import (
	"goimpl/token"
)

// StubLexer is a stub lexer for unit testing, conforms to the lexer.LexerType interface.
type StubLexer struct {
	Toks []token.Token
}

func (l *StubLexer) NextToken() token.Token {
	if len(l.Toks) == 0 {
		return token.Token{Type: token.EOF, Literal: ""}
	}
	tok := l.Toks[0]
	l.Toks = l.Toks[1:]
	return tok
}
