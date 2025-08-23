package ast

import (
	"goimpl/token"
	"strings"
)

// Function expression type fn(x, y) { ... }
type Fun struct {
	Tok  token.Token  // fn keyword token
	Args []Identifier // function arguments (x, y,...)
	Body Block        // {...} function body
}

// implement the Node interface
func (f Fun) TokenLiteral() string { return "fn" }

func (f Fun) String() string {
	args := make([]string, 0, len(f.Args))
	for _, arg := range f.Args {
		args = append(args, arg.Value)
	}
	return "fn(" + strings.Join(args, ", ") + ") " + f.Body.String()
}

// TODO: implement Dump or remove and rely on String() only?
func (f Fun) Dump(i int) string { return "Fn Dump not implemented!" }

// implement the expressionNode interface
func (Fun) expressionNode() {}
