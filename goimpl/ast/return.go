package ast

import (
	"goimpl/token"
	"strings"
)

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

func (r Return) Dump(ident int) string {
	out := strings.Repeat(" ", ident) + "ast.Return{ //start of Return\n"
	out += strings.Repeat(" ", ident+1)
	out += "Tok: token.Token{ Type: token.RETURN, Literal: \""
	out += r.Tok.Literal + "\" },\n"
	out += strings.Repeat(" ", ident+1) + "Value: "
	if r.Value != nil {
		out += r.Value.Dump(ident + 1)
	} else {
		out += "nil"
	}
	out += "\n" + strings.Repeat(" ", ident) + "} //end of Return"
	return out
}
