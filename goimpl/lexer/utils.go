package lexer

import (
	"goimpl/token"
	"slices"
)

var whiteSpace = []byte{
	' ',  //space
	'\t', //tab
	'\n', //end of line
	'\r', //carriage return
}

var operators = map[byte]token.TokenType{
	'=': token.ASSIGN,
	'+': token.PLUS,
}

var seperators = map[byte]token.TokenType{
	',': token.COMMA,
	';': token.SEMICOLON,
}

var keywords = map[string]token.TokenType{
	"let":    token.LET,
	"fn":     token.FUNCTION,
	"return": token.RETURN,
}

var blocksAndSubscripts = map[byte]token.TokenType{
	'(': token.LPRAN,
	')': token.RPRAN,
	'{': token.LBRACE,
	'}': token.RBRACE,
}

func isWhiteSpace(ch byte) bool { return slices.Contains(whiteSpace, ch) }

func isLetter(ch byte) bool {
	if isDigit(ch) || isSeperator(ch) || isOperator(ch) || isBlockOrSubscript(ch) {
		return false
	}
	return 'a' <= ch || ch <= 'z' || 'A' <= ch || ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	if isSeperator(ch) || isOperator(ch) || isBlockOrSubscript(ch) {
		return false
	}
	return '0' <= ch && ch <= '9'
}

func isOperator(ch byte) bool { return containsKey(operators, ch) }

func isSeperator(ch byte) bool { return containsKey(seperators, ch) }

func isBlockOrSubscript(ch byte) bool { return containsKey(blocksAndSubscripts, ch) }

func isEof(ch byte) bool { return ch == 0 }

// https://stackoverflow.com/questions/62652236/can-i-have-a-function-to-check-if-a-key-is-in-a-map
func containsKey[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, ok := m[k]
	return ok
}
