package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currTok.Type {
	case token.LET:
		return p.parseLetStatment()
	case token.RETURN:
		return p.parseReturnStatment()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatment() ast.Statement {
	stmt := &ast.Let{Tok: p.currTok}

	if p.peekTok.Type != token.IDENTIFIER {
		p.errors = append(p.errors, fmt.Sprintf("expected IDENTIFIER, got %s", p.peekTok.Type))
		return nil
	}

	p.nextToken()
	stmt.Name = &ast.Identifier{Tok: p.currTok, Value: p.currTok.Literal}
	if p.peekTok.Type != token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("expected ASSIGN, got %s", p.peekTok.Type))
		return nil
	}

	p.nextToken() // consume the ASSIGN token
	p.nextToken() // consume the value token (expression)
	if expr := p.parseExpression(internal.LOWEST); expr != nil {
		stmt.Value = expr
	} else {
		return nil
	}

	for p.currTok.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// FIXME: at the moment value is not captured, just skipping until we find a semicolon
func (p *Parser) parseReturnStatment() ast.Statement {
	stmt := &ast.Return{Tok: p.currTok}
	//FIXME: Implement the rest of the parsing logic, skipping expression at the moment
	for p.nextToken(); p.currTok.Type != token.SEMICOLON; p.nextToken() {
	}
	return stmt
}

func (p *Parser) invalidStatment() ast.Statement {
	return ast.Error{Message: fmt.Sprintf("invalid token %s", p.currTok.Type)}
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
