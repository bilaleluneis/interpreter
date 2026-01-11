package com.monkeylang.token

final class TokensSuite extends munit.FunSuite:

  private val tokens = Tokens(
    Token.LET,
    Token.IDENT("myVar"),
    Token.ASSIGN,
    Token.INT(5),
    Token.SEMICOLON
  )

  test("Tokens toString"):
    assertEquals(tokens.toString, "[LET IDENT(myVar) ASSIGN INT(5) SEMICOLON]")

  test("Tokens iteration"):
    val it = tokens.iterator
    while it.hasNext do println(it.next())
