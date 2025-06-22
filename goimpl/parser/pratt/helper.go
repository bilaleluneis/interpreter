package pratt

import (
	"goimpl/parser/internal"
)

func (p *Parser) advance() {
	p.currTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
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
