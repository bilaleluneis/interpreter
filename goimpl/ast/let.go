package ast

import "goimpl/token"

type Let struct {
	Tok   token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (Let) statmentNode() {}

func (l Let) TokenLiteral() string { return l.Tok.Literal }

func (l Let) String() string {
	out := l.Tok.Literal + " " + l.Name.String() + " = "
	if l.Value != nil {
		out += l.Value.String()
	}
	return out + ";"
}

func (l Let) Dump() string {
	return `ast.Let{
	Tok: ` + l.Tok.Literal + `,
	Name: ` + l.Name.Dump() + `,
	Value: ` + l.Value.Dump() + `
}`
}
