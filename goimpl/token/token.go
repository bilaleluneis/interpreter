package token

import "slices"

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) isOneOf(types ...TokenType) bool {
	return slices.Contains(types, t.Type)
}

func (t Token) String() string {
	return t.Literal
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{tokenType, string(ch)}
}

//goland:noinspection GoNameStartsWithPackageName
type TokenType string

//goland:noinspection GoCommentStart
const (
	ILLIGAL TokenType = "ILLIGAL"
	EOF     TokenType = "EOF"

	//Identifiers + literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"

	//Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"
	MINUS  TokenType = "-"
	BANG   TokenType = "!"
	ASTER  TokenType = "*"
	SLASH  TokenType = "/"
	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NEQ    TokenType = "!="

	//Delimeters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPRAN     TokenType = "("
	RPRAN     TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"

	//Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	RETURN   TokenType = "RETURN"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
)

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
