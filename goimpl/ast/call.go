package ast

import (
	"goimpl/token"
	"strings"
)

// Call Call: represents a function call expression.
// <expression>(<comma separated expression list>)
// Examples:
//
//	add(1, 2 * 3, 4 + 5)
//	foobar()
//	fn (x, y) { x + y }(10, 20)
//	fn(x, y, fn(a, b) { a + b })(1, 2, 3, 4)
type Call struct {
	Tok      token.Token  // The '(' token
	Function Expression   // The function being called
	Args     []Expression // The arguments to the function
}

func (c Call) expressionNode() {}

func (c Call) TokenLiteral() string { return c.Tok.Literal }

func (c Call) String() string {
	var args []string
	for _, a := range c.Args {
		args = append(args, a.String())
	}
	return c.Function.String() + "(" + joinExpressions(args) + ")"
}

func joinExpressions(e []string) string {
	var b strings.Builder
	for i, expr := range e {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(expr)
	}
	return b.String()
}
