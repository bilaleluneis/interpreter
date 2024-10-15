package ast

import "goimpl/token"

type Return struct {
	Tok   token.Token // the token.RETURN token
	Value Expression
}

func (Return) statmentNode() {}

func (r Return) TokenLiteral() string { return r.Tok.Literal }

func (r Return) String() string {
	out := r.Tok.Literal + " "
	if r.Value != nil {
		out += r.Value.String()
	}
	return out + ";"
}

func (r Return) Dump() string {
	return `ast.Return{
	Tok: token.Token{ Type: token.RETURN, Literal: ` + r.Tok.Literal + `,
    Value: ` + r.Value.Dump() + `
}`
}
