package ast

import (
	"fmt"
	"goimpl/token"
	"strings"
)

// Boolean represents a boolean expression in the AST.
type Boolean struct {
	Tok   token.Token // The token.BOOL token
	Value bool
}

func (Boolean) expressionNode()        {}
func (b Boolean) TokenLiteral() string { return b.Tok.Literal }
func (b Boolean) String() string       { return b.Tok.Literal }
func (b Boolean) Dump(ident int) string {
	identation := strings.Repeat("\t", ident)
	out := identation + "ast.Boolean{ //start of Boolean\n"
	out += identation + "\tTok: token.Token{ Type: token.BOOL, Literal: " + b.Tok.Literal + "},\n"
	out += identation + "\tValue: " + fmt.Sprintf("%t", b.Value) + "\n"
	out += identation + "} //end of Boolean\n"
	return out
}
