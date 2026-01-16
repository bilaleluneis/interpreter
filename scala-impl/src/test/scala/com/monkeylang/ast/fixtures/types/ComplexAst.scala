package com.monkeylang.ast.fixtures.types

import com.monkeylang.ast.StatementNode
import com.monkeylang.ast.ExpressionNode
import com.monkeylang.ast.ErrorNode

class ComplexAst(val e: ExpressionNode, val s: StatementNode, val er: ErrorNode)
    extends StatementNode:

  override def toString: String = s"ComplexAst($e, $s, $er)"

object ComplexAst:
  def unapply(complexAst: ComplexAst): (ExpressionNode, StatementNode, ErrorNode) =
    (complexAst.e, complexAst.s, complexAst.er)
