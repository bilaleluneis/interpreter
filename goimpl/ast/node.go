package ast

import "fmt"

type Node interface {
	fmt.Stringer // must implement the String() method
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
