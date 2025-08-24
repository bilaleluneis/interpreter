package ast

import (
	"goimpl/token"
)

type ExpressionStatement struct {
	Tok     token.Token // the first token of the Exprssn
	Exprssn Expression
}

func (ExpressionStatement) statmentNode() {}

func (e ExpressionStatement) TokenLiteral() string {
	return e.Tok.Literal
}

func (e ExpressionStatement) String() string {
	if e.Exprssn != nil {
		return e.Exprssn.String()
	}
	return ""
}
