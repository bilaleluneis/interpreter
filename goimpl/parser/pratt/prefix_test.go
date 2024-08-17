package pratt

import (
	"goimpl/ast"
	"goimpl/lexer"
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
}{
	// -5;
	{"-", 5, lexer.NewStubLexer([]token.Token{
		{Type: token.MINUS, Literal: "-"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}})},

	// !5;
	{"!", 5, lexer.StubLexer([]token.Token{
		{Type: token.BANG, Literal: "!"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""}})},
}

func TestPrefix(t *testing.T) {
	for _, tt := range prefixTests {
		expectedOperator := tt.expectedOperator
		expectedValue := tt.expectedValue
		lexr := tt.lexr
		p := New(&lexr)
		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printErrs(p)
			t.FailNow()
		}
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement) // stmt is of type *ast.ExpressionStatement
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		prefix, ok := stmt.Exprssn.(*ast.PrefixExpression) // expression statement is of type *ast.PrefixExpression
		if !ok {
			t.Fatalf("exp not *ast.PrefixExpression. got=%T", stmt.Exprssn)
		}
		if prefix.Operator != expectedOperator {
			t.Fatalf("prefix.Operator not %s. got=%s", expectedOperator, prefix.Operator)
		}
		intLit, ok := prefix.Right.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("prefix.Right not *ast.IntegerLiteral. got=%T", prefix.Right)
		}
		if intLit.Value != expectedValue {
			t.Fatalf("prefix.Right.Value not %d. got=%d", expectedValue, intLit.Value)
		}
	}
}
