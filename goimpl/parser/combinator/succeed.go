package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

func succeed[L lexer.LexerType, A ast.Statement](s A) CombinatorParser[L, A] {
	return func(l L) (A, L) {
		return s, l
	}
}
