package test

import (
	"goimpl/lexer"
	"goimpl/parser"
	"goimpl/parser/combinator"
	"goimpl/parser/pratt"
)

type lxrTyp = lexer.StubLexer // alias for lexer.StubLexer

type parsersUnderTest struct {
	lexer   lxrTyp
	parsers map[string]parser.ParserType
}

func testParsers(lxr lxrTyp) *parsersUnderTest {
	return &parsersUnderTest{
		lexer:   lxr,
		parsers: make(map[string]parser.ParserType),
	}
}

func (p *parsersUnderTest) initPratt() *parsersUnderTest {
	l := p.lexer.GetCopy()
	p.parsers["pratt"] = pratt.New(&l)
	return p
}

func (p *parsersUnderTest) initCombinator(parserF ...combinator.ParserFunc[lxrTyp, *lxrTyp]) *parsersUnderTest {
	p.parsers["combinator"] = combinator.New(p.lexer.GetCopy(), parserF...)
	return p
}
