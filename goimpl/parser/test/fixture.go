package test

import (
	"goimpl/lexer"
	"goimpl/parser"
	"goimpl/parser/combinator"
	"goimpl/parser/pratt"
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
		p.parsers["combinator"] = combinator.New(p.lexer, parserF...)
	}
	return p
}
