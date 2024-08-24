package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

// TODO: create constraint for PT type

type LPT[L any] interface {
	lexer.Lexer
	GetCopy() *L //FIXME: VType[L] interface returns L for GetCopy() instead of *L
	*L
}

type CombinatorParser[L any, PT interface {
	NextToken() token.Token
	GetCopy() *L
	*L
}] func(L) (ast.Statement, L)

func (p *CombinatorParser[L, PT]) ParseProgram() *ast.Program {
	return &ast.Program{Statements: []ast.Statement{}}
}

type ParseResult[L any, PT interface {
	NextToken() token.Token
	GetCopy() *L
	*L
}] struct {
	Lxr   L
	Stmnt ast.Statement
}

func (pr ParseResult[L, PT]) Parse(cparser ...CombinatorParser[L, PT]) []ParseResult[L, PT] {
	var results []ParseResult[L, PT]
	for _, parser := range cparser {
		cl := PT(&pr.Lxr).GetCopy()
		stmnt, lxrResult := parser(*cl)
		results = append(results, ParseResult[L, PT]{Lxr: lxrResult, Stmnt: stmnt})
	}
	return results
}

func NewParseResult[L any, PT interface {
	NextToken() token.Token
	GetCopy() *L
	*L
}](l L) ParseResult[L, PT] {
	return ParseResult[L, PT]{Lxr: l}
}
