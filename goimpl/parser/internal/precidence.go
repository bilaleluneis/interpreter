package internal

import "goimpl/token"

type Operator string

// Define prefix operators.
const (
	_     Operator = ""
	BANG  Operator = "!"
	MINUS Operator = "-"
)

type Precidence int

// Define precedence levels from lowest to highest.
const (
	_           Precidence = iota
	LOWEST                 // Base level for starting expression parsing
	EQUALS                 // ==, !=
	LESSGREATER            // >, <
	SUM                    // +, -
	PRODUCT                // *, /
	PREFIX                 // -X, !X
	CALL                   // myFunction(X)
)

// PrecidenceMap defines operator precedence for all infix operators.
var PrecidenceMap = map[token.TokenType]Precidence{
	// Equality operators
	token.EQ:  EQUALS,
	token.NEQ: EQUALS,

	// Comparison operators
	token.LT: LESSGREATER,
	token.GT: LESSGREATER,

	// Arithmetic operators
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.ASTER: PRODUCT,

	// Call (function calls)
	token.LPRAN: CALL,
}
