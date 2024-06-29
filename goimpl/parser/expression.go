package parser

import (
	"fmt"
	"goimpl/ast"
	"goimpl/token"
)

type precidence int

const (
	_ precidence = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var parseExpressionStatement parseStatement = func(p *Parser) (ast.Statement, error) {
	stmt := &ast.ExpressionStatement{Tok: p.currTok}
	stmt.Exprssn, _ = parseExprWithPriority(p, LOWEST) //FIXME: handle error
	if p.peekTok.Type == token.SEMICOLON {
		p.nextToken()
	}
	return stmt, nil
}

func parseExprWithPriority(p *Parser, precedence precidence) (ast.Expression, error) {
	prefix := p.prefixParseFns[p.currTok.Type]
	if prefix == nil {
		return nil, fmt.Errorf("no prefix parse function for %s found", p.currTok.Type)
	}
	return prefix(p), nil
}

var parseIdentifier prefixParseFn = func(parser *Parser) ast.Expression {
	return &ast.Identifier{Tok: parser.currTok, Value: parser.currTok.Literal}
}
