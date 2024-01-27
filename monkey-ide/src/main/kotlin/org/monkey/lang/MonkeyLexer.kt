package org.monkey.lang

import com.intellij.lexer.Lexer
import com.intellij.lexer.LexerPosition
import com.intellij.psi.TokenType
import com.intellij.psi.tree.IElementType
import org.monkey.lang.psi.MonkeyTypes

class MonkeyLexer : Lexer() {

    private var code: CharSequence = ""
    private var startOffset: Int = 0
    private var endOffset: Int = 0
    private var currState: Int = 0

    // to be used with overriden methods from Lexer class for current token
    private var currTokenStart: Int = 0
    private var currTokenEnd: Int = 0


    override fun getTokenStart(): Int = currTokenStart
    override fun getTokenEnd(): Int = currTokenEnd

    override fun getBufferSequence(): CharSequence = code
    override fun getState(): Int = currState
    override fun getBufferEnd(): Int = endOffset

    override fun start(buffer: CharSequence, startOffset: Int, endOffset: Int, initialState: Int) {
        this.code = buffer
        this.startOffset = startOffset //starts at 0
        this.endOffset = endOffset - 1 // [we can access 0 to endOffset - 1] and not overflow
        currState = initialState //starts at 0 stays at 0 during lexing and ends at 2 when nothing to lex
    }

    override fun getTokenType(): IElementType? {
        if (currState == 2) return null
        return with(code[startOffset]) {
            when {
                isWhitespace() -> whiteSpaceElement
                equals(':') -> seperatorElement
                else -> literalElement
            }
        }
    }

    override fun advance() {
        currState = 2
        if (startOffset < endOffset){
            currState = 0
            startOffset++
        }
    }

    override fun restore(position: LexerPosition) {
        startOffset = position.offset
        currState = position.state
    }

    override fun getCurrentPosition(): LexerPosition = object : LexerPosition {
        override fun getOffset(): Int = startOffset //TODO: is this correct calculation?
        override fun getState(): Int = currState
    }

    private val seperatorElement: IElementType
        get() {
            currTokenStart = startOffset
            currTokenEnd = startOffset + 1
            return MonkeyTypes.SEPARATOR
        }

    private val whiteSpaceElement: IElementType
        get() {
            currTokenStart = startOffset
            currTokenEnd = startOffset + 1
            return TokenType.WHITE_SPACE
        }

    private val literalElement: IElementType
        get() {
            var literal = code[startOffset].toString()
            currTokenStart = startOffset
            currTokenEnd = startOffset
            var peek: Char?
            while (peekAt(++currTokenEnd).also { peek = it } != null && peek!!.isLetter()) {
                literal += peek
            }
            startOffset = currTokenEnd - 1 // last peeked char is not part of literal
            return MonkeyTypes.KEY
        }

    private fun peekAt(index: Int): Char? {
        return if (index <= endOffset) code[index] else null
    }

}
