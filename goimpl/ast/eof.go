package ast

import "strings"

type Eof struct{}

func (Eof) statmentNode()        {}
func (Eof) TokenLiteral() string { return "EOF" }
func (Eof) String() string       { return "EOF" }

func (Eof) Dump(ident int) string {
	identation := strings.Repeat("\t", ident)
	out := identation + "ast.Eof{ //start of Eof\n"
	out += identation + "} //end of Eof"
	return out
}
