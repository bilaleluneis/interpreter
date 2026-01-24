package com.monkeylang.lexer
import com.monkeylang.token.Token

final class LazyLexer(input: String) extends Lexer:
  override def iterator: Iterator[Token] =
    new Iterator[Token]:
      private var chars: List[Char] = input.toList
      override def hasNext: Boolean = chars.nonEmpty
      override def next(): Token =
        chars match
          case head :: tail =>
            chars = tail
            Token.EOF
          case _ =>
            Token.EOF
