package parser

import "goimpl/ast"

type ParserType interface {
	ParseProgram() (ast.Program, bool)
}
