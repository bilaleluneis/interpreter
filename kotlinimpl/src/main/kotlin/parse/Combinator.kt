package parse

import lex.Lexer
import org.apache.commons.lang3.SerializationUtils

class Combinator(val lexer: Lexer) : () -> Combinator {
    override fun invoke(): Combinator {
        val updatedLexer = SerializationUtils.clone(lexer)
        val result = Combinator(updatedLexer)
        return result
    }
}
