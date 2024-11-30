package test

import (
	"goimpl/ast"
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
	expectedAst      *ast.PrefixExpression
}{
	// -5;
	{"-", 5, lexer.NewStubLexer([]token.Token{
		{Type: token.MINUS, Literal: "-"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}}),
		&ast.PrefixExpression{
			Operator: "-",
			Right:    &ast.IntegerLiteral{Tok: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		},
	},

	// !5;
	{"!", 5, lexer.NewStubLexer([]token.Token{
		{Type: token.BANG, Literal: "!"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}}),
		&ast.PrefixExpression{
			Operator: "!",
			Right:    &ast.IntegerLiteral{Tok: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		},
	},
}

func TestPrefixExpression(t *testing.T) {
	for _, tt := range prefixTests {
		for pname, parser := range testParsers(&tt.lexr).initPratt().initCombinator(combinator.PrefixInt).parsers {
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
				exprStmt, ok := stmt.(*ast.ExpressionStatement)
				if !ok {
					t.Fatalf("expected *ast.ExpressionStatement, got %T", stmt)
				}
				if exprStmt.Exprssn.Dump(1) != tt.expectedAst.Dump(1) {
					t.Fatalf("expected %s, got %s", tt.expectedAst.Dump(1), exprStmt.Exprssn.Dump(1))
				}
			})
		}
	}
}
