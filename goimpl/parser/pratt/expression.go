package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
	"strconv"
)

type PrefixParseFn func(*Parser) ast.Expression
type InfixParseFn func(*Parser, ast.Expression) ast.Expression

func (p *Parser) initPrefixParseFns() {
	p.prefixParseFns = map[token.TokenType]PrefixParseFn{
		token.IDENTIFIER: parseIdentifierExpr,
		token.INT:        parseIntegerLiteral,
		token.BANG:       parsePrefixExpr,
		token.MINUS:      parsePrefixExpr,
		token.PLUS:       parsePrefixExpr,
		token.TRUE:       parseBooleanExpr,
		token.FALSE:      parseBooleanExpr,
		token.LPRAN:      parseGroupedExpression,
		token.FUNCTION:   parseFunctionLiteral,
		token.IF:         parseIfExpression,
	}
}

func (p *Parser) initInfixParseFns() {
	p.infixParseFns = map[token.TokenType]InfixParseFn{
		token.PLUS:  parseInfixExpr,
		token.MINUS: parseInfixExpr,
		token.SLASH: parseInfixExpr,
		token.ASTER: parseInfixExpr,
		token.EQ:    parseInfixExpr,
		token.NEQ:   parseInfixExpr,
		token.LT:    parseInfixExpr,
		token.GT:    parseInfixExpr,
		token.LPRAN: parseCallExpression,
	}
}

// parseExpression implements the Pratt parsing for expressions with operator precedence.
func (p *Parser) parseExpression(precedence internal.Precidence) ast.Expression {
	prefix := p.prefixParseFns[p.currTok.Type]
	if prefix == nil {
		return &ast.InvalidExpression{
			Message: fmt.Sprintf(internal.ErrExpectedPrefixParseFn, p.currTok.Type),
		}
	}

	leftExpr := prefix(p)

	for p.peekTok.Type != token.SEMICOLON && precedence < p.peekPrecidence() {
		infix := p.infixParseFns[p.peekTok.Type]
		if infix == nil {
			return leftExpr
		}

		p.advance()
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
		parser.advance()
		expr.Right = parser.parseExpression(precedence)
		return expr
	}

	// parseCallExpression handles function calls like add(1, 2)
	parseCallExpression InfixParseFn = func(parser *Parser, left ast.Expression) ast.Expression {
		call := &ast.Call{Tok: parser.currTok, Function: left}

		// consume '('
		parser.advance()

		// no arguments: ()
		if parser.currTok.Type == token.RPRAN {
			return call
		}

		// at least one argument
		call.Args = append(call.Args, parser.parseExpression(internal.LOWEST))

		// additional comma-separated args
		for parser.peekTok.Type == token.COMMA {
			parser.advance() // consume current token (last arg)
			parser.advance() // move to next token (start of next arg)
			call.Args = append(call.Args, parser.parseExpression(internal.LOWEST))
		}

		// expect closing ')'
		if parser.peekTok.Type != token.RPRAN {
			return &ast.InvalidExpression{
				Message: fmt.Sprintf(internal.ErrExpectedClosePren, parser.peekTok.Type),
			}
		}
		parser.advance() // consume ')'
		return call
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
		parser.advance()
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
		return &ast.InvalidExpression{
			Message: fmt.Sprintf(internal.ErrExpectedIntegerLiteral, parser.currTok.Literal),
		}
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
		parser.advance() // consume the '('

		expr := parser.parseExpression(internal.LOWEST)

		if parser.peekTok.Type != token.RPRAN {
			return &ast.InvalidExpression{
				Message: fmt.Sprintf(internal.ErrExpectedOpenPren, parser.peekTok.Type),
			}
		}

		parser.advance() // consume the ')'
		return expr
	}

	// parse a function : ex fn(x, y) { return (x + y); }
	// the block doesn't end up with a semicolon
	parseFunctionLiteral PrefixParseFn = func(parser *Parser) ast.Expression {
		fun := &ast.Fun{Tok: parser.currTok}
		if parser.peekTok.Type != token.LPRAN {
			return &ast.InvalidExpression{
				Message: fmt.Sprintf(internal.ErrExpectedOpenPren, parser.peekTok.Type),
			}
		}
		parser.advance() // consume 'fn'
		parser.advance() // consume '('
		for parser.currTok.Type == token.IDENTIFIER {
			fun.Args = append(fun.Args, ast.Identifier{
				Tok:   parser.currTok,
				Value: parser.currTok.Literal,
			})
			parser.advance() // consume identifier
			if parser.currTok.Type == token.COMMA {
				parser.advance() // consume ','
			} else {
				break
			}
		}
		if parser.currTok.Type != token.RPRAN {
			return &ast.InvalidExpression{
				Message: fmt.Sprintf(internal.ErrExpectedClosePren, parser.currTok.Type),
			}
		}
		parser.advance() // consume ')'
		if parser.currTok.Type != token.LBRACE {
			return &ast.InvalidExpression{
				Message: fmt.Sprintf(internal.ErrExpectedOpenBrace, parser.currTok.Type),
			}
		}
		blockStmt := parser.parseBlockStatement()
		switch funBody := blockStmt.(type) {
		case *ast.Error:
			return &ast.InvalidExpression{
				Message: funBody.Message,
			}
		case *ast.Block:
			fun.Body = *funBody
		}
		return fun
	}
)
