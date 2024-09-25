package test

import (
	"goimpl/ast"
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

	for pname, parser := range testParsers(l).initPratt().initCombinator(combinator.IdentifierExpr).parsers {
		var program ast.Program
		var literal string
		var ok bool
		if program, ok = parser.ParseProgram(); !ok {
			t.Errorf("\nexpected ok program for parser %s, got: !ok", pname)
		} else if !isExpressionStmt(program.Statements[0]) {
			t.Errorf("\nexpected *ast.ExpressionStatement, got: %T", program.Statements[0])
		} else if literal, ok = isIdentifier(program.Statements[0]); !ok {
			t.Errorf("\nexpected identifier expression")
		} else if literal != "foobar" {
			t.Errorf("\nexpected identifier 'foobar', got: %s", literal)
		}
	}
}
