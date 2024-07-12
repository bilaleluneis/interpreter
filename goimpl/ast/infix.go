package ast

import (
	"goimpl/debug/treedrawer/tree"
	"goimpl/token"
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

func visualizeInfixExpression(ie InfixExpression, parent *tree.Tree) {
	operator := parent.AddChild(tree.NodeString(ie.Operator))
	visualizeExpression(ie.Left, operator)
	visualizeExpression(ie.Right, operator)
}
