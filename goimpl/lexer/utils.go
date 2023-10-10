package lexer

import (
	"slices"
)

var whiteSpace = []byte{
	' ',  //space
	'\t', //tab
	'\n', //end of line
	'\r', //carriage return
}

var operators = []byte{
	'=',
	'+',
}

//var keywords = []String{
//	"let",
//	"fn",
//}

func isWhiteSpace(ch byte) bool { return slices.Contains(whiteSpace, ch) }

func isLetter(ch byte) bool { return 'a' <= ch || ch <= 'z' || 'A' <= ch || ch <= 'Z' || ch == '_' }

func isOperator(ch byte) bool { return slices.Contains(operators, ch) }
