// Package lexer defines the lexer interface used by the interpreter.
package lexer

import (
	"goimpl/token"
)

// Lexer is an interface that can be used in stubbing tests
type Lexer interface {
	NextToken() token.Token
}
