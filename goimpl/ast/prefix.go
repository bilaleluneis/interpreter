package ast

import (
	"goimpl/token"
	"strings"
)

// PrefixExpression <PrefixExpression> ::= <prefix operator> <expression>
type PrefixExpression struct {
	Tok      token.Token // The prefix token, e.g. "!" or "-"
	Operator string
	Right    Expression
}

func (PrefixExpression) expressionNode() {}

func (pe PrefixExpression) TokenLiteral() string { return pe.Tok.Literal }

func (pe PrefixExpression) String() string {
	return "(" + pe.Operator + pe.Right.String() + ")"
}

func (pe PrefixExpression) Dump(ident int) string {
	out := "ast.PrefixExpression{ //start ofPrefixExpression\n"
	indentation := strings.Repeat("\t", ident+1)
	out += indentation + "Tok: token.Token{Type: " + pe.Tok.String()
	out += ", Literal: " + pe.Tok.Literal + "},\n"
	out += indentation + "Operator: " + pe.Operator + ",\n"
	out += indentation + "Right: " + func() string {
		if pe.Right != nil {
			return pe.Right.Dump(ident + 1)
		}
		return "nil"
	}()
	out += "\n" + strings.Repeat("\t", ident) + "} //end of PrefixExpression"
	return out
}
