package parse

import lex.LazyLexer
import org.junit.jupiter.api.Test

internal class CombinatorTest {
    @Test
    fun initCombinator() {
//        val combinator = Combinator(LazyLexer("=+,;(){}"))
//        val result = combinator()
//        println(result)
        takeCombinator { Combinator(LazyLexer("=+,;(){}")) }
    }
}

fun takeCombinator(combinator: () -> Combinator) {
    val result: Combinator = combinator()
    result.invoke()
    println(result)
}
