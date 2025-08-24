package ast

import (
	"goimpl/token"
)

type Block struct {
	Tok        token.Token // { token.LBRACE
	Statements []Statement
}

func (Block) statmentNode()          {}
func (b Block) TokenLiteral() string { return b.Tok.Literal }
func (b Block) String() string {
	out := "{"
	if len(b.Statements) > 0 {
		out += "\n"
	}
	for _, s := range b.Statements {
		out += s.String() + "\n"
	}
	return out + "}"
}
