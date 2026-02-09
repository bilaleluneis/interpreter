package token

import "testing"

func Test_immutability(t *testing.T) { //nolint:paralleltest
	tokens := NewTokens(
		NewToken(LET, 'l'),
		NewToken(IDENTIFIER, 'x'),
		NewToken(ASSIGN, '='),
		NewToken(INT, '5'),
		NewToken(SEMICOLON, ';'),
	)

	tokensCopy := FromTokens(tokens)

	tokens.Set(0, NewToken(ELSE, 'e'))

	if tokensCopy.Get(0).Type != LET {
		t.Errorf("Expected original tokens to be unchanged, but it was modified")
	}
}
