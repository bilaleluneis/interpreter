package ast

import "goimpl/token"

type Block struct {
	token.Token // { token.LBRACE
	Statements  []Statement
}

func (Block) statmentNode()          {}
func (b Block) TokenLiteral() string { return b.Token.Literal }
func (b Block) String() string {
	out := "{"
	for _, s := range b.Statements {
		out += s.String()
	}
	return out + "}"
}

func (b Block) Dump() string {
	out := `ast.Block{
	Tok: ` + b.Token.Literal + `,
	Statements: []ast.Statement{`
	for _, s := range b.Statements {
		out += s.Dump() + ","
	}
	out += `}
}`
	return out
}
