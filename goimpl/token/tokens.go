package token

import "strings"

type tokens struct {
	toks []Token
}

func (t tokens) Len() int {
	return len(t.toks)
}

func (t tokens) Append(tk Token) {
	_ = append(t.toks, tk)
}

func (t tokens) Get(i int) Token {
	return t.toks[i]
}

func (t tokens) Set(i int, tk Token) {
	t.toks[i] = tk
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

func FromTokens(tks tokens) tokens {
	newTks := make([]Token, len(tks.toks))
	copy(newTks, tks.toks)
	return tokens{
		toks: newTks,
	}
}
