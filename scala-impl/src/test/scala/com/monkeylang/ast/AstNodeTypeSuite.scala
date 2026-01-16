package com.monkeylang.ast

import com.monkeylang.ast.AstNodeType.ExpressionType
import com.monkeylang.ast.fixtures.types.IntType
import com.monkeylang.ast.fixtures.types.ComplexAst
import com.monkeylang.ast.fixtures.types.DummyError

final class AstNodeTypeSuite extends munit.FunSuite:

  test("switch over AstNode"):
    val exprNode: AstType = ExpressionType(Identifier("myVar"))

    exprNode match
      case ExpressionType(expr) =>
        assertEquals(expr.toString, "ast.Identifier(myVar)")
      case _ =>
        fail("Expected ExpressionType but got different AstNodeType")

  test("AstNodeType should have correc for identfier expression"):
    val identifier = ExpressionType(Identifier("x"))
    assertEquals(identifier.toString, "ExpressionType(ast.Identifier(x))")

  test("AstNodeType should have correct for integer expression"):
    val intExpr = ExpressionType(IntType(10))
    assertEquals(intExpr.toString, "ExpressionType(ast.IntType(10))")

  test("able stringify ComplexAst node"):
    val complexAst = ComplexAst(
      e = Identifier("y"),
      s = Let(Identifier("z"), IntType(5)),
      er = DummyError("an error occurred")
    )

    var complexAstStrWant = "ComplexAst("
    complexAstStrWant += "ast.Identifier(y), "
    complexAstStrWant += "ast.Let(ast.Identifier(z), ast.IntType(5)), "
    complexAstStrWant += "DummyError(an error occurred))"

    assertEquals(complexAst.toString, complexAstStrWant)

  test("able to switch on ComplexAst node type"):
    val complexAst = ComplexAst(
      e = Identifier("y"),
      s = Let(Identifier("z"), IntType(5)),
      er = DummyError("an error occurred")
    )

    complexAst match
      case ComplexAst(e, s, er) =>
        assertEquals(e.toString, "ast.Identifier(y)")
        assertEquals(s.toString, "ast.Let(ast.Identifier(z), ast.IntType(5))")
        assertEquals(er.toString, "DummyError(an error occurred)")
