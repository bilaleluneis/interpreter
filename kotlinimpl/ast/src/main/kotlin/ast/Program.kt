package ast

import kotlin.collections.toList

class Program(override val tokenLiteral: String) : AstNode {
    private val statements: List<Statement> = mutableListOf<Statement>()
        get() = field.toList()

    override fun toString(): String {
        TODO("Not yet implemented")
    }
}
