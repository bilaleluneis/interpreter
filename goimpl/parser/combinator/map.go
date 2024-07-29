package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

func mapStatement[L lexer.LexerType, A ast.Statement](
	p CombinatorParser[L, A],
	mapF func(lexerType L) A) CombinatorParser[L, A] {
	return func(l L) (A, L) {
		_, l = p(l)
		//if s == nil {
		//	var astErr A
		//	return astErr, l
		//}
		return mapF(l), l
	}
}
