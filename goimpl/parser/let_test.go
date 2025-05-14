package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
	"time"
)

// FIXME: review bellow test and checks once you update statment.go parseLetStatement
func TestLet(t *testing.T) {
	for name, fix := range letStatementTests {
		t.Run(name, func(t *testing.T) {
			done := make(chan struct{})
			go func() {
				defer close(done)

				lex := lexer.NewStubLexer(fix.tokens)
				var p ParserType = pratt.New(&lex)
				prog, ok := p.ParseProgram()

				if fix.wantErr {
					if ok {
						t.Error("expected parser errors but got none")
					}
					return
				}

				if !ok {
					t.Error("parser had errors")
					return
				}

				if len(prog.Statements) != 1 {
					t.Errorf("program has wrong number of statements. got=%d", len(prog.Statements))
					return
				}

				stmt, ok := prog.Statements[0].(*ast.Let)
				if !ok {
					t.Errorf("program.Statements[0] is not *ast.Let. got=%T", prog.Statements[0])
					return
				}

				if fix.checkVal {
					if stmt.Name == nil || stmt.Name.Value != "x" {
						t.Errorf("let statement name not 'x'. got=%v", stmt.Name)
					}

					if stmt.Value == nil {
						t.Error("let statement value is nil")
					}

					if fix.expected != "" && stmt.String() != fix.expected {
						t.Errorf("wrong statement string. got=%q, want=%q", stmt.String(), fix.expected)
					}
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
