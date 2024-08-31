package lexer

import "goimpl/token"

// StubLexer is a stub lexer for unit testing, conforms to the lexer.LexerType interface.
type StubLexer []token.Token

func (l *StubLexer) NextToken() token.Token {
	if len(*l) == 0 {
		return token.Token{Type: token.EOF, Literal: ""}
	}
	tok := (*l)[0]
	*l = (*l)[1:]
	return tok
}

func (l StubLexer) GetCopy() StubLexer {
	copied := NewStubLexer(l)
	return copied
}

func NewStubLexer(toks []token.Token) StubLexer {
	l := StubLexer{}
	for _, tok := range toks {
		l = append(l, tok)
	}
	return l
}
