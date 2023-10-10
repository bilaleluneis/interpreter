package lexer

import "goimpl/token"

type Lexer struct {
	input               string
	prevReadIndx        int
	nextReadIndex       int
	charUnderInspection byte
}

func New(input string) *Lexer { return &Lexer{input: input} }

func (l *Lexer) NextToken() token.Token {

	for l.selectCharToInspect(); isWhiteSpace(l.charUnderInspection); l.selectCharToInspect() {
	}

	//check if this is a IDENTIFIER or KEYWORD token
	if isLetter(l.charUnderInspection) {

	}

	switch l.charUnderInspection {

	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	case '=':
		return token.NewToken(token.ASSIGN, l.charUnderInspection)
	case '+':
		return token.NewToken(token.PLUS, l.charUnderInspection)
	case ',':
		return token.NewToken(token.COMMA, l.charUnderInspection)
	case '(':
		return token.NewToken(token.LPRAN, l.charUnderInspection)
	case ')':
		return token.NewToken(token.RPRAN, l.charUnderInspection)
	case '{':
		return token.NewToken(token.LBRACE, l.charUnderInspection)
	case '}':
		return token.NewToken(token.RBRACE, l.charUnderInspection)
	case ';':
		return token.NewToken(token.SEMICOLON, l.charUnderInspection)
	default:
		return token.NewToken(token.ILLIGAL, l.charUnderInspection)

	}

}

func (l *Lexer) selectCharToInspect() {
	if l.nextReadIndex >= len(l.input) {
		l.charUnderInspection = 0
	} else {
		l.charUnderInspection = l.input[l.nextReadIndex]
	}
	l.prevReadIndx = l.nextReadIndex
	l.nextReadIndex++
}
