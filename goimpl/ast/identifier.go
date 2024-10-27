package ast

import (
	"goimpl/token"
	"strings"
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

func (i Identifier) Dump(ident int) string {
	out := strings.Repeat("\t", ident)
	out += "ast.Identifier{ //start of Identifier\n"
	out += strings.Repeat("\t", ident+1)
	out += "Tok: token.Token{ Type: token.IDENT, Literal: \"" + i.Tok.Literal + "\"},\n"
	out += strings.Repeat("\t", ident+1)
	out += "Value: \"" + i.Value + "\"\n"
	out += strings.Repeat("\t", ident)
	out += "} //end of Identifier"
	return out
}
