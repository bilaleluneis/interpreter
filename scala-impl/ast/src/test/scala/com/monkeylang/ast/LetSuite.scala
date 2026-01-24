package com.monkeylang.ast

import com.monkeylang.ast.fixtures.types.IntType

final class LetSuite extends munit.FunSuite:

  test("toString should return the correct format for Let statement"):
    val ident = Identifier("myVar")
    val value = IntType(42)
    val letStmt = Let(ident, value)
    assertEquals(
      letStmt.toString,
      "ast.Let(ast.Identifier(myVar), ast.IntType(42))"
    )
