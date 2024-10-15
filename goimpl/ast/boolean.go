package ast

import "goimpl/token"

// Boolean represents a boolean expression in the AST.
type Boolean struct {
	Tok   token.Token // The token.BOOL token
	Value bool
}

func (Boolean) expressionNode()        {}
func (b Boolean) TokenLiteral() string { return b.Tok.Literal }
func (b Boolean) String() string       { return b.Tok.Literal }
func (b Boolean) Dump() string {
	return `ast.Boolean{
	Tok: ` + b.Tok.Literal + `,
	Value: ` + b.Tok.Literal + `
}`
}
