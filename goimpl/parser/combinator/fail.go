package combinator

import (
	"goimpl/ast"
	"goimpl/token"
)

func Fail[L any, PT interface {
	NextToken() token.Token
	GetCopy() L
	*L
}](l L) (ast.Statement, L) {
	return ast.Error{Message: "failed to lex"}, l
}
