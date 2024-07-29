package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

var Fail CombinatorParser[lexer.LexerType, ast.Error] = func(l lexer.LexerType) (ast.Error, lexer.LexerType) {
	var astErr ast.Error
	return astErr, l
}
