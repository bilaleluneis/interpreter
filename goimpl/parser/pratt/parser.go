package pratt

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser"
	"goimpl/token"
)

type PrattParser struct {
	lexer          lexer.Lexer
	currTok        token.Token
	peekTok        token.Token
	errors         []string
	prefixParseFns map[token.TokenType]parser.PrefixParseFn
	infixParseFns  map[token.TokenType]parser.InfixParseFn
}

func New(l lexer.Lexer) *PrattParser {
	p := &PrattParser{
		lexer:          l,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]parser.PrefixParseFn),
		infixParseFns:  make(map[token.TokenType]parser.InfixParseFn),
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

func (p *PrattParser) Errors() []string {
	return p.errors
}

func (p *PrattParser) ParseProgram() *ast.Program {
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
