package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

type Result[L any, CL lexer.CopyableLexer[L]] struct {
	lxr   L
	stmnt ast.Statement
}

type ParserFunc[L any, CL lexer.CopyableLexer[L]] func(L) Result[L, CL]

type resultList[L any, CL lexer.CopyableLexer[L]] []Result[L, CL]

func (resList resultList[L, CL]) drop(f func(ast.Statement) bool) resultList[L, CL] {
	var filteredList resultList[L, CL]
	for _, result := range resList {
		if !f(result.stmnt) {
			filteredList = append(filteredList, result)
		}
	}
	return filteredList
}

func (resList resultList[L, CL]) first() (Result[L, CL], bool) {
	if len(resList) == 0 {
		return Result[L, CL]{}, false
	}
	return resList[0], true
}

type Parser[L any, CL lexer.CopyableLexer[L]] struct {
	currLexer L
	parsers   []ParserFunc[L, CL]
}

func (p Parser[L, CL]) ParseProgram() *ast.Program {
	var parsedStatements []ast.Statement
	lxr := p.currLexer
	parseNext := true
	isError := func(stmnt ast.Statement) bool { _, isError := stmnt.(ast.Error); return isError }
	for parseNext {
		// if we have reached EOF then stop parsing
		if p.isEof(lxr) {
			parseNext = false
			continue
		}
		if result, ok := p.parse(lxr).drop(isError).first(); ok {
			lxr = result.lxr //update to use lexer associated success parsing, that could have advanced
			parsedStatements = append(parsedStatements, result.stmnt)
		} else {
			parseNext = false
			parsedStatements = append(parsedStatements, ast.Error{Message: "No valid statement parsed"})
		}
	}
	return &ast.Program{Statements: parsedStatements}
}

func (p Parser[L, CL]) parse(currLxr L) resultList[L, CL] {
	var resultList resultList[L, CL]
	for _, parser := range p.parsers {
		lxr := CL(&currLxr).GetCopy() // make copy of lexer
		result := parser(lxr)
		resultList = append(resultList, result)
	}
	return resultList
}

func (p Parser[L, CL]) isEof(currLxr L) bool {
	cp := CL(&currLxr).GetCopy()
	return CL(&cp).NextToken().Type == token.EOF
}

func New[L any, CL lexer.CopyableLexer[L]](lexer L, parsers ...ParserFunc[L, CL]) Parser[L, CL] {
	return Parser[L, CL]{currLexer: lexer, parsers: parsers}
}
