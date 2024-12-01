package ast

interface AstNode{
    val tokenLiteral : String
    override fun toString(): String
}

interface Expression : AstNode

interface Statement : AstNode
