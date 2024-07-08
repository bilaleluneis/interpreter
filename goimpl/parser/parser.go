package parser

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

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

	// Register prefix parse functions
	p.registerPrefix(token.IDENTIFIER, parseIdentifierExpr)
	p.registerPrefix(token.INT, parseIntegerLiteral)
	p.registerPrefix(token.BANG, parsePrefixExpr)
	p.registerPrefix(token.MINUS, parsePrefixExpr)

	// Register infix parse functions
	p.registerInfix(token.PLUS, parseInfixExpr)

	p.currTok = p.lexer.NextToken()
	p.peekTok = p.lexer.NextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
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
