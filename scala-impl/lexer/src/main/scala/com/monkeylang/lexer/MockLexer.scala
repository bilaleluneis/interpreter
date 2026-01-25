package com.monkeylang.lexer

import scala.collection.immutable.ArraySeq
import com.monkeylang.token.Token

final class MockLexer(toks: ArraySeq[Token]) extends Lexer, Iterable[Token]:
  override def iterator: Iterator[Token] =
    new Iterator[Token]:
      private var pos: Int = 0
      override def hasNext: Boolean = pos < toks.length
      override def next(): Token =
        if hasNext then
          val tok = toks(pos)
          pos += 1
          tok
        else Token.EOF
