package token

sealed interface Token

data object Illigal : Token
data object Eof : Token
data object Assign : Token
data object Plus : Token
data object Comma : Token
data object SemiColon : Token
data object Lpran : Token
data object Rpran : Token
data object Lbrace : Token
data object Rbrace : Token
data object Func : Token
data object Let : Token
data class IntValue(val value: Int) : Token
data class Identifier(val value: String) : Token


fun lookup(tok: String): Token {
    return with(tok) {
        when {
            isBlank() -> Eof
            equals("=") -> Assign
            equals("+") -> Plus
            equals(",") -> Comma
            equals(";") -> SemiColon
            equals("(") -> Lpran
            equals(")") -> Rpran
            equals("{") -> Lbrace
            equals("}") -> Rbrace
            equals("fn") -> Func
            equals("let") -> Let
            all { it.isDigit() } -> IntValue(toInt())
            all { it.isLetter() } -> Identifier(this)
            else -> Illigal
        }
    }
}
