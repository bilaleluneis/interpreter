package parser

import (
	"goimpl/ast"
	"goimpl/token"
)

var parseReturnStatment parseStatement = func(p *Parser) (ast.Statement, error) {
	stmt := &ast.Return{Tok: p.currTok}
	//FIXME: Implement the rest of the parsing logic, skipping expression at the moment
	for p.nextToken(); p.currTok.Type != token.SEMICOLON; p.nextToken() {
	}
	return stmt, nil
}
