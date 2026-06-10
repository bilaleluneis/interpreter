package token

import (
	"fmt"
	"testing"
)

func Test_iteration(t *testing.T) {
	tokens := NewTokens(
		NewToken(LET, 'l'),
		NewToken(IDENTIFIER, 'x'),
		NewToken(ASSIGN, '='),
		NewToken(INT, '5'),
		NewToken(SEMICOLON, ';'),
	)

	var i Token
	for i = range tokens.All() {
		fmt.Print(i)
	}
	expect := Token{Type: SEMICOLON, Literal: ";"}

	if i != expect {
		t.Errorf("expected last token to be SEMICOLON, got %v", i)
	}

}

// test that multiple iterations over the same tokens value work correctly
func Test_multiple_iterations(t *testing.T) {
	tokens := NewTokens(
		NewToken(LET, 'l'),
		NewToken(IDENTIFIER, 'x'),
		NewToken(ASSIGN, '='),
		NewToken(INT, '5'),
		NewToken(SEMICOLON, ';'),
	)

	var i Token
	for i = range tokens.All() {
		fmt.Print(i)
	}
	expect := Token{Type: SEMICOLON, Literal: ";"}

	if i != expect {
		t.Errorf("expected last token to be SEMICOLON, got %v", i)
	}

	for i = range tokens.All() {
		fmt.Print(i)
	}

	if i != expect {
		t.Errorf("expected last token to be SEMICOLON, got %v", i)
	}
}
