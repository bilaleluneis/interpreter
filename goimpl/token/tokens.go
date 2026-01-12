package token

import "strings"

type Tokens []Token

func (t Tokens) Len() int {
	return len(t)
}

func (t Tokens) String() string {
	var result strings.Builder
	result.WriteString("[ ")
	for _, token := range t {
		result.WriteString(token.String() + " ")
	}
	result.WriteString("]")

	return result.String()
}
