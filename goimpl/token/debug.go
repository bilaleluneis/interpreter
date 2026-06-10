package token

import "strings"

func (t Token) String() string {
	return "Token{" + t.Literal + "}"
}

func (t tokens) String() string {
	var result strings.Builder
	result.WriteString("tokens:[ ")
	for _, token := range t.toks {
		result.WriteString(token.String())
		result.WriteString(" ")
	}
	result.WriteString("]")

	return result.String()
}
