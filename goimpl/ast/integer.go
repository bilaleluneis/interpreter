package ast

import (
	"fmt"
	"goimpl/token"
)

type IntegerLiteral struct {
	Tok   token.Token // the token.INT token
	Value int64
}

func (IntegerLiteral) expressionNode() {}

func (il IntegerLiteral) TokenLiteral() string { return il.Tok.Literal }

func (il IntegerLiteral) String() string { return fmt.Sprintf("%d", il.Value) }
