package com.monkeylang.lexer

final class LazyLexer(input: String) extends Lexer:
  override def iterator: Iterator[Token] =
    new Iterator[Token]:
      private var nextToken: Token = ???
      override def hasNext: Boolean = ???
      override def next(): Token = Token.EOF
