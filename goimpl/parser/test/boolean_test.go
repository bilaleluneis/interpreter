package test

import (
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

var booleanTests = []struct {
	input    string
	expected bool
	lexer    lexer.StubLexer
}{
	{"true;", true, lexer.NewStubLexer([]token.Token{
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})},
}

// FIXME: combinator parser not used in this test
func TestBooleanExpression(t *testing.T) {
	for _, tt := range booleanTests {
		for pname, parser := range testParsers(&tt.lexer).initPratt().parsers {
			if program, ok := parser.ParseProgram(); !ok {
				fail(pname, t, "expected ok program got !ok")
			} else if expr := toExpression(program.Statements[0]); expr == nil {
				fail(pname, t, "expected expression")
			} else if boolean := toBoolean(expr); boolean == nil {
				fail(pname, t, "expected boolean expression")
			} else {
				success(pname, t, "parser %s passed with result %s", pname, program)
			}
		}
	}
}
