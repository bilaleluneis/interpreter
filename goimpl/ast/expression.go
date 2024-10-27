package ast

import (
	"goimpl/token"
	"strings"
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

func (e ExpressionStatement) Dump(ident int) string {
	out := "ast.ExpressionStatement{ //start of ExpressionStatment\n"
	indentation := strings.Repeat("\t", ident)
	out += indentation + "Tok: token.Token{ Type: " + e.Tok.String()
	out += ", Literal: " + e.Tok.Literal + "},\n"
	out += indentation + "Exprssn: " + func() string {
		if e.Exprssn != nil {
			return e.Exprssn.Dump(ident + 1)
		}
		return "nil"
	}() + "\n"
	out += strings.Repeat("\t", ident-1) + "} //end of ExpressionStatment"
	return out
}
