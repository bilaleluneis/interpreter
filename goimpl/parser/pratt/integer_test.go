package pratt

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

func TestIntegerLiteralExpr(t *testing.T) {
	// 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	p := New(&l)
	program := p.ParseProgram()
	printErrs(p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement) // stmt is of type *ast.ExpressionStatement
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	integ, ok := stmt.Exprssn.(*ast.IntegerLiteral) // expression statement is of type *ast.IntegerLiteral
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Exprssn)
	}
	if integ.Value != 5 {
		t.Fatalf("integ.Value not %d. got=%d", 5, integ.Value)
	}
	if integ.TokenLiteral() != "5" {
		t.Fatalf("integ.Tok.Literal not %s. got=%s", "5", integ.Tok.Literal)
	}
}
