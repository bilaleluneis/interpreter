package com.monkeylang.lexer

import munit.FunSuite
import scala.collection.immutable.ArraySeq

final class MockLexerSuite extends FunSuite:

  private val tokens = ArraySeq(
    Token.LET,
    Token.IDENT("five"),
    Token.ASSIGN,
    Token.INT(5),
    Token.SEMICOLON,
    Token.EOF
  )

  private val lexer: Lexer = MockLexer(tokens)

  test("MockLexer should return tokens in order"):
    val result = lexer.toList
    assertEquals(result, tokens.toList)

  test("MockLexer allow use of iterators"):
    lexer.foreach(print)
