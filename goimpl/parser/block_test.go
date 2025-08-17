package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
)

func TestBlock(t *testing.T) {
	for name, fix := range blockTests {
		t.Run(name, func(t *testing.T) {
			lexr := lexer.NewStubLexer(fix.tokens)
			var p ParserType = pratt.New(&lexr)
			program := p.ParseProgram()

			numStatements := len(program.Statements)
			if numStatements != 1 {
				t.Errorf("program has wrong number of statments. got=%d", numStatements)
				return
			}

			stmt := program.Statements[0]

			switch stmt := stmt.(type) {
			case *ast.Error:
				if fix.expectedErrrMsg == "" {
					t.Errorf("unexpected error: %s", stmt.Message)
					return
				}
				expErr := fix.expectedErrrMsg
				if stmt.Message != expErr {
					t.Errorf("wrong error message. got=%q, want=%q", stmt.Message, expErr)
				}
			case *ast.Block:
				if fix.expectedBlock == "" {
					t.Error("expected error but got block statement")
					return
				}
				blockStmt := stmt.String()
				expBlock := fix.expectedBlock
				if blockStmt != expBlock {
					t.Errorf("wrong block string. got=%q, want=%q", blockStmt, expBlock)
				}
			default:
				t.Errorf("unexpected statement type: %T", stmt)
			}
		})
	}
}
