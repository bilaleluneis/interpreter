package ast

type Node interface {
	String() string
	TokenLiteral() string
}

type Statement interface {
	Node
	statmentNode() //marker interface
}

type Expression interface {
	Node
	expressionNode() //marker interface
}
