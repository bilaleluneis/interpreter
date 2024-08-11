package lex

import token.Token
import java.io.Serializable

interface Lexer : Serializable {
    fun next(): Token
}
