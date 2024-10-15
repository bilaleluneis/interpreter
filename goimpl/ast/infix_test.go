package ast

import (
	"goimpl/token"
	"testing"
)

func TestInfixExpressionAst(t *testing.T) {
	tests := []struct {
		input    InfixExpression
		expected string
	}{
		{
			input: InfixExpression{
				Left:     &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
				Tok:      token.Token{Type: token.PLUS, Literal: "+"},
				Operator: "+",
				Right:    &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
			},
			expected: "(5 + 5)",
		},
		{
			input: InfixExpression{
				Left:     &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
				Tok:      token.Token{Type: token.PLUS, Literal: "+"},
				Operator: "+",
				Right: &InfixExpression{
					Left:     &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
					Tok:      token.Token{Type: token.ASTER, Literal: "*"},
					Operator: "*",
					Right:    &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
				},
			},
			expected: "(5 + (5 * 5))",
		},
	}

	for _, tt := range tests {
		program := &Program{Statements: []Statement{&ExpressionStatement{Exprssn: &tt.input}}}
		actual := program.Top()
		expected := tt.expected
		if actual != expected {
			t.Errorf("expected %q, got %q", expected, actual)
		} else {
			t.Logf("program ast succes! %s", program)
		}
	}
}
