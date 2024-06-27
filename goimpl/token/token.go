package token

type Token struct {
	Type    TokenType
	Literal string
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
)
