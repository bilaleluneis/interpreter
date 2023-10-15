package lexer

import (
	"goimpl/token"
)

type Lexer struct {
	input               string
	prevReadIndx        int
	nextReadIndex       int
	charUnderInspection byte
}

func New(input string) (*Lexer, bool) {
	if len(input) == 0 {
		return nil, false
	}
	return &Lexer{input: input}, true
}

// NextToken TODO: refactor to reduce if/else and manual advanceLoc()
func (l *Lexer) NextToken() token.Token {

	for isWhiteSpace(l.peek()) {
		l.advanceLoc()
	}

	l.charUnderInspection = l.peek()

	tokenToReturn := token.NewToken(token.ILLIGAL, l.charUnderInspection)
	if isEof(l.charUnderInspection) {
		tokenToReturn = token.Token{Type: token.EOF, Literal: ""}
	} else if isOperator(l.charUnderInspection) {
		tokenToReturn = token.NewToken(operators[l.charUnderInspection], l.charUnderInspection)
		l.advanceLoc()
	} else if isSeperator(l.charUnderInspection) {
		tokenToReturn = token.NewToken(seperators[l.charUnderInspection], l.charUnderInspection)
		l.advanceLoc()
	} else if isBlockOrSubscript(l.charUnderInspection) {
		tokenToReturn = token.NewToken(blocksAndSubscripts[l.charUnderInspection], l.charUnderInspection)
		l.advanceLoc()
	} else if isDigit(l.charUnderInspection) {
		digits := l.captureDigits()
		tokenToReturn = token.Token{Type: token.INT, Literal: digits}
	} else if isLetter(l.charUnderInspection) {
		//capture the litral string
		litral := l.captureLiteral()

		//check if literal is a keyword
		if tokenType, ok := keywords[litral]; ok {
			tokenToReturn = token.Token{Type: tokenType, Literal: litral}
		} else {
			tokenToReturn = token.Token{Type: token.IDENTIFIER, Literal: litral}
		}
	}

	return tokenToReturn

}

func (l *Lexer) advanceLoc() {
	l.prevReadIndx = l.nextReadIndex
	l.nextReadIndex++
}

func (l *Lexer) peek() byte {
	if l.nextReadIndex >= len(l.input) {
		return 0
	}
	return l.input[l.nextReadIndex]
}

func (l *Lexer) captureLiteral() string {
	var litral string
	for ; !isWhiteSpace(l.peek()) && isLetter(l.peek()); l.advanceLoc() {
		letter := l.peek()
		litral += string(letter)
		l.charUnderInspection = letter
	}
	return litral
}

func (l *Lexer) captureDigits() string {
	var digits string
	for ; !isWhiteSpace(l.peek()) && isDigit(l.peek()); l.advanceLoc() {
		digit := l.peek()
		digits += string(digit)
		l.charUnderInspection = digit
	}
	return digits
}
