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
				t.Errorf("\nexpected ok program for parser %s, got: !ok", pname)
			} else if !isExpressionStmt(program.Statements[0]) {
				t.Errorf("\nexpected *ast.ExpressionStatement, got: %T", program.Statements[0])
			} else if prefix, ok := isPrefixExpression(program.Statements[0]); !ok {
				t.Errorf("\nexpected *ast.PrefixExpression")
			} else if prefix.Operator != tt.expectedOperator {
				t.Errorf("\nexpected operator %s, got: %s", tt.expectedOperator, prefix.Operator)
			} //else if literal, ok := isIntLiteral(prefix.Right); !ok {
			// 	t.Errorf("\nexpected *ast.IntegerLiteral")
			// } else if literal != tt.expectedValue {
			// 	t.Errorf("\nexpected token literal %d, got: %d", tt.expectedValue, literal)
			// }
		}
	}
}
