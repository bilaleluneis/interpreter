package com.monkeylang.ast

import com.monkeylang.ast.AstNodeType.ExpressionType
import com.monkeylang.ast.fixtures.types.IntType

final class AstNodeTypeSuite extends munit.FunSuite:

  test("AstNodeType should have correc for identfier expression"):
    val identifier = ExpressionType(Identifier("x"))
    assertEquals(identifier.toString, "ExpressionType(ast.Identifier(x))")

  test("AstNodeType should have correct for integer expression"):
    val intExpr = ExpressionType(IntType(10))
    assertEquals(intExpr.toString, "ExpressionType(ast.IntType(10))")
