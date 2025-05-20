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
			go func() {
				defer close(done)

				lex := lexer.NewStubLexer(fix.tokens)
				var p ParserType = pratt.New(&lex)
				prog := p.ParseProgram()

				numStatements := len(prog.Statements)
				if numStatements != 1 {
					t.Errorf("program has wrong number of statements. got=%d", numStatements)
					return
				}

				stmt := prog.Statements[0]

				switch stmt := stmt.(type) {
				case *ast.Error:
					if !fix.wantErr {
						t.Errorf("got unexpected error: %s", stmt.Message)
						return
					}
					if stmt.Message != fix.errorMsg {
						got := stmt.Message
						want := fix.errorMsg
						t.Errorf("wrong error message. got=%q, want=%q", got, want)
					}
				case *ast.Let:
					if fix.wantErr {
						t.Error("expected error but got let statement")
						return
					}
					if fix.expected != "" && stmt.String() != fix.expected {
						got := stmt.String()
						want := fix.expected
						t.Errorf("wrong statement string. got=%q, want=%q", got, want)
					}
				default:
					t.Errorf("statement is not *ast.Let or *ast.Error. got=%T", stmt)
				}
			}()

			select {
			case <-done:
				return
			case <-time.After(500 * time.Millisecond):
				t.Error("test exceeded timeout of 500ms")
			}
		})
	}
}
