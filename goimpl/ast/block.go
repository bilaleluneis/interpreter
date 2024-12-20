package ast

import (
	"goimpl/token"
	"strings"
)

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

func (b Block) Dump(ident int) string {
	identation := strings.Repeat("\t", ident)
	out := identation + "ast.Block{ //start of Block\n"
	out += identation + "\tToken: token.Token{ Type: token.LBRACE, Literal: " + b.Token.Literal + "},\n"
	out += identation + "\tStatements: []ast.Statement{\n"
	for _, s := range b.Statements {
		out += s.Dump(ident + 1)
		out += ",\n"
	}
	out += identation + "}\n"
	out += identation + "} //end of Block"
	return out
}
