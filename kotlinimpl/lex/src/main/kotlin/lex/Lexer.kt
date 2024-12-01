package lex

import java.io.Serializable

interface Lexer : Serializable {
    fun next(): Token
}
