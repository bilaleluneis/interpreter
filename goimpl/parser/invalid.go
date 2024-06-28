package parser

import (
	"fmt"
	"goimpl/ast"
)

var parseInvalid parseStatement = func(p *Parser) (ast.Statement, error) {
	return nil, fmt.Errorf("invalid token %s", p.currTok.Type)
}
