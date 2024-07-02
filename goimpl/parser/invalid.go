package parser

import (
	"fmt"
	"goimpl/ast"
)

var parseInvalid parseStatement = func(p *Parser) ast.Statement {
	p.errors = append(p.errors, fmt.Sprintf("invalid token %s", p.currTok.Type))
	return nil
}
