package ast

import "goimpl/token"

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

func (ie InfixExpression) Dump() string {
	out := `ast.InfixExpression{
	Left: ` + ie.Left.Dump() + `,
	Tok: ` + ie.Tok.Literal + `,
	Operator: ` + ie.Operator + `,
	Right: `
	if ie.Right != nil {
		out += ie.Right.Dump()
	} else {
		out += "nil"
	}
	out += `
}`
	return out
}
