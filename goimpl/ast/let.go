package ast

import "goimpl/token"

type Let struct {
	Tok   token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (l Let) statmentNode() {}

func (l Let) TokenLiteral() string { return l.Tok.Literal }

func (l Let) String() string {
	out := l.Tok.Literal + " " + l.Name.String() + " = "
	if l.Value != nil {
		out += l.Value.String()
	}
	return out + ";"
}
