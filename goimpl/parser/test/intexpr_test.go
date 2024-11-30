package test

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

func TestIntLiteralExpr(t *testing.T) {
	var lxr = lexer.NewStubLexer([]token.Token{
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	var expectedAst = &ast.ExpressionStatement{
		Tok:     token.Token{Type: token.INT, Literal: "5"},
		Exprssn: &ast.IntegerLiteral{Tok: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
	}

	for pname, parser := range testParsers(&lxr).initPratt().initCombinator(combinator.IntExpr).parsers {
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
				t.Fatalf("expected: \n %s \n got: \n %s \n", expectedAst.Dump(1), stmt.Dump(1))
			}
		})
	}
}
