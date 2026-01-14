package com.monkeylang.ast

final class LetSuite extends munit.FunSuite:

  test("toString should return the correct format for Let statement"):
    val ident = Identifier("myVar")
    val value = Identifier("anotherVar")
    val letStmt = Let(ident, value)
    assertEquals(
      letStmt.toString,
      "ast.Let(ast.Identifier(myVar), ast.Identifier(anotherVar))"
    )
