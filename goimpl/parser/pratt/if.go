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
var parseIfExpression PrefixParseFn = func(parser *Parser) ast.Expression {

	ifExpr := &ast.IfExpression{Tok: token.Lookup("if")}
	err := &ast.InvalidExpression{Tok: ifExpr.Tok}
	parser.advance() // consume 'if'

	// Validate '(' and non-empty condition
	if parser.currTok.Type != token.LPRAN {
		err.Message = fmt.Sprintf(internal.ErrExpectedOpenPren, parser.currTok.Type)
		return err
	}
	if parser.peekTok.Type == token.RPRAN {
		err.Message = internal.ErrEmptyExpression
		return err
	}

	// Parse condition
	cond := parser.parseExpression(internal.LOWEST)
	if invalid, ok := cond.(*ast.InvalidExpression); ok {
		invalid.Tok = ifExpr.Tok
		return invalid
	}
	ifExpr.Condition = cond

	// Validate ')'
	if parser.currTok.Type != token.RPRAN {
		err.Message = fmt.Sprintf(internal.ErrExpectedClosePren, parser.currTok.Type)
		return err
	}
	parser.advance() // consume ')'

	// Validate '{'
	if parser.currTok.Type != token.LBRACE {
		err.Message = fmt.Sprintf(internal.ErrExpectedOpenBrace, parser.currTok.Type)
		return err
	}

	// Parse consequence block
	cons := parser.parseBlockStatement()
	if block, ok := cons.(*ast.Block); ok {
		ifExpr.Conseq = *block

		// Optional else block
		if parser.currTok.Type == token.ELSE {
			// consume 'else'
			parser.advance()
			if parser.currTok.Type != token.LBRACE {
				err.Message = fmt.Sprintf(internal.ErrExpectedOpenBrace, parser.currTok.Type)
				return err
			}
			alt := parser.parseBlockStatement()
			if altBlock, ok := alt.(*ast.Block); ok {
				ifExpr.Alt = *altBlock
			} else {
				err.Message = "Expected block statement as alternative"
				return err
			}
		}

		return ifExpr
	}

	err.Message = "Expected block statement as consequence"
	return err
}
