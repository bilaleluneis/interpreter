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
	stmt := &ast.Let{Tok: p.currTok}

	if p.peekTok.Type != token.IDENTIFIER {
		msg := fmt.Sprintf("expected IDENTIFIER, got %s", p.peekTok.Type)
		p.errors = append(p.errors, msg)
		return nil
	}

	p.nextToken()
	stmt.Name = &ast.Identifier{
		Tok:   p.currTok,
		Value: p.currTok.Literal,
	}

	if p.peekTok.Type != token.ASSIGN {
		msg := fmt.Sprintf("expected ASSIGN, got %s", p.peekTok.Type)
		p.errors = append(p.errors, msg)
		return nil
	}

	p.nextToken() // consume the ASSIGN token
	p.nextToken() // consume the value token (expression)

	if expr := p.parseExpression(internal.LOWEST); expr != nil {
		stmt.Value = expr
	} else {
		return nil
	}

	// FIXME: look at bellow , I need to return ast.Error instead when parsing fails

	// Advance until we find a semicolon or EOF
	for p.currTok.Type != token.SEMICOLON && p.currTok.Type != token.EOF {
		p.nextToken()
	}

	if p.currTok.Type == token.EOF {
		msg := fmt.Sprintf("expected semicolon, got EOF")
		p.errors = append(p.errors, msg)
		return nil
	}

	return stmt
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
	if expression := p.parseExpression(internal.LOWEST); expression != nil {
		stmt := &ast.ExpressionStatement{Tok: p.currTok, Exprssn: expression}

		if p.peekTok.Type == token.SEMICOLON {
			p.nextToken()
			return stmt
		}
	}

	p.errors = append(p.errors, "could not parse expression statement")
	return nil
}
