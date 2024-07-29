package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

type CombinatorParser[L lexer.LexerType, A ast.Statement] func(L) (A, L)

func (p *CombinatorParser[L, A]) ParseProgram() *ast.Program {
	return &ast.Program{Statements: []ast.Statement{}}
}

// TODO: implement parse
func parse[L lexer.LexerType, A ast.Statement](c CombinatorParser[L, A], l L) {}
