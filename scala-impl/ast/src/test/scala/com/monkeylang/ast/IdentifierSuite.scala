package com.monkeylang.ast

import munit.FunSuite
import munit.FunSuite

final class IdentifierSuite extends FunSuite:

  test("toString should return the correct format for 'foo'"):
    val id = Identifier("foo")
    assertEquals(id.toString, "ast.Identifier(foo)")
