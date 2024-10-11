package ast

import (
	"goimpl/token"
	"testing"
)

func TestPrefixExpressionAst(t *testing.T) {
	tests := []struct {
		input    PrefixExpression
		expected string
	}{
		{
			input: PrefixExpression{
				Tok:      token.Token{Type: token.MINUS, Literal: "-"},
				Operator: "-",
				Right:    &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
			},
			expected: "(-5)",
		},
		{
			input: PrefixExpression{
				Tok:      token.Token{Type: token.BANG, Literal: "!"},
				Operator: "!",
				Right:    &IntegerLiteral{Value: 5, Tok: token.Token{Type: token.INT, Literal: "5"}},
			},
			expected: "(!5)",
		},
	}

	for _, tt := range tests {
		program := &Program{Statements: []Statement{&ExpressionStatement{Exprssn: &tt.input}}}
		actual := program.top()
		expected := tt.expected
		if actual != expected {
			t.Errorf("expected %q, got %q", expected, actual)
		} else {
			t.Logf("program ast succes! %s", program)
		}
	}
}
