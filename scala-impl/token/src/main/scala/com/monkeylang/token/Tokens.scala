package com.monkeylang.token

import com.monkeylang.token.Token

def of(tokens: Token*): Iterable[Token] =
  new Iterable[Token]:
    // overrides toString on the iterable which is returned
    override def toString: String = "[" + tokens.map(_.toString).mkString(" ") + "]"
    override def iterator: Iterator[Token] =
      new Iterator[Token]:
        private var pos: Int = 0
        override def hasNext: Boolean = pos < tokens.length
        override def next(): Token =
          if hasNext then
            val tok = tokens(pos)
            pos += 1
            tok
          else Token.EOF
        // overrides toString on the iterator
        override def toString: String = "[" + tokens.map(_.toString) + "]"
