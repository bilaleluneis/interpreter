package parser

import (
	"fmt"
	"goimpl/ast"
	"goimpl/token"
)

var parseExpressionStatement parseStatement = func(p *Parser) ast.Statement {
	stmt := &ast.ExpressionStatement{Tok: p.currTok}
	stmt.Exprssn = parseExpr(p, LOWEST)
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

var parseIdentifier prefixParseFn = func(parser *Parser) ast.Expression {
	return &ast.Identifier{Tok: parser.currTok, Value: parser.currTok.Literal}
}

var parsePrefix prefixParseFn = func(parser *Parser) ast.Expression {
	expr := &ast.PrefixExpression{Tok: parser.currTok, Operator: parser.currTok.Literal} //operator ex: ! or -
	parser.nextToken()
	expr.Right = parseExpr(parser, PREFIX) //parse the right side of the operator
	return expr
}

var parseExpr parseExpression = func(p *Parser, precedence precidence) ast.Expression {
	prefix := p.prefixParseFns[p.currTok.Type]
	if prefix == nil {
		p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", p.currTok.Type))
		return nil
	}
	return prefix(p)
}
