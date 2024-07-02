package parser

import (
	"fmt"
	"goimpl/ast"
	"goimpl/token"
)

var parseLetStatment parseStatement = func(p *Parser) ast.Statement {
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
