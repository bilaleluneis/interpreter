package test

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

// FIXME: combinator parser not used in this test
func TestBooleanExpression(t *testing.T) {
	var lxr = lexer.NewStubLexer([]token.Token{
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	var expectedAst = &ast.ExpressionStatement{
		Tok:     token.Token{Type: token.TRUE, Literal: "true"},
		Exprssn: &ast.Boolean{Tok: token.Token{Type: token.TRUE, Literal: "true"}, Value: true},
	}

	for pname, parser := range testParsers(&lxr).initPratt().parsers {
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
