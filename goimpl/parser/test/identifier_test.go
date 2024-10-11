package test

import (
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	tokens := []token.Token{
		{Type: token.IDENTIFIER, Literal: "foobar"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	l := lexer.NewStubLexer(tokens)

	for pname, parser := range testParsers(&l).initPratt().initCombinator(combinator.IdentifierExpr).parsers {
		if program, ok := parser.ParseProgram(); !ok {
			fail(pname, t, "got not ok program")
		} else if expr := toExpression(program.Statements[0]); expr == nil {
			fail(pname, t, "expected expression")
		} else if literal := toIdentifier(expr); literal == nil {
			fail(pname, t, "expected identifier")
		} else if literal.TokenLiteral() != "foobar" {
			fail(pname, t, "expected token literal foobar, got: %s", literal.TokenLiteral())
		} else {
			success(pname, t, "parser %s passed with result %s", pname, program)
		}
	}
}
