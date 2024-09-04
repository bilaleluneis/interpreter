package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

func TestLet(t *testing.T) {
	//let x = 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	pr := NewParseResult(l)
	parseResult := pr.Parse(Let[lexer.StubLexer])
	filterResult := Filtr(parseResult, func(s ast.Statement) bool { return s.TokenLiteral() == "let" })
	if len(filterResult) != 1 {
		t.Fatal("filter should return one result got", len(filterResult))
	}

}
