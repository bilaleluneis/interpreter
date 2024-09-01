package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

func Fail[L any, _ lexer.CopyableLexer[L]](l L) (ast.Statement, L) {
	return ast.Error{Message: "failed to lex"}, l
}
