package ast

import (
	"goimpl/token"
)

// InvalidExpression represents an invalid expression encountered during parsing
type InvalidExpression struct {
	Tok     token.Token // the token that caused the invalid expression
	Message string      // error message describing why the expression is invalid
}

func (InvalidExpression) expressionNode() {}

func (i InvalidExpression) TokenLiteral() string {
	return i.Tok.Literal
}

func (i InvalidExpression) String() string {
	return "invalid expression: " + i.Message
}
