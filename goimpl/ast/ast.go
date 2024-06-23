package ast

import "goimpl/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statment() //marker interface
}

type Expression interface {
	Node
	expression() //marker interface
}

// Identifier can also be an Expression
type Identifier struct {
	Tok   token.Token // the token.IDENT token
	Value string
}

func (Identifier) expression() {}
func (i Identifier) TokenLiteral() string {
	return i.Tok.Literal
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
