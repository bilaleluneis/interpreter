package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

func TestLetFunc(t *testing.T) {
	//let x = 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	// call the Let function and pass copy of lexer
	// similar to what combinator.Parser.parse() does
	lexerUnderTest := l.GetCopy()
	result := Let(lexerUnderTest)

	// check that the lexer we got back is not same instance as passed
	if &l == result.lxr {
		t.Fatalf("expected different lexer instance got same")
	}

	// check that the original lexer is not advanced
	if l.NextToken().Type != token.LET {
		t.Fatalf("original lexer has advanced")
	}

	// check that lexer we got back emits token.EOF on advance
	if result.lxr.NextToken().Type != token.EOF {
		t.Fatalf("expected EOF got %s", result.lxr.NextToken().Literal)
	}

	// check that the statement is a let statement
	if _, ok := result.stmnt.(ast.Let); !ok {
		t.Fatalf("expected let statement got %T", result.stmnt)
	}
	//TODO: might want to check the values of the let statement
}
