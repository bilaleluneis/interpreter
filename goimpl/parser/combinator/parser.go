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

// ParseProgram will always create a new deep copy of lexer
// and then pass it to list of parsers, the result will be
// collected and parse failures filtered out, the first
// successful parse result will be appended to the parsed statments list
// and lexer associated with that parse will be used to parse next
func (p Parser[L, CL]) ParseProgram() (ast.Program, bool) {
	var parsedStatements []ast.Statement
	lxr := p.currLexer
	parseNext := true
	isError := func(stmnt ast.Statement) bool { _, isError := stmnt.(ast.Error); return isError }
	for parseNext {
		parseNext = false // we default to not parsting again, unless we have more to parse
		if p.isEof(lxr) { // if we have reached EOF then skip rest of loop
			continue
		}
		if result, ok := p.parse(lxr).drop(isError).first(); ok {
			lxr = result.lxr //update to use lexer associated success parsing, that could have advanced
			parsedStatements = append(parsedStatements, result.stmnt)
			parseNext = true // we might have more to parse, attempt again
		}
	}
	return ast.Program{Statements: parsedStatements}, len(parsedStatements) > 0
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
