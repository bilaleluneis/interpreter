package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser"
	"goimpl/token"
	"strconv"
)

func (p *Parser) parseExpression(precedence precidence) ast.Expression {
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

var parseInfixExpr parser.InfixParseFn = func(parser parser.ParserType, left ast.Expression) ast.Expression {
	if prttParser, ok := parser.(*Parser); ok {
		expr := &ast.InfixExpression{Tok: prttParser.currTok, Operator: prttParser.currTok.Literal, Left: left}
		precedence := prttParser.currPrecidence()
		prttParser.nextToken()
		expr.Right = prttParser.parseExpression(precedence)
		return expr
	}
	return nil
}

var parseIdentifierExpr parser.PrefixParseFn = func(parser parser.ParserType) ast.Expression {
	if prttParser, ok := parser.(*Parser); ok {
		return &ast.Identifier{Tok: prttParser.currTok, Value: prttParser.currTok.Literal}
	}
	return nil
}

var parsePrefixExpr parser.PrefixParseFn = func(parser parser.ParserType) ast.Expression {
	if prttParser, ok := parser.(*Parser); ok {
		expr := &ast.PrefixExpression{Tok: prttParser.currTok, Operator: prttParser.currTok.Literal} //operator ex: ! or -
		prttParser.nextToken()
		expr.Right = prttParser.parseExpression(PREFIX) //parse the right side of the operator
		return expr
	}
	return nil
}

var parseIntegerLiteral parser.PrefixParseFn = func(parser parser.ParserType) ast.Expression {
	if prttParser, ok := parser.(*Parser); ok {
		literal := &ast.IntegerLiteral{Tok: prttParser.currTok}
		if value, err := strconv.ParseInt(prttParser.currTok.Literal, 0, 64); err == nil {
			literal.Value = value
			return literal
		}
		err := "could not parse " + prttParser.currTok.Literal + " as integer"
		prttParser.errors = append(prttParser.errors, err)
		return nil
	}
	return nil //FIXME: add error
}
