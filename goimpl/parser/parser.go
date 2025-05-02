package parser

import "goimpl/ast"

// ParserType defines the core functionality of a parser implementation.
// Any parser implementation must be able to parse a complete program and
// indicate whether parsing was successful.
type ParserType interface {
	// ParseProgram parses a complete program and returns the AST along with a success flag.
	// The boolean return indicates if parsing completed without errors.
	ParseProgram() (ast.Program, bool)
}
