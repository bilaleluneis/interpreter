package ast

import (
	"goimpl/debug/treedrawer/tree"
	"goimpl/token"
)

type Return struct {
	Tok   token.Token // the token.RETURN token
	Value Expression
}

func (Return) statmentNode() {}

func (r Return) TokenLiteral() string { return r.Tok.Literal }

func (r Return) String() string {
	out := r.Tok.Literal + " "
	if r.Value != nil {
		out += r.Value.String()
	}
	return out + ";"
}

func visualizeReturn(r Return, parent *tree.Tree) {
	ret := parent.AddChild(tree.NodeString(r.Tok.Literal))
	visualizeExpression(r.Value, ret)
}
