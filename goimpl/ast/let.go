package ast

import (
	"goimpl/debug/treedrawer/tree"
	"goimpl/token"
)

type Let struct {
	Tok   token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (Let) statmentNode() {}

func (l Let) TokenLiteral() string { return l.Tok.Literal }

func (l Let) String() string {
	out := l.Tok.Literal + " " + l.Name.String() + " = "
	if l.Value != nil {
		out += l.Value.String()
	}
	return out + ";"
}

func visualizeLet(l Let, parent *tree.Tree) {
	let := parent.AddChild(tree.NodeString(l.Tok.Literal))
	visualizeIdentifier(*l.Name, let)
	visualizeExpression(l.Value, let)
}
