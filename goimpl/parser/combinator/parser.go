package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

type CombinatorParser[T any, CL lexer.CopyableLexer[T]] func(T) (ast.Statement, T)

func (p *CombinatorParser[T, CL]) ParseProgram() *ast.Program {
	return &ast.Program{Statements: []ast.Statement{}}
}

type ParseResult[T any, CL lexer.CopyableLexer[T]] struct {
	Lxr   T
	Stmnt ast.Statement
}

func (pr ParseResult[T, CL]) Parse(cparser ...CombinatorParser[T, CL]) []ParseResult[T, CL] {
	var results []ParseResult[T, CL]
	for _, parser := range cparser {
		cl := CL(&pr.Lxr).GetCopy()
		stmnt, lxrResult := parser(cl)
		results = append(results, ParseResult[T, CL]{Lxr: lxrResult, Stmnt: stmnt})
	}
	return results
}

func NewParseResult[T any, CL lexer.CopyableLexer[T]](l T) ParseResult[T, CL] {
	return ParseResult[T, CL]{Lxr: l}
}
