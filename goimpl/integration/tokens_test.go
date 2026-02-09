package integration

import (
	. "goimpl/token"
	"testing"
)

func TestXxx(t *testing.T) {
	tokens := NewTokens(
		NewToken(LET, 'l'),
		NewToken(IDENTIFIER, 'x'),
		NewToken(ASSIGN, '='),
		NewToken(INT, '5'),
		NewToken(SEMICOLON, ';'),
	)

	copy := FromTokens(tokens)

	copy.Set(0, NewToken(ELSE, 'e'))

	if tokens.Get(0).Type != LET {
		t.Errorf("Expected original tokens to be unchanged, but it was modified")
	}
}
