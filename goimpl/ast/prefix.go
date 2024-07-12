package ast

import (
	"goimpl/debug/treedrawer/tree"
	"goimpl/token"
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

func visualizePrefixExpression(pe PrefixExpression, parent *tree.Tree) {
	operator := parent.AddChild(tree.NodeString(pe.Operator))
	visualizeExpression(pe.Right, operator)
}
