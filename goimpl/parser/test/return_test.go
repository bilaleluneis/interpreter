package test

import (
	"fmt"
	"goimpl/lexer"
	"goimpl/parser/combinator"
	"goimpl/token"
	"testing"
)

// TODO: test doesnt not check Value on return statement
func TestReturnStatment(t *testing.T) {
	// return 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.RETURN, Literal: "return"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	for pname, parser := range testParsers(&l).initPratt().initCombinator(combinator.Retrn).parsers {
		fmt.Printf("invoking parser: %s\n", pname)
		if program, ok := parser.ParseProgram(); !ok {
			fail(pname, t, "got not ok program")
		} else if toReturn(program.Statements[0]) == nil {
			fail(pname, t, "expected return statement")
		} else {
			success(pname, t, "parser %s passed with result %s", pname, program)
		}
	}
}
