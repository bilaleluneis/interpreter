package ast

import "strings"

type Error struct {
	Message string
}

func (Error) statmentNode()          {}
func (e Error) TokenLiteral() string { return e.Message }
func (e Error) String() string       { return e.Message }

func (e Error) Dump(ident int) string {
	identation := strings.Repeat("\t", ident)
	out := identation + "ast.Error{ //start of Error\n"
	out += identation + "\tMessage: " + e.Message + "\n"
	out += identation + "} //end of Error"
	return out
}
