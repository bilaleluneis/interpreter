package pratt

import (
	"goimpl/parser/internal"
	"goimpl/token"
)

func (p *Parser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn PrefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn InfixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekPrecidence() internal.Precidence {
	if p, ok := internal.PrecidenceMap[p.peekTok.Type]; ok {
		return p
	}
	return internal.LOWEST
}

func (p *Parser) currPrecidence() internal.Precidence {
	if p, ok := internal.PrecidenceMap[p.currTok.Type]; ok {
		return p
	}
	return internal.LOWEST
}
