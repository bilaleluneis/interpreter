package com.monkeylang.token

import com.monkeylang.token.Token

final class Tokens(tokens: Token*) extends Iterable[Token]:
  override def iterator: Iterator[Token] = new Iterator[Token]:
    private var pos = 0
    def remaining: Seq[Token] = tokens.drop(pos)
    override def hasNext: Boolean = pos < tokens.length
    override def next(): Token =
      if hasNext then
        val tok = tokens(pos)
        pos += 1
        tok
      else Token.EOF
    override def toString: String =
      s"[${remaining.map(_.toString).mkString(" ")}]"

  override def toString: String =
    s"[${tokens.map(_.toString).mkString(" ")}]"
