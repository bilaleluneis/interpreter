package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
)

type Parser[L any, CL lexer.CopyableLexer[L]] func(L) (ast.Statement, L)

func (p *Parser[L, CL]) ParseProgram() *ast.Program {
	return &ast.Program{Statements: []ast.Statement{}}
}

type ParseResult[L any, CL lexer.CopyableLexer[L]] struct {
	Lxr   L
	Stmnt ast.Statement
}

func (pr ParseResult[L, CL]) Parse(cparser ...Parser[L, CL]) []ParseResult[L, CL] {
	var results []ParseResult[L, CL]
	for _, parser := range cparser {
		cl := CL(&pr.Lxr).GetCopy()
		stmnt, lxrResult := parser(cl)
		results = append(results, ParseResult[L, CL]{Lxr: lxrResult, Stmnt: stmnt})
	}
	return results
}

func NewParseResult[L any, CL lexer.CopyableLexer[L]](l L) ParseResult[L, CL] {
	return ParseResult[L, CL]{Lxr: l}
}

func Filtr[L any, CL lexer.CopyableLexer[L]](p []ParseResult[L, CL], f func(ast.Statement) bool) []ParseResult[L, CL] {
	var filtered []ParseResult[L, CL]
	for _, p := range p {
		if f(p.Stmnt) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
