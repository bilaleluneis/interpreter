package internal

import "goimpl/token"

type Operator string

const (
	_     Operator = ""
	BANG  Operator = "!"
	MINUS Operator = "-"
)

type Precidence int

const (
	_ Precidence = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var PrecidenceMap = map[token.TokenType]Precidence{
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.ASTER: PRODUCT,
	token.LT:    LESSGREATER,
	token.GT:    LESSGREATER,
	token.EQ:    EQUALS,
	token.NEQ:   EQUALS,
}
