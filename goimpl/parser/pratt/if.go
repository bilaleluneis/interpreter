package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
)

// parseIfExpression parses an if expression.
// if(<condition>) { <consequence> } else { <alternative> }
var parseIfExpression PrefixParseFn = func(parser *Parser) ast.Expression {
	ifExpr := &ast.IfExpression{
		Tok: token.Lookup("if"),
	}
	parser.advance() // consume the 'if' token

	if parser.currTok.Type != token.LPRAN {
		return &ast.InvalidExpression{
			Tok:     token.Lookup("if"),
			Message: fmt.Sprintf(internal.ErrExpectedOpenPren, parser.currTok.Type),
		}
	}
	condition := parseGroupedExpression(parser)
	switch cond := condition.(type) {
	case *ast.InvalidExpression:
		cond.Tok = ifExpr.Tok //TODO: revisit InvalidExpression Tok usage
		return cond
	default:
		ifExpr.Condition = cond
	}

	return ifExpr
}
