package com.monkeylang.token

final class TokensSuite extends munit.FunSuite:

  private val tokens = of(
    Token.LET,
    Token.IDENT("myVar"),
    Token.ASSIGN,
    Token.INT(5),
    Token.SEMICOLON
  )

  test("Tokens toString"):
    assertEquals(tokens.toString, "[LET IDENT(myVar) = INT(5) ;]")

  test("Tokens iteration"):
    tokens.filter(_ != Token.SEMICOLON).toList match
      case List(Token.LET, Token.IDENT(name), Token.ASSIGN, Token.INT(value)) =>
        assertEquals(name, "myVar")
        assertEquals(value, 5)
      case _ =>
        fail("Tokens did not match expected sequence")
