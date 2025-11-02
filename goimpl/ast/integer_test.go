package ast

import (
	"goimpl/token"
	"testing"
)

func TestIntegerLiteral_TokenLiteralAndString(t *testing.T) {
	cases := []struct {
		lit string
		val int64
	}{
		{"5", 5},
		{"0", 0},
	}

	for _, c := range cases {
		t.Run(c.lit, func(t *testing.T) {
			tok := token.Token{Type: token.INT, Literal: c.lit}
			il := IntegerLiteral{Tok: tok, Value: c.val}

			if got := il.TokenLiteral(); got != c.lit {
				t.Fatalf("TokenLiteral() = %q, want %q", got, c.lit)
			}

			if got := il.String(); got != c.lit {
				t.Fatalf("String() = %q, want %q", got, c.lit)
			}
		})
	}
}
