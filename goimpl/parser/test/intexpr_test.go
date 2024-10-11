package test

import (
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

func TestIntLiteralExpr(t *testing.T) {
	// 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	for pname, parser := range testParsers(&l).initPratt().initCombinator(combinator.IntExpr).parsers {
		if program, ok := parser.ParseProgram(); !ok {
			fail(pname, t, "got not ok program")
		} else if expr := toExpression(program.Statements[0]); expr == nil {
			fail(pname, t, "expected expression")
		} else if literal := toIntegerLiteral(expr); literal == nil {
			fail(pname, t, "expected integer literal")
		} else if literal.TokenLiteral() != "5" {
			fail(pname, t, "expected token literal 5, got: %s", literal.TokenLiteral())
		} else {
			success(pname, t, "parser %s passed with result %s", pname, program)
		}
	}
}
