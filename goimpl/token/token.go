package token

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) isOneOf(types ...TokenType) bool {
	for _, typ := range types {
		if t.Type == typ {
			return true
		}
	}
	return false
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
