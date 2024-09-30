package test

import (
	"goimpl/ast"
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
		var program ast.Program
		var literal string
		var ok bool
		if program, ok = parser.ParseProgram(); !ok {
			t.Errorf("\nexpected ok program for parser %s, got: !ok", pname)
		} else if !isExpressionStmt(program.Statements[0]) {
			t.Errorf("\nexpected *ast.ExpressionStatement, got: %T", program.Statements[0])
		} else if literal, ok = isIntLiteral(program.Statements[0]); !ok {
			t.Errorf("\nexpected *ast.IntegerLiteral")
		} else if literal != "5" {
			t.Errorf("\nexpected token literal 5, got: %s", literal)
		}
	}
}
