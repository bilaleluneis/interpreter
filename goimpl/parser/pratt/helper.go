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

// peekPrecidence returns the precedence of the peek token.
func (p *Parser) peekPrecidence() internal.Precidence {
	if precidence, exists := internal.PrecidenceMap[p.peekTok.Type]; exists {
		return precidence
	}
	return internal.LOWEST
}

// currPrecidence returns the precedence of the current token
func (p *Parser) currPrecidence() internal.Precidence {
	if precidence, exists := internal.PrecidenceMap[p.currTok.Type]; exists {
		return precidence
	}
	return internal.LOWEST
}
