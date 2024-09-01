package lexer

import (
	"goimpl/common"
	"goimpl/token"
)

// Lexer is an interface that can be used in stubbing tests
type Lexer interface {
	NextToken() token.Token
}

type CopyableLexer[T any] interface {
	Lexer
	common.VType[T]
	*T
}
