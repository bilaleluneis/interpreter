package parser

import (
	"fmt"
	"goimpl/ast"
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

	//FIXME: Implement the rest of the parsing logic
	//for now just skip until we encounter a semicolon
	//ignoring what comes after the equal sign
	for p.currTok.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatment() ast.Statement {
	stmt := &ast.Return{Tok: p.currTok}
	//FIXME: Implement the rest of the parsing logic, skipping expression at the moment
	for p.nextToken(); p.currTok.Type != token.SEMICOLON; p.nextToken() {
	}
	return stmt
}

func (p *Parser) invalidStatment() ast.Statement {
	p.errors = append(p.errors, fmt.Sprintf("invalid token %s", p.currTok.Type))
	return nil
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{Tok: p.currTok}
	stmt.Exprssn = p.parseExpression(LOWEST)
	if stmt.Exprssn == nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse expression statement"))
		return nil
	}
	if p.peekTok.Type == token.SEMICOLON {
		p.nextToken()
		return stmt
	}
	return nil
}
