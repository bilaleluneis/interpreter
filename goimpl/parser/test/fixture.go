package test

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser"
	"goimpl/parser/combinator"
	"goimpl/parser/pratt"
	"testing"
)

type parsersUnderTest[L lexer.CopyableLexer[L]] struct {
	lexer   L
	parsers map[string]parser.ParserType
}

func testParsers[L lexer.CopyableLexer[L]](lxr L) *parsersUnderTest[L] {
	return &parsersUnderTest[L]{
		lexer:   lxr,
		parsers: make(map[string]parser.ParserType),
	}
}

func (p *parsersUnderTest[L]) initPratt() *parsersUnderTest[L] {
	l := p.lexer.GetCopy()
	p.parsers["pratt"] = pratt.New(l)
	return p
}

func (p *parsersUnderTest[L]) initCombinator(parserF ...combinator.ParserFunc[L]) *parsersUnderTest[L] {
	if len(parserF) != 0 {
		p.parsers["combinator"] = combinator.New(p.lexer.GetCopy(), parserF...)
	}
	return p
}

type lexerToAstResult struct {
	lxr             lexer.StubLexer
	expectedAst     ast.Statement
	expectedProgram string
}

func (lta *lexerToAstResult) test(t *testing.T, p parser.ParserType) {
	program, ok := p.ParseProgram()
	if !ok {
		t.Logf("expected ok program got !ok")
		t.FailNow()
	}
	if program.Top() != lta.expectedProgram {
		t.Logf("expected %s got %s", lta.expectedProgram, program.Top())
		t.FailNow()
	}
	if program.Statements[0] != lta.expectedAst {
		t.Logf("expected %T got %T", lta.expectedAst, program.Statements[0])
		t.FailNow()
	}
}
