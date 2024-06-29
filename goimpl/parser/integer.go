package parser

import (
	"goimpl/ast"
	"strconv"
)

var parseInteger prefixParseFn = func(parser *Parser) ast.Expression {
	literal := &ast.IntegerLiteral{Tok: parser.currTok}
	if value, err := strconv.ParseInt(parser.currTok.Literal, 0, 64); err == nil {
		literal.Value = value
		return literal
	}
	err := "could not parse " + parser.currTok.Literal + " as integer"
	parser.errors = append(parser.errors, err)
	return nil
}
