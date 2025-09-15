package pratt

import (
	"fmt"
	"goimpl/ast"
	"goimpl/parser/internal"
	"goimpl/token"
)

// parseIfExpression parses an if expression.
// if(<condition>) { <consequence> } else { <alternative> }
// TODO: need to handle else part
// TODO: need to clean up code and reduce size ..
// have parser return something like (ast.Expression, error) instead of
var parseIfExpression PrefixParseFn = func(parser *Parser) ast.Expression {
	ifExpr := &ast.IfExpression{
		Tok: token.Lookup("if"),
	}

	parser.advance() // consume 'if'

	if parser.currTok.Type != token.LPRAN {
		return &ast.InvalidExpression{
			Tok:     token.Lookup("if"),
			Message: fmt.Sprintf(internal.ErrExpectedOpenPren, parser.currTok.Type),
		}
	}

	if parser.peekTok.Type == token.RPRAN {
		return &ast.InvalidExpression{
			Tok:     token.Lookup("if"),
			Message: internal.ErrEmptyExpression,
		}
	}

	condition := parser.parseExpression(internal.LOWEST)
	switch cond := condition.(type) {
	case *ast.InvalidExpression:
		cond.Tok = token.Lookup("if") //TODO: revisit InvalidExpression Tok usage
		return cond
	default:
		ifExpr.Condition = cond
	}

	if parser.currTok.Type != token.RPRAN {
		return &ast.InvalidExpression{
			Tok:     token.Lookup("if"),
			Message: fmt.Sprintf(internal.ErrExpectedClosePren, parser.currTok.Type),
		}
	}

	parser.advance() // consume ')'

	if parser.currTok.Type != token.LBRACE {
		return &ast.InvalidExpression{
			Tok:     token.Lookup("if"),
			Message: fmt.Sprintf(internal.ErrExpectedOpenBrace, parser.currTok.Type),
		}
	}

	consequence := parser.parseBlockStatement()
	switch cons := consequence.(type) {
	case *ast.Block:
		ifExpr.Conseq = *cons
	default:
		return &ast.InvalidExpression{
			Tok:     token.Lookup("if"),
			Message: "Expected block statement as consequence",
		}
	}

	return ifExpr
}
