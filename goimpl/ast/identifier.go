package ast

import (
	"goimpl/token"
)

// Identifier can also be an Expression
type Identifier struct {
	Tok   token.Token // the token.IDENT token
	Value string
}

func (Identifier) expressionNode() {}

func (i Identifier) TokenLiteral() string {
	return i.Tok.Literal
}

func (i Identifier) String() string {
	return i.Value
}
