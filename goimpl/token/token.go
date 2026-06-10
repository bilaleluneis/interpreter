// Package token defines the token types for interpreter.
package token

import "slices"

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) isOneOf(types ...TokenType) bool {
	return slices.Contains(types, t.Type)
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{tokenType, string(ch)}
}
