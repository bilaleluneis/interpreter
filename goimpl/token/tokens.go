package token

import "strings"

type tokens struct {
	toks []Token
}

func (t tokens) Len() int {
	return len(t.toks)
}

func (t tokens) String() string {
	var result strings.Builder
	result.WriteString("[ ")
	for _, token := range t.toks {
		result.WriteString(token.String() + " ")
	}
	result.WriteString("]")

	return result.String()
}

func NewTokens(tks ...Token) tokens {
	return tokens{
		toks: tks,
	}
}
