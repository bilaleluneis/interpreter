package org.monkey.lang

import com.intellij.lexer.Lexer
import com.intellij.psi.tree.IElementType
import com.intellij.testFramework.LexerTestCase

internal abstract class MonkeyLexerTestBase : LexerTestCase() {
    override fun getDirPath() = "" //TODO "src/test/resources/org/monkey/lang" to use test data
    override fun createLexer(): Lexer = MonkeyLexer()
    fun match(code: String, tokens: String) = doTest(code, tokens)
    fun match(code: String, vararg tokens: String) {
        val lexer = createLexer()
        lexer.start(code)
        var index = 0
        var tokenType: IElementType?
        while (lexer.tokenType.also{tokenType = it} != null) {
            //TODO: use lexer.tokenText to verify token text in source
            assertEquals(tokens[index], tokenType.toString())
            lexer.advance()
            index++
        }
    }
}

