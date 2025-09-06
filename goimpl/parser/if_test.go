package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
	"time"
)

func TestIf(t *testing.T) {
	for name, fix := range ifTests {
		t.Run(name, func(t *testing.T) {
			done := make(chan struct{})
			go runIfTest(t, fix, done)
			select {
			case <-done:
				// test finished
			case <-time.After(500 * time.Second):
				t.Errorf("test %q timed out", name)
			}
		})
	}
}

func runIfTest(t *testing.T, fix ifTestCase, done chan struct{}) {
	defer close(done)

	lex := lexer.NewStubLexer(fix.tokens)
	var p ParserType = pratt.New(&lex)
	prog := p.ParseProgram()
	numStatements := len(prog.Statements)
	if numStatements != 1 {
		t.Errorf("program has wrong number of statements. got=%d", numStatements)
		return
	}

	switch ifExpr := prog.Statements[0].(type) {
	case *ast.Error:
		if fix.expectedErrMsg == "" {
			t.Errorf("unexpected error: %s", ifExpr.Message)
			return
		}
		if ifExpr.Message != fix.expectedErrMsg {
			t.Errorf("wrong error message. got=%q, want=%q", ifExpr.Message, fix.expectedErrMsg)
		}
	default:
		if fix.expectedIfExpression == "" {
			t.Error("expected error but got valid if expression")
			return
		}
		got := ifExpr.String()
		want := fix.expectedIfExpression
		if got != want {
			t.Errorf("wrong if expression string. got=%q, want=%q", got, want)
		}
	}
}
