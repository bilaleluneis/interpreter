package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

var infixTests = []struct {
	expectedLeft  string
	expectedOp    string
	expectedRight string
	lexr          lexer.LexerType
}{
	// 5 + 5;
	{"5", "+", "5", &stubLexer{
		tokens: []token.Token{
			{Type: token.INT, Literal: "5"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
	}},
}

func TestInfix(t *testing.T) {
	for _, tt := range infixTests {
		expectedLeft := tt.expectedLeft
		expectedOp := tt.expectedOp
		expectedRight := tt.expectedRight
		lexr := tt.lexr
		p := New(lexr)
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
		infix, ok := stmt.Exprssn.(*ast.InfixExpression) // expression statement is of type *ast.InfixExpression
		if !ok {
			t.Fatalf("exp not *ast.InfixExpression. got=%T", stmt.Exprssn)
		}
		if infix.Operator != expectedOp {
			t.Fatalf("infix.Operator is not '%s'. got=%s", expectedOp, infix.Operator)
		}
		if infix.Left.String() != expectedLeft {
			t.Fatalf("infix.Left is not %s. got=%s", expectedLeft, infix.Left.String())
		}
		if infix.Right.String() != expectedRight {
			t.Fatalf("infix.Right is not %s. got=%s", expectedRight, infix.Right.String())
		}
	}
}
