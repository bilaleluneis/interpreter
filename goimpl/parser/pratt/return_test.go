package pratt

import (
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

// TODO: checks
// cast to see if statment is of type return
// returnStmt, ok := stmt.(*ast.ReturnStatement)
func TestReturnStatment(t *testing.T) {
	// return 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.RETURN, Literal: "return"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	p := New(&l)
	_, ok := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printErrs(p)
		t.FailNow()
	}
	if !ok {
		t.Fatalf("ParseProgram() failed")
	}
}
