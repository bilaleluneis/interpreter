package pratt

import (
	"goimpl/parser"
	"goimpl/token"
)

func (p *Parser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn parser.PrefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn parser.InfixParseFn) {
	p.infixParseFns[tokenType] = fn
}
