package ast

import "goimpl/token"

// IfExpression represents an if-else expression in the AST.
// if <condition> { <consequence> } else { <alternative> }
type IfExpression struct {
	Tok       token.Token // the 'if' token
	Condition Expression
	Conseq    Block
	Alt       Block // optional else block, can be empty
}

func (ie IfExpression) expressionNode() {}

func (ie IfExpression) TokenLiteral() string { return ie.Tok.Literal }

func (ie IfExpression) String() string {
	out := "if"
	out += "("
	if ie.Condition == nil {
		out += "<nil>)"
	} else {
		out += ie.Condition.String() + ")"
	}
	out += ie.Conseq.String()
	if len(ie.Alt.Statements) > 0 {
		out += " else " + ie.Alt.String()
	}
	return out
}
