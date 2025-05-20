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
	errStmt := &ast.Error{}
	stmt := &ast.Let{Tok: p.currTok}

	if p.peekTok.Type != token.IDENTIFIER {
		errStmt.Message = fmt.Sprintf(internal.LetErrExpectedIdentifier, p.peekTok.Type)
		return errStmt
	}

	p.nextToken()
	stmt.Name = &ast.Identifier{
		Tok:   p.currTok,
		Value: p.currTok.Literal,
	}

	if p.peekTok.Type != token.ASSIGN {
		errStmt.Message = fmt.Sprintf(internal.LetErrExpectedAssign, p.peekTok.Type)
		return errStmt
	}

	p.nextToken() // consume the ASSIGN token
	p.nextToken() // consume the value token (expression)

	if expr := p.parseExpression(internal.LOWEST); expr != nil {
		stmt.Value = expr
	} else {
		errStmt.Message = fmt.Sprintf(internal.LetErrExpectedExpression, p.currTok.Type)
		return errStmt
	}

	// Advance until we find a semicolon or EOF
	for p.currTok.Type != token.SEMICOLON && p.currTok.Type != token.EOF {
		p.nextToken()
	}

	if p.currTok.Type == token.EOF {
		errStmt.Message = fmt.Sprintf(internal.LetErrExpectedSemicolon, token.EOF)
		return errStmt
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

	return &ast.Error{
		//FIXME: add error message
		// Message: fmt.Sprintf(internal.ExprErrExpectedSemicolon, p.peekTok.Type),
	}

}
