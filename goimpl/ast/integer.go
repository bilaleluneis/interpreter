package ast

import (
	"goimpl/debug/treedrawer/tree"
	"goimpl/token"
)

type IntegerLiteral struct {
	Tok   token.Token // the token.INT token
	Value int64
}

func (IntegerLiteral) expressionNode() {}

func (il IntegerLiteral) TokenLiteral() string { return il.Tok.Literal }

func (il IntegerLiteral) String() string { return il.Tok.Literal }

func visualizeIntegerLiteral(il IntegerLiteral, parent *tree.Tree) {
	parent.AddChild(tree.NodeString(il.Tok.Literal))
}
