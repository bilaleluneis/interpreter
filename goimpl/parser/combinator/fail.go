package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

func Fail[L any, CL lexer.CopyableLexer[L]](l L) Result[L, CL] {
	return Result[L, CL]{l, ast.Error{Message: "Failed to Parse"}}
}
