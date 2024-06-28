package parser

import (
	"fmt"
	"goimpl/ast"
	"goimpl/token"
)

var parseLetStatment parseStatement = func(p *Parser) (ast.Statement, error) {
	stmt := &ast.Let{Tok: p.currTok}
	if p.peekTok.Type != token.IDENTIFIER {
		return nil, fmt.Errorf("expected IDENTIFIER, got %s", p.peekTok.Type)
	}
	p.nextToken()
	stmt.Name = &ast.Identifier{Tok: p.currTok, Value: p.currTok.Literal}
	if p.peekTok.Type != token.ASSIGN {
		return nil, fmt.Errorf("expected ASSIGN, got %s", p.peekTok.Type)
	}

	//FIXME: Implement the rest of the parsing logic
	//for now just skip until we encounter a semicolon
	//ignoring what comes after the equal sign
	for p.currTok.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt, nil
}
