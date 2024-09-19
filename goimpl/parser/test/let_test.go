package test

import (
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {
	// let x = 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	for pname, parser := range testParsers(l).initPratt().initCombinator(combinator.Let).parsers {
		program, ok := parser.ParseProgram()
		if !ok {
			t.Errorf("\nexpected ok program for parser %s, got: !ok", pname)
		} else {
			t.Logf("\npass--parser %s, got: %s", pname, program)
		}
	}
}

func TestParseLetStatementError(t *testing.T) {
	// let x 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	for pname, parser := range testParsers(l).initPratt().initCombinator(combinator.Let).parsers {
		if program, ok := parser.ParseProgram(); ok {
			t.Errorf("\nexpected !ok program for parser %s, got: %s", pname, program)
		}
	}
}
