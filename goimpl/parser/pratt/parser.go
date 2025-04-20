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
	errors         []string
	prefixParseFns map[token.TokenType]PrefixParseFn
	infixParseFns  map[token.TokenType]InfixParseFn
}

type PrefixParseFn func(*Parser) ast.Expression
type InfixParseFn func(*Parser, ast.Expression) ast.Expression

func New(l lexer.Lexer) *Parser {
	p := &Parser{
		lexer:          l,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]PrefixParseFn),
		infixParseFns:  make(map[token.TokenType]InfixParseFn),
	}

	// Register prefix parse functions
	p.registerPrefix(token.IDENTIFIER, parseIdentifierExpr)
	p.registerPrefix(token.INT, parseIntegerLiteral)
	p.registerPrefix(token.BANG, parsePrefixExpr)
	p.registerPrefix(token.MINUS, parsePrefixExpr)
	p.registerPrefix(token.TRUE, parseBooleanExpr)
	p.registerPrefix(token.FALSE, parseBooleanExpr)

	// Register infix parse functions
	p.registerInfix(token.PLUS, parseInfixExpr)

	p.currTok = p.lexer.NextToken()
	p.peekTok = p.lexer.NextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() (ast.Program, bool) {
	var parsedStatements []ast.Statement
	for stmt := p.parseStatement(); len(p.errors) == 0; stmt = p.parseStatement() {
		parsedStatements = append(parsedStatements, stmt)

		p.nextToken()
		if p.peekTok.Type == token.EOF {
			break
		}
	}
	return ast.Program{Statements: parsedStatements}, len(p.errors) == 0
}
