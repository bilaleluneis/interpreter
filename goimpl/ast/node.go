package ast

type Node interface {
	String() string
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() //marker interface
}

type Expression interface {
	Node
	expressionNode() //marker interface
}
