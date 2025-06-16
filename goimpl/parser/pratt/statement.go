package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
)

// Statement parsing methods for the Pratt parser
func (p *Parser) parseStatement() ast.Statement {
	switch p.currTok.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement parses a let statement in the form: let <identifier> = <expression>;
func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.Let{}

	if p.peekTok.Type != token.IDENTIFIER {
		return &ast.Error{Message: fmt.Sprintf(internal.LetErrExpectedIdentifier, p.peekTok.Type)}
	}

	p.nextToken()
	stmt.Name = &ast.Identifier{
		Tok:   p.currTok,
		Value: p.currTok.Literal,
	}

	if p.peekTok.Type != token.ASSIGN {
		return &ast.Error{Message: fmt.Sprintf(internal.LetErrExpectedAssign, p.peekTok.Type)}
	}

	p.nextToken() // consume the ASSIGN token
	p.nextToken() // consume the value token (expression)

	expr := p.parseExpression(internal.LOWEST)

	switch expr := expr.(type) {
	case *ast.InvalidExpression:
		return &ast.Error{
			Message: fmt.Sprintf(internal.LetErrExpectedExpression, p.currTok.Literal),
		}

	default:
		stmt.Value = expr
		for p.currTok.Type != token.SEMICOLON && p.currTok.Type != token.EOF {
			p.nextToken()
		}
		if p.currTok.Type == token.EOF {
			return &ast.Error{Message: fmt.Sprintf(internal.LetErrExpectedSemicolon, token.EOF)}
		}
		return stmt
	}
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.Return{Tok: p.currTok}

	p.nextToken() // consume 'return' token

	// Parse the return value expression
	stmt.Value = p.parseExpression(internal.LOWEST)

	// Advance until semicolon if present
	if p.peekTok.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	var stmt ast.Statement
	expr := p.parseExpression(internal.LOWEST)

	switch expr := expr.(type) {
	case *ast.InvalidExpression:
		return &ast.Error{Message: expr.Message}

	default: // assume it's a valid expression
		stmt = &ast.ExpressionStatement{Tok: p.currTok, Exprssn: expr}
	}

	if p.peekTok.Type == token.SEMICOLON {
		p.nextToken()
		return stmt
	}

	return &ast.Error{
		Message: fmt.Sprintf(internal.LetErrExpectedSemicolon, p.peekTok.Type),
	}

}
