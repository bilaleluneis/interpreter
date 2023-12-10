package lex

import token.Eof
import token.Token
import token.lookup

class Lexer(private val input: String) {

    private val endOfFileFlag = (0).toChar()
    private var prevReadPos = 0
    private var nextReadPos = 0

    private val underInspection: Char
        get() {
            val ch = peek
            prevReadPos = nextReadPos
            nextReadPos++
            return ch
        }

    private val peek get() = input.elementAtOrElse(nextReadPos) { _ -> endOfFileFlag }

    fun next(): Token {
        return with(underInspection) {
            when {
                isWhitespace() -> next()
                equals(endOfFileFlag) -> Eof
                isDigit() -> digitToken(this)
                isLetter() -> literalToken(this)
                else -> lookup(toString())
            }
        }
    }

    private fun literalToken(initial: Char): Token {
        var literal = initial.toString()
        while (peek.isLetter()) {
            literal += underInspection
        }
        return lookup(literal)
    }

    private fun digitToken(initial: Char): Token {
        var digits = initial.toString()
        while (peek.isDigit()) {
            digits += underInspection
        }
        return lookup(digits)
    }

}

