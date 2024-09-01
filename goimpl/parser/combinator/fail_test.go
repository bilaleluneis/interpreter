package combinator

import (
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

func TestFail(t *testing.T) {
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})
	// TODO:
	// - run with multiple combinators and check that you get always diff lexer instance
	// - check that orginal lexer is not mutated
	// - use iterator to go over combinators
	pr := NewParseResult(l)
	result := pr.Parse(Fail[lexer.StubLexer])
	if len(result) == 0 {
		t.Fatalf("expected error")
	}
}
