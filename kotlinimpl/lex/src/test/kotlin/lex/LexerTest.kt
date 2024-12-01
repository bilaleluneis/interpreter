package lex

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

internal class LexerTest {

    @Test
    fun basicLexing() {
        val lexer = LazyLexer("=+,;(){}")
        arrayOf(
            Assign,
            Plus,
            Comma,
            SemiColon,
            Lpran,
            Rpran,
            Lbrace,
            Rbrace,
            Eof,
        ).forEach { assertEquals(lexer.next(), it) }
    }

    @Test
    fun minLangConstruct() {
        val input = """
            let five = 5;
            let ten = 10;
            let add = fn(x,y) { x + y};
            let result = add(five, ten);
            """
        val lexer = LazyLexer(input)
        arrayOf(
            // let five = 5;
            Let,
            Identifier("five"),
            Assign,
            IntValue(5),
            SemiColon,

            // let ten = 10;
            Let,
            Identifier("ten"),
            Assign,
            IntValue(10),
            SemiColon,

            // let add = fn(x,y){x+y};
            Let,
            Identifier("add"),
            Assign,
            Func,
            Lpran,
            Identifier("x"),
            Comma,
            Identifier("y"),
            Rpran,
            Lbrace,
            Identifier("x"),
            Plus,
            Identifier("y"),
            Rbrace,
            SemiColon,

            // let result = add(five, ten);
            Let,
            Identifier("result"),
            Assign,
            Identifier("add"),
            Lpran,
            Identifier("five"),
            Comma,
            Identifier("ten"),
            Rpran,
            SemiColon,

            Eof
        ).forEach { assertEquals(lexer.next(), it) }
    }

}
