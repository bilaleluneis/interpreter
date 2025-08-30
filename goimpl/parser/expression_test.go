// filepath: /root/developer/interpreter/goimpl/parser/expression_test.go
package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
	"time"
)

func TestExpression(t *testing.T) {
	for name, fix := range expressionTests {
		t.Run(name, func(t *testing.T) {
			done := make(chan struct{})
			go runExpressionTest(t, fix, done)
			select {
			case <-done:
				// test finished
			case <-time.After(500 * time.Microsecond):
				t.Errorf("test %q timed out", name)
			}
		})
	}
}

func runExpressionTest(t *testing.T, fix expressionTestCase, done chan struct{}) {
	defer close(done)

	lex := lexer.NewStubLexer(fix.tokens)
	var p ParserType = pratt.New(&lex)
	prog := p.ParseProgram()
	numStatements := len(prog.Statements)
	if numStatements != 1 {
		t.Errorf("program has wrong number of statements. got=%d", numStatements)
		return
	}

	switch expr := prog.Statements[0].(type) {
	case *ast.Error:
		if fix.expectedErrMsg == "" {
			t.Errorf("unexpected error: %s", expr.Message)
			return
		}
		if expr.Message != fix.expectedErrMsg {
			t.Errorf("wrong error message. got=%q, want=%q", expr.Message, fix.expectedErrMsg)
		}
	default:
		if fix.expectedExpression == "" {
			t.Error("expected error but got valid expression")
			return
		}
		got := expr.String()
		want := fix.expectedExpression
		if got != want {
			t.Errorf("wrong expression string. got=%q, want=%q", got, want)
		}
	}
}
