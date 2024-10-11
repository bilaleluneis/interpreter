package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
	"strconv"
)

func (p *Parser) parseExpression(precedence internal.Precidence) ast.Expression {
	prefix := p.prefixParseFns[p.currTok.Type]
	if prefix == nil {
		p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", p.currTok.Type))
		return nil
	}
	leftExpr := prefix(p)
	for p.peekTok.Type != token.SEMICOLON && precedence < p.peekPrecidence() {
		infix := p.infixParseFns[p.peekTok.Type]
		if infix == nil {
			return leftExpr
		}
		p.nextToken()
		leftExpr = infix(p, leftExpr)
	}
	return leftExpr
}

var parseInfixExpr InfixParseFn = func(parser *Parser, left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{Tok: parser.currTok, Operator: parser.currTok.Literal, Left: left}
	precedence := parser.currPrecidence()
	parser.nextToken()
	expr.Right = parser.parseExpression(precedence)
	return expr
}

var parseIdentifierExpr PrefixParseFn = func(parser *Parser) ast.Expression {
	return &ast.Identifier{Tok: parser.currTok, Value: parser.currTok.Literal}
}

var parsePrefixExpr PrefixParseFn = func(parser *Parser) ast.Expression {
	expr := &ast.PrefixExpression{Tok: parser.currTok, Operator: parser.currTok.Literal} //operator ex: ! or -
	parser.nextToken()
	expr.Right = parser.parseExpression(internal.PREFIX) //parse the right side of the operator
	return expr
}

var parseIntegerLiteral PrefixParseFn = func(parser *Parser) ast.Expression {
	literal := &ast.IntegerLiteral{Tok: parser.currTok}
	if value, err := strconv.ParseInt(parser.currTok.Literal, 0, 64); err == nil {
		literal.Value = value
		return literal
	}
	err := "could not parse " + parser.currTok.Literal + " as integer"
	parser.errors = append(parser.errors, err)
	return nil
}

var parseBooleanExpr PrefixParseFn = func(parser *Parser) ast.Expression {
	return &ast.Boolean{Tok: parser.currTok, Value: parser.currTok.Type == token.TRUE}
}
