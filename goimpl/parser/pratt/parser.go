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
		errors:         make([]string, 0),
		prefixParseFns: make(map[token.TokenType]PrefixParseFn),
		infixParseFns:  make(map[token.TokenType]InfixParseFn),
	}

	// Register prefix parse functions
	prefixFns := map[token.TokenType]PrefixParseFn{
		token.IDENTIFIER: parseIdentifierExpr,
		token.INT:        parseIntegerLiteral,
		token.BANG:       parsePrefixExpr,
		token.MINUS:      parsePrefixExpr,
		token.TRUE:       parseBooleanExpr,
		token.FALSE:      parseBooleanExpr,
		token.LPRAN:      parseGroupedExpression,
	}
	for tok, fn := range prefixFns {
		p.registerPrefix(tok, fn)
	}

	// Register infix parse functions
	infixFns := map[token.TokenType]InfixParseFn{
		token.PLUS:  parseInfixExpr,
		token.MINUS: parseInfixExpr,
		token.SLASH: parseInfixExpr,
		token.ASTER: parseInfixExpr,
		token.EQ:    parseInfixExpr,
		token.NEQ:   parseInfixExpr,
		token.LT:    parseInfixExpr,
		token.GT:    parseInfixExpr,
	}
	for tok, fn := range infixFns {
		p.registerInfix(tok, fn)
	}

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
