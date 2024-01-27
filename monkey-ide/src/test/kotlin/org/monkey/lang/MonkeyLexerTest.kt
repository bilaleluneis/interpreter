package org.monkey.lang

internal class MonkeyLexerTest : MonkeyLexerTestBase() {
    fun testKey() = match("foo", "MonkeyTokenType.KEY ('foo')")
    fun testValue() = match(
        "foo :",
        "MonkeyTokenType.KEY",
        "WHITE_SPACE",
        "MonkeyTokenType.SEPARATOR",
    )
}

