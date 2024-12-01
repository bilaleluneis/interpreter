package ast

object Eof : Statement {
    override val tokenLiteral: String
        get() = "EOF"
    override fun toString(): String {
        return tokenLiteral
    }
}
