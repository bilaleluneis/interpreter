package pratt

import (
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
	"strconv"
)

// parseExpression implements the Pratt parsing algorithm for expressions with operator precedence.
func (p *Parser) parseExpression(precedence internal.Precidence) ast.Expression {
	prefix := p.prefixParseFns[p.currTok.Type]
	if prefix == nil {
		// FIXME: need to return as.Error
		//fmt.Sprintf("no prefix parse function for %s found", p.currTok.Type)
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

// Parser functions for different expression types.
var (
	// parseInfixExpr handles operators like +, -, *, /, ==, etc.
	parseInfixExpr InfixParseFn = func(parser *Parser, left ast.Expression) ast.Expression {
		expr := &ast.InfixExpression{
			Tok:      parser.currTok,
			Operator: parser.currTok.Literal,
			Left:     left,
		}
		precedence := parser.currPrecidence()
		parser.nextToken()
		expr.Right = parser.parseExpression(precedence)
		return expr
	}

	// parseIdentifierExpr handles variable names and other identifiers.
	parseIdentifierExpr PrefixParseFn = func(parser *Parser) ast.Expression {
		return &ast.Identifier{
			Tok:   parser.currTok,
			Value: parser.currTok.Literal,
		}
	}

	// parsePrefixExpr handles unary operators like ! and - .
	parsePrefixExpr PrefixParseFn = func(parser *Parser) ast.Expression {
		expr := &ast.PrefixExpression{
			Tok:      parser.currTok,
			Operator: parser.currTok.Literal,
		}
		parser.nextToken()
		expr.Right = parser.parseExpression(internal.PREFIX)
		return expr
	}

	// parseIntegerLiteral handles numeric literals.
	parseIntegerLiteral PrefixParseFn = func(parser *Parser) ast.Expression {
		literal := &ast.IntegerLiteral{Tok: parser.currTok}
		value, err := strconv.ParseInt(parser.currTok.Literal, 0, 64)
		if err == nil {
			literal.Value = value
			return literal
		}
		//FIXME: need to return as.Error
		//"could not parse " + parser.currTok.Literal + " as integer"
		return nil
	}

	// parseBooleanExpr handles true/false literals.
	parseBooleanExpr PrefixParseFn = func(parser *Parser) ast.Expression {
		return &ast.Boolean{
			Tok:   parser.currTok,
			Value: parser.currTok.Type == token.TRUE,
		}
	}

	// parseGroupedExpression handles expressions in parentheses.
	parseGroupedExpression PrefixParseFn = func(parser *Parser) ast.Expression {
		parser.nextToken() // consume the '('

		expr := parser.parseExpression(internal.LOWEST)

		if parser.peekTok.Type != token.RPRAN {
			//FIXME: need to return as.Error
			// "expected )"
			return nil
		}

		parser.nextToken() // consume the ')'
		return expr
	}
)
