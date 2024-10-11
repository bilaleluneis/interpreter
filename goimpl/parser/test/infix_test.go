package test

import (
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

var infixTests = []struct {
	expectedLeft  string
	expectedOp    string
	expectedRight string
	lexr          lexer.StubLexer
}{
	// 5 + 5;
	{"5", "+", "5", lexer.NewStubLexer([]token.Token{
		{Type: token.INT, Literal: "5"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})},
}

func TestInfixExpression(t *testing.T) {
	for _, tt := range infixTests {
		for pname, parser := range testParsers(&tt.lexr).initPratt().initCombinator(combinator.InfixInt).parsers {
			if program, ok := parser.ParseProgram(); !ok {
				fail(pname, t, "nexpected ok program got !ok")
			} else if expr := toExpression(program.Statements[0]); expr == nil {
				fail(pname, t, "\nexpected *ast.ExpressionStatement, got: %T", program.Statements[0])
			} else if infix := toInfixExpression(expr); infix == nil {
				fail(pname, t, "\nexpected *ast.InfixExpression")
			} else if infix.Operator != tt.expectedOp {
				fail(pname, t, "\nexpected operator %s, got: %s", tt.expectedOp, infix.Operator)
			} else if leftExpr := toIntegerLiteral(infix.Left); leftExpr == nil {
				fail(pname, t, "\nexpected *ast.IntegerLiteral")
			} else if leftExpr.TokenLiteral() != tt.expectedLeft {
				fail(pname, t, "\nexpected token literal %s, got: %s", tt.expectedLeft, leftExpr.TokenLiteral())
			} else if right := toIntegerLiteral(infix.Right); right == nil {
				fail(pname, t, "\nexpected *ast.IntegerLiteral")
			} else if right.TokenLiteral() != tt.expectedRight {
				fail(pname, t, "\nexpected token literal %s, got: %s", tt.expectedRight, right.TokenLiteral())
			} else {
				success(pname, t, "parser %s passed with result %s", pname, program)
			}
		}
	}
}
