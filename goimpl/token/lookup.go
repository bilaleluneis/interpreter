package token

import (
	"regexp"
	"strconv"
)

// Lookup returns the corresponding token for a given identifier.
// TODO: write up a test for this function
func Lookup(ident string) Token {
	switch {
	case ident == "+":
		return Token{Type: PLUS, Literal: ident}
	case ident == "-":
		return Token{Type: MINUS, Literal: ident}
	case ident == "*":
		return Token{Type: ASTER, Literal: ident}
	case ident == "/":
		return Token{Type: SLASH, Literal: ident}
	case ident == "<":
		return Token{Type: LT, Literal: ident}
	case ident == ">":
		return Token{Type: GT, Literal: ident}
	case ident == "==":
		return Token{Type: EQ, Literal: ident}
	case ident == "!=":
		return Token{Type: NEQ, Literal: ident}
	case ident == "=":
		return Token{Type: ASSIGN, Literal: ident}
	case ident == ",":
		return Token{Type: COMMA, Literal: ident}
	case ident == ";":
		return Token{Type: SEMICOLON, Literal: ident}
	case ident == "(":
		return Token{Type: LPRAN, Literal: ident}
	case ident == ")":
		return Token{Type: RPRAN, Literal: ident}
	case ident == "{":
		return Token{Type: LBRACE, Literal: ident}
	case ident == "}":
		return Token{Type: RBRACE, Literal: ident}
	case ident == "!":
		return Token{Type: BANG, Literal: ident}
	case ident == "fn":
		return Token{Type: FUNCTION, Literal: ident}
	case ident == "let":
		return Token{Type: LET, Literal: ident}
	case ident == "return":
		return Token{Type: RETURN, Literal: ident}
	case ident == "if":
		return Token{Type: IF, Literal: ident}
	case ident == "else":
		return Token{Type: ELSE, Literal: ident}
	case ident == "true":
		return Token{Type: TRUE, Literal: ident}
	case ident == "false":
		return Token{Type: FALSE, Literal: ident}
	case func() bool { _, err := strconv.ParseUint(ident, 0, 64); return err == nil }():
		return Token{Type: INT, Literal: ident}
	case regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`).MatchString(ident):
		return Token{Type: IDENTIFIER, Literal: ident}
	default:
		return Token{Type: ILLIGAL, Literal: ident}
	}
}
