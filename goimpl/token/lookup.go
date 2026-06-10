package token

func Lookup(ident string) Token {
	switch ident {
	case "+":
		return Token{Type: PLUS, Literal: ident}
	case "-":
		return Token{Type: MINUS, Literal: ident}
	case "*":
		return Token{Type: ASTER, Literal: ident}
	case "/":
		return Token{Type: SLASH, Literal: ident}
	case "<":
		return Token{Type: LT, Literal: ident}
	case ">":
		return Token{Type: GT, Literal: ident}
	case "==":
		return Token{Type: EQ, Literal: ident}
	case "!=":
		return Token{Type: NEQ, Literal: ident}
	case "=":
		return Token{Type: ASSIGN, Literal: ident}
	case ",":
		return Token{Type: COMMA, Literal: ident}
	case ";":
		return Token{Type: SEMICOLON, Literal: ident}
	case "(":
		return Token{Type: LPRAN, Literal: ident}
	case ")":
		return Token{Type: RPRAN, Literal: ident}
	case "{":
		return Token{Type: LBRACE, Literal: ident}
	case "}":
		return Token{Type: RBRACE, Literal: ident}
	case "!":
		return Token{Type: BANG, Literal: ident}
	case "fn":
		return Token{Type: FUNCTION, Literal: ident}
	case "let":
		return Token{Type: LET, Literal: ident}
	case "return":
		return Token{Type: RETURN, Literal: ident}
	case "if":
		return Token{Type: IF, Literal: ident}
	case "else":
		return Token{Type: ELSE, Literal: ident}
	case "true":
		return Token{Type: TRUE, Literal: ident}
	case "false":
		return Token{Type: FALSE, Literal: ident}
	default:
		return Token{Type: IDENTIFIER, Literal: ident}
	}
}
