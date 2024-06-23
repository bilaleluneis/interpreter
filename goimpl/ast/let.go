package ast

import "goimpl/token"

type Let struct {
	Tok   token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (l Let) statment()            {}
func (l Let) TokenLiteral() string { return l.Tok.Literal }
