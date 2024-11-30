package test

import (
	"goimpl/ast"
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
	expectedAst   *ast.InfixExpression
}{
	// 5 + 5;
	{"5", "+", "5", lexer.NewStubLexer([]token.Token{
		{Type: token.INT, Literal: "5"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}}),
		&ast.InfixExpression{
			Left:     &ast.IntegerLiteral{Tok: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
			Operator: "+",
			Right:    &ast.IntegerLiteral{Tok: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		},
	},
}

func TestInfixExpression(t *testing.T) {
	for _, tt := range infixTests {
		for pname, parser := range testParsers(&tt.lexr).initPratt().initCombinator(combinator.InfixInt).parsers {
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
