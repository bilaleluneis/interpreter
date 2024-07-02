package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

type parseStatement func(*Parser) ast.Statement
type parseExpression func(*Parser, precidence) ast.Expression
type prefixParseFn func(*Parser) ast.Expression
type infixParseFn func(*Parser, ast.Expression) ast.Expression

type Parser struct {
	lexer          lexer.LexerType
	currTok        token.Token
	peekTok        token.Token
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l lexer.LexerType) *Parser {
	p := &Parser{
		lexer:          l,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}
	register(p)
	p.currTok = p.lexer.NextToken()
	p.peekTok = p.lexer.NextToken()
	return p
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}
	for p.currTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	var parseFn = parseInvalid

	switch p.currTok.Type {
	case token.LET:
		parseFn = parseLetStatment
	case token.RETURN:
		parseFn = parseReturnStatment
	default:
		parseFn = parseExpressionStatement
	}

	return parseFn(p)
}

func register(p *Parser) {
	p.registerPrefix(token.IDENTIFIER, parseIdentifier)
	p.registerPrefix(token.INT, parseInteger)
	p.registerPrefix(token.BANG, parsePrefix)
	p.registerPrefix(token.MINUS, parsePrefix)
}
