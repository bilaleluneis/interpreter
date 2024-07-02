package parser

import "goimpl/token"

type precidence int

const (
	_ precidence = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precidenceMap = map[token.TokenType]precidence{
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.ASTER: PRODUCT,
	token.LT:    LESSGREATER,
	token.GT:    LESSGREATER,
	token.EQ:    EQUALS,
	token.NEQ:   EQUALS,
}
