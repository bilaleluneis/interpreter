package ast

import "goimpl/token"

type IntegerLiteral struct {
	Tok   token.Token // the token.INT token
	Value int64
}

func (IntegerLiteral) expressionNode() {}

func (il IntegerLiteral) TokenLiteral() string { return il.Tok.Literal }

func (il IntegerLiteral) String() string { return il.Tok.Literal }

func (il IntegerLiteral) Dump() string {
	return `ast.IntegerLiteral{
	Tok: ` + il.Tok.Literal + `,
	Value: ` + il.Tok.Literal + `
}`
}
