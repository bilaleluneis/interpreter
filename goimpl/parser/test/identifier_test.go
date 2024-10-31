package test

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	var lxr = lexer.NewStubLexer([]token.Token{
		{Type: token.IDENTIFIER, Literal: "foobar"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	var expectedAst = &ast.ExpressionStatement{
		Tok:     token.Token{Type: token.IDENTIFIER, Literal: "foobar"},
		Exprssn: &ast.Identifier{Tok: token.Token{Type: token.IDENTIFIER, Literal: "foobar"}, Value: "foobar"},
	}

	for pname, parser := range testParsers(&lxr).initPratt().initCombinator(combinator.IdentifierExpr).parsers {
		t.Run(pname, func(t *testing.T) {
			t.Logf("testing parser %s", pname)
			program, ok := parser.ParseProgram()
			if !ok {
				t.Fatalf("parser returned error")
			}
			if len(program.Statements) != 1 {
				t.Fatalf("expected 1 statement, got %d", len(program.Statements))
			}
			stmt := program.Statements[0]
			if stmt.Dump(1) != expectedAst.Dump(1) {
				t.Fatalf("expected %s, got %s", expectedAst.Dump(1), stmt.Dump(1))
			}
		})
	}
}
