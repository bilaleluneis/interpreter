package ast

import (
	"goimpl/token"
	"testing"
)

func TestProgram(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&Let{
				Tok: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Tok:   token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Tok:   token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.TokenLiteral() != "let" {
		t.Errorf("program.TokenLiteral not 'let'. got=%q", program.TokenLiteral())
	}
}
