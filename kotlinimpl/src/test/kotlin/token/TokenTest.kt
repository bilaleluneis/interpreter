package token

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import org.junit.jupiter.api.Assertions.assertEquals
import java.util.stream.Stream

internal class TokenTest {

    @ParameterizedTest
    @MethodSource("tokens")
    fun tokens(input: String, expected: Token) {
        assertEquals(lookup(input), expected)
    }

    // Tests Source
    private companion object {
        @JvmStatic
        fun tokens(): Stream<Arguments> = Stream.of(
            Arguments.of("", Eof),
            Arguments.of("=", Assign),
            Arguments.of("+", Plus),
            Arguments.of(",", Comma),
            Arguments.of(";", SemiColon),
            Arguments.of("(", Lpran),
            Arguments.of(")", Rpran),
            Arguments.of("{", Lbrace),
            Arguments.of("}", Rbrace),
            Arguments.of("fn", Func),
            Arguments.of("let", Let),
            Arguments.of("10", IntValue(10)),
            Arguments.of("tmp", Identifier("tmp")),
            Arguments.of("<TP@", Illigal),
        )
    }

}
