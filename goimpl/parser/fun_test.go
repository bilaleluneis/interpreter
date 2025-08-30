package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
	"time"
)

func TestFun(t *testing.T) {
	for name, fix := range funTests {
		t.Run(name, func(t *testing.T) {
			done := make(chan struct{})
			go runFunTest(t, fix, done)
			select {
			case <-done:
				// test finished
			case <-time.After(500 * time.Second):
				t.Errorf("test %q timed out", name)
			}
		})
	}
}

func runFunTest(t *testing.T, fix funTestCase, done chan struct{}) {
	defer close(done)

	lex := lexer.NewStubLexer(fix.tokens)
	var p ParserType = pratt.New(&lex)
	prog := p.ParseProgram()
	numStatements := len(prog.Statements)
	if numStatements != 1 {
		t.Errorf("program has wrong number of statements. got=%d", numStatements)
		return
	}

	switch fun := prog.Statements[0].(type) {
	case *ast.Error:
		if fix.expectedErrMsg == "" {
			t.Errorf("unexpected error: %s", fun.Message)
			return
		}
		if fun.Message != fix.expectedErrMsg {
			t.Errorf("wrong error message. got=%q, want=%q", fun.Message, fix.expectedErrMsg)
		}
	default:
		if fix.expectedFunExpression == "" {
			t.Error("expected error but got valid function")
			return
		}
		got := fun.String()
		want := fix.expectedFunExpression
		if got != want {
			t.Errorf("wrong function string. got=%q, want=%q", got, want)
		}
	}
}
