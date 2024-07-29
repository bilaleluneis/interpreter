package combinator

import (
	"goimpl/parser"
	"goimpl/token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {
	// let x = 5;
	_ = &parser.StubLexer{
		Toks: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
	}

	//func(lexr lexer.LexerType) CombinatorParser {
	//	return func(lexerType lexer.LexerType) CombinatorParser {
	//		return nil
	//	}
	//}(l)

}
