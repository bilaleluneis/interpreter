package ast

import (
	"goimpl/token"
)

type Let struct {
	Tok   token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (Let) statmentNode() {}

func (l Let) TokenLiteral() string { return l.Tok.Literal }

func (l Let) String() string {
	if l.Name != nil && l.Value != nil {
		return "let " + l.Name.Value + " = " + l.Value.String() + ";"
	}
	return "let <incomplete>"
}
