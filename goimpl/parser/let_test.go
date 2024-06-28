package parser

import (
	"goimpl/token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {

	// let x = 5;
	l := &stubLexer{
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
	}

	p := New(l)
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printErrs(p)
		t.FailNow()
	}
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
}

func TestParseLetStatementError(t *testing.T) {
	// let x 5;
	l := &stubLexer{
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
	}

	p := New(l)
	_ = p.ParseProgram()
	if len(p.Errors()) == 0 {
		t.Fatalf("expected error")
	}
	printErrs(p)
}
