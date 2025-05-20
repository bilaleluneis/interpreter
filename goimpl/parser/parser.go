package parser

import "goimpl/ast"

// ParserType defines the core functionality of a parser implementation.
// Any parser implementation must be able to parse a complete program and
// indicate whether parsing was successful or return ast.Error.
type ParserType interface {
	// ParseProgram parses a complete program and returns the AST or ast.Error.
	ParseProgram() ast.Program
}
