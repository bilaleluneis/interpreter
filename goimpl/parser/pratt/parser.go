package pratt

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

type Parser struct {
	lexer          lexer.Lexer
	currTok        token.Token
	peekTok        token.Token
	prefixParseFns map[token.TokenType]PrefixParseFn
	infixParseFns  map[token.TokenType]InfixParseFn
}

func New(l lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.initPrefixParseFns()
	p.initInfixParseFns()
	p.currTok = p.lexer.NextToken()
	p.peekTok = p.lexer.NextToken()
	return p
}

func (p *Parser) ParseProgram() ast.Program {
	var parsedStatements []ast.Statement
	for p.currTok.Type != token.EOF {
		stmt := p.parseStatement()
		parsedStatements = append(parsedStatements, stmt)
		switch stmt.(type) {
		case *ast.Error:
			return ast.Program{Statements: parsedStatements}
		default:
			p.advance()
		}
	}
	return ast.Program{Statements: parsedStatements}
}
