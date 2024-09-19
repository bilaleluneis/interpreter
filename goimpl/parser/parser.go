package parser

import "goimpl/ast"

type ParserType interface {
	ParseProgram() (ast.Program, bool)
}

type PrefixParseFn func(ParserType) ast.Expression
type InfixParseFn func(ParserType, ast.Expression) ast.Expression
