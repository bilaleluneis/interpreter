package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

func Fail[L lexer.LexerConstraint[L]](l L) Result[L] {
	return Result[L]{l, ast.Error{Message: "Failed to Parse"}}
}
