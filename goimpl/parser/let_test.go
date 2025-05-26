package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
	"time"
)

func TestLet(t *testing.T) {
	for name, fix := range letStatementTests {
		t.Run(name, func(t *testing.T) {
			done := make(chan struct{})
			go runLetTest(t, fix, done)
			select {
			case <-done:
				return
			case <-time.After(500 * time.Millisecond):
				t.Error("test exceeded timeout of 500ms")
			}
		})
	}
}

func runLetTest(t *testing.T, fix letTestCase, done chan struct{}) {
	defer close(done)

	lex := lexer.NewStubLexer(fix.tokens)
	var p ParserType = pratt.New(&lex)
	prog := p.ParseProgram()
	numStatments := len(prog.Statements)
	if numStatments != 1 {
		t.Errorf("program has wrong number of statements. got=%d", numStatments)
		return
	}

	stmt := prog.Statements[0]

	switch stmt := stmt.(type) {
	case *ast.Error:
		if fix.expectedErrMsg == "" {
			t.Errorf("unexpected error: %s", stmt.Message)
			return
		}
		if stmt.Message != fix.expectedErrMsg {
			t.Errorf("wrong error message. got=%q, want=%q", stmt.Message, fix.expectedErrMsg)
		}
	case *ast.Let:
		if fix.expectedStatement == "" {
			t.Error("expected error but got let statement")
			return
		}
		statement := stmt.String()
		if statement != fix.expectedStatement {
			t.Errorf("wrong statement string. got=%q, want=%q", statement, fix.expectedStatement)
		}
	default:
		t.Errorf("unexpected statement type: %T", stmt)
	}
}
