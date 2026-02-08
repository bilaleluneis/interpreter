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

	tokens = append(tokens, NewToken(EOF, 0))
	tokens[0] = NewToken(LET, 'l')

	if tokens.Len() != 6 {
		t.Errorf("Expected 5 tokens, got %d", tokens.Len())
	}
}
