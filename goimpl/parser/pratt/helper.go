package pratt

import (
	"goimpl/parser"
	"goimpl/token"
)

func (p *PrattParser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *PrattParser) registerPrefix(tokenType token.TokenType, fn parser.PrefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *PrattParser) registerInfix(tokenType token.TokenType, fn parser.InfixParseFn) {
	p.infixParseFns[tokenType] = fn
}
