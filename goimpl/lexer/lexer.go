package lexer

import (
	"goimpl/common"
	"goimpl/token"
)

// Lexer is an interface that can be used in stubbing tests
type Lexer interface {
	NextToken() token.Token
}

func CopyOf[V common.VType[L], L Lexer](lexer V) L {
	return lexer.GetCopy()
}
