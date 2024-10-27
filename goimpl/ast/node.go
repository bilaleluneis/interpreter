package ast

type Node interface {
	String() string
	Dump(ident int) string // Dump returns a string representation of the node object
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
