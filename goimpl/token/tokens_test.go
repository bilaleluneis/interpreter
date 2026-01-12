package token

import "testing"

func TestTokens_String(t *testing.T) { //nolint:paralleltest
	tokens := Tokens{
		{Type: IDENTIFIER, Literal: "foo"},
		{Type: ASSIGN, Literal: "="},
		{Type: INT, Literal: "42"},
		{Type: SEMICOLON, Literal: ";"},
	}

	expected := "[ foo = 42 ; ]"
	if tokens.String() != expected {
		t.Errorf("expected %q, got %q", expected, tokens.String())
	}
}
