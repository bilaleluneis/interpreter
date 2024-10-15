package ast

import "goimpl/token"

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

func (pe PrefixExpression) Dump() string {
	out := `ast.PrefixExpression{
	Tok: ` + pe.Tok.Literal + `,
	Operator: ` + pe.Operator + `,
	Right: `
	if pe.Right != nil {
		out += pe.Right.Dump()
	} else {
		out += "nil"
	}
	out += `
}`
	return out
}
