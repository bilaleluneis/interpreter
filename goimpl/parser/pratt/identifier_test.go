package pratt

import (
	"goimpl/ast"
	"goimpl/lexer"
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

	p := New(&l)
	program, _ := p.ParseProgram()
	printErrs(p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement) // stmt is of type *ast.ExpressionStatement
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := stmt.Exprssn.(*ast.Identifier) // expression statement is of type *ast.Identifier
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Exprssn)
	}
	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.Tok.Literal != "foobar" {
		t.Fatalf("ident.Tok.Literal not %s. got=%s", "foobar", ident.Tok.Literal)
	}
}
