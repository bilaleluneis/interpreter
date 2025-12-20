package ast

import (
	"goimpl/token"
	"strings"
)

type Block struct {
	Tok        token.Token // { token.LBRACE
	Statements []Statement
}

func (Block) statementNode()         {}
func (b Block) TokenLiteral() string { return b.Tok.Literal }
func (b Block) String() string {
	var out strings.Builder
	out.WriteString("{")
	if len(b.Statements) > 0 {
		out.WriteString("\n")
	}
	for _, s := range b.Statements {
		out.WriteString(s.String() + "\n")
	}
	return out.String() + "}"
}
