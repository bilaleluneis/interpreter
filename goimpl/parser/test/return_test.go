package test

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	// return 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.RETURN, Literal: "return"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	expectedAst := &ast.Return{
		Tok: token.Token{Type: token.RETURN, Literal: "return"},
		Value: &ast.IntegerLiteral{
			Tok:   token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
	}

	for pname, parser := range testParsers(&l).initPratt().initCombinator(combinator.Retrn).parsers {
		t.Run(pname, func(t *testing.T) {
			t.Logf("testing parser %s", pname)
			program, ok := parser.ParseProgram()
			if !ok {
				t.Fatalf("expected ok program got !ok")
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
