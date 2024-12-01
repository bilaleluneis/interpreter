package ast

class Error(val message : String ) : Statement {
    override val tokenLiteral: String
        get() = "ERROR"
    override fun toString(): String {
        return tokenLiteral
    }
}


