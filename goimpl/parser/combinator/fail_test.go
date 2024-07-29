package combinator

import (
	"goimpl/lexer"
	"goimpl/parser"
	"goimpl/token"
	"testing"
)

func TestFail(t *testing.T) {
	var l lexer.LexerType = &parser.StubLexer{
		Toks: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
	}

	parse(Fail, l)
}
