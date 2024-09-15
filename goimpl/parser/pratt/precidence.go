package pratt

import "goimpl/parser/internal"

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
