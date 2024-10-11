package test

import (
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

// <PrefixExpression> ::= <prefix operator> <expression>
// <prefix operator> ::= "-" | "!"
// example input: -5;

var prefixTests = []struct {
	expectedOperator string
	expectedValue    int64
	lexr             lexer.StubLexer
}{
	// -5;
	{"-", 5, lexer.NewStubLexer([]token.Token{
		{Type: token.MINUS, Literal: "-"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}})},

	// !5;
	{"!", 5, lexer.StubLexer([]token.Token{
		{Type: token.BANG, Literal: "!"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}})},
}

func TestPrefixExpression(t *testing.T) {
	for _, tt := range prefixTests {
		for pname, parser := range testParsers(&tt.lexr).initPratt().initCombinator(combinator.PrefixInt).parsers {
			if program, ok := parser.ParseProgram(); !ok {
				fail(pname, t, "expected ok program got !ok")
			} else if expr := toExpression(program.Statements[0]); expr == nil {
				fail(pname, t, "expected expression got nil")
			} else if prefix := toPrefixExpression(expr); prefix == nil {
				fail(pname, t, "expected ast.PrefixExpression")
			} else if prefix.Operator != tt.expectedOperator {
				fail(pname, t, "expected operator %s, got: %s", tt.expectedOperator, prefix.Operator)
			} else if literal := toIntegerLiteral(prefix.Right); literal == nil {
				fail(pname, t, "expected right side of expression to be and integer literal")
			} else if literal.Value != tt.expectedValue {
				fail(pname, t, "expected value %d, got: %d", tt.expectedValue, literal.Value)
			} else {
				success(pname, t, "parser %s passed with result %s", pname, program)
			}
		}
	}
}
