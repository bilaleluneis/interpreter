package ast

import (
	"goimpl/token"
	"strings"
)

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

func (l Let) Dump(ident int) string {
	out := "ast.Let{ //start of Let\n"
	indentation := strings.Repeat("\t", ident+1)
	out += indentation + "Tok: token.Token{Type: token.LET, Literal: \""
	out += l.Tok.Literal + "\"},\n"
	out += indentation + "Name: " + func() string {
		if l.Name != nil {
			return l.Name.Dump(ident + 1)
		}
		return "nil"
	}() + ",\n"
	out += indentation + "Value: " + func() string {
		if l.Value != nil {
			return l.Value.Dump(ident + 1)
		}
		return "nil"
	}() + "\n"
	out += strings.Repeat("\t", ident) + "} //end of Let"
	return out
}
