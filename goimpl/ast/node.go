package ast

import "regexp"

type Node interface {
	String() string
	Dump() string // Dump returns a string representation of the node
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

// TrimDump removes all whitespace from the Dump output
func TrimDump(n Node) string {
	return TrimString(n.Dump())
}

// TrimString removes all whitespace from the string
func TrimString(s string) string {
	whitespaceRegex := regexp.MustCompile(`\s+`)
	return whitespaceRegex.ReplaceAllString(s, "")
}
