package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

// TestFailFunc tests the Fail function from the fail.go file
// function should return back lexer with no advancement in token
func TestFailFunc(t *testing.T) {
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	// call the Fail function and pass copy of lexer
	// similar to what combinator.Parser.parse() does
	result := Fail(l.GetCopy())

	// check that the lexer we got back is not same instance as passed
	if &l == &result.lxr {
		t.Fatalf("expected different lexer instance got same")
	}

	// check that the lexer we got back is not advanced
	if l.NextToken().Type != result.lxr.NextToken().Type {
		t.Fatalf("expected same token type got different")
	}

	// check that the statement is an error
	if _, ok := result.stmnt.(ast.Error); !ok {
		t.Fatalf("expected error statement got %T", result.stmnt)
	}
}

func TestFailParse(t *testing.T) {
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	result := New(l, Fail[lexer.StubLexer]).ParseProgram()

	// check that we got one statment
	if len(result.Statements) != 1 {
		t.Fatalf("expected 1 statement got %d", len(result.Statements))
	}

	// check that the statement is an error
	if _, ok := result.Statements[0].(ast.Error); !ok {
		t.Fatalf("expected error statement got %T", result.Statements[0])
	}

	// check that the error message is correct
	if result.Statements[0].TokenLiteral() != "No valid statement parsed" {
		t.Fatalf("expected [No valid statement parsed] got %s", result.Statements[0].TokenLiteral())
	}
}
