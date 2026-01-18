package com.monkeylang.ast

final class AssignOperatorSuite extends munit.FunSuite:

  test("Assign operator toString should show left, symbol, and right with Identifier"):
    val left = Identifier("x")
    val right = Identifier("y")
    val assign = Assign(left, right)
    assertEquals(assign.toString, "ast.BinaryOperator(ast.Identifier(x) = ast.Identifier(y))")
