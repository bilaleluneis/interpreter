package ast

import (
	"goimpl/debug/treedrawer/tree"
	"goimpl/token"
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

func visualizeExpressionStatement(exs ExpressionStatement, parent *tree.Tree) {
	expression := exs.Exprssn
	if expression != nil {
		visualizeExpression(expression, parent)
	}
}

func visualizeExpression(ex Expression, parent *tree.Tree) {
	switch exprType := ex.(type) {
	case *InfixExpression:
		visualizeInfixExpression(*exprType, parent)
	case *PrefixExpression:
		visualizePrefixExpression(*exprType, parent)
	case *IntegerLiteral:
		visualizeIntegerLiteral(*exprType, parent)
	case *Identifier:
		visualizeIdentifier(*exprType, parent)
	}
}
