package parser

import (
	"goimpl/token"
)

func (p *Parser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
