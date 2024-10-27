package ast

import (
	"goimpl/token"
	"strings"
)

// InfixExpression <InfixExpression> ::= <expression> <infix operator> <expression>
type InfixExpression struct {
	Left     Expression
	Tok      token.Token // The operator token, e.g. "+"
	Operator string
	Right    Expression
}

func (InfixExpression) expressionNode() {}

func (ie InfixExpression) TokenLiteral() string { return ie.Tok.Literal }

func (ie InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Operator + " " + ie.Right.String() + ")"
}

func (ie InfixExpression) Dump(ident int) string {
	out := "ast.InfixExpression{ //start of InfixExpression \n"
	indentation := strings.Repeat("\t", ident+1)
	out += indentation + "Left: " + func() string {
		if ie.Left != nil {
			return ie.Left.Dump(ident + 1)
		}
		return "nil"
	}()
	out += ",\n"
	out += indentation + "Tok: token.Token{Type: " + ie.Tok.String()
	out += ", Literal: " + ie.Tok.Literal + "},\n"
	out += indentation + "Operator: " + ie.Operator + ",\n"
	out += indentation + "Right: " + func() string {
		if ie.Right != nil {
			return ie.Right.Dump(ident + 1)
		}
		return "nil"
	}()
	out += "\n" + strings.Repeat("\t", ident) + "} //end of InfixExpression"
	return out
}
