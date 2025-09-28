package com.monkeylang.lexer

import munit.FunSuite
import scala.collection.immutable.ArraySeq

final class MockLexerSuite extends FunSuite:
  test("MockLexer should return tokens in order"):
    val tokens = ArraySeq(
      Token.LET,
      Token.IDENT("five"),
      Token.ASSIGN,
      Token.INT(5),
      Token.SEMICOLON,
      Token.EOF
    )
    val lexer: Lexer = MockLexer(tokens)
    val result = lexer.toList
    assertEquals(result, tokens.toList)
