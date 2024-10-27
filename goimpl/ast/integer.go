package ast

import (
	"goimpl/token"
	"strings"
)

type IntegerLiteral struct {
	Tok   token.Token // the token.INT token
	Value int64
}

func (IntegerLiteral) expressionNode() {}

func (il IntegerLiteral) TokenLiteral() string { return il.Tok.Literal }

func (il IntegerLiteral) String() string { return il.Tok.Literal }

func (il IntegerLiteral) Dump(ident int) string {
	out := "ast.IntegerLiteral{ //start of IntegerLiteral\n"
	indentation := strings.Repeat("\t", ident+1)
	out += indentation + "Tok: token.Token{ Type: token.INT, Literal: "
	out += il.Tok.Literal + "},\n"
	out += indentation + "Value: " + il.Tok.Literal + "\n"
	out += strings.Repeat("\t", ident) + "} //end of IntegerLiteral"
	return out
}
