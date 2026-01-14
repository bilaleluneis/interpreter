package com.monkeylang.ast

import com.monkeylang.ast.AstNodeType.ExpressionType

final class AstNodeTypeSuite extends munit.FunSuite:

  test("AstNodeType should have correc for identfier expression"):
    val identifier = ExpressionType(Identifier("x"))
    assertEquals(identifier.toString, "ExpressionType(ast.Identifier(x))")
