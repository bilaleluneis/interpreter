package lexer

import "goimpl/token"

type LazyLexer struct {
	input               string
	prevReadIndx        int
	nextReadIndex       int
	charUnderInspection byte
}

func NewLazyLexer(input string) LazyLexer {
	return LazyLexer{input: input}
}

func (l *LazyLexer) GetCopy() *LazyLexer {
	lcopy := LazyLexer{input: l.input}
	lcopy.prevReadIndx = l.prevReadIndx
	lcopy.nextReadIndex = l.nextReadIndex
	lcopy.charUnderInspection = l.charUnderInspection
	return &lcopy
}

// NextToken TODO: refactor to reduce if/else and manual advanceLoc()
func (l *LazyLexer) NextToken() token.Token {

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

func (l *LazyLexer) advanceLoc() {
	l.prevReadIndx = l.nextReadIndex
	l.nextReadIndex++
}

func (l *LazyLexer) peek() byte {
	if l.nextReadIndex >= len(l.input) {
		return 0
	}
	return l.input[l.nextReadIndex]
}

func (l *LazyLexer) captureLiteral() string {
	var litral string
	for ; !isWhiteSpace(l.peek()) && isLetter(l.peek()); l.advanceLoc() {
		letter := l.peek()
		l.charUnderInspection = letter
		litral += string(l.charUnderInspection)
	}
	return litral
}

func (l *LazyLexer) captureDigits() string {
	var digits string
	for ; !isWhiteSpace(l.peek()) && isDigit(l.peek()); l.advanceLoc() {
		digit := l.peek()
		l.charUnderInspection = digit
		digits += string(l.charUnderInspection)
	}
	return digits
}
