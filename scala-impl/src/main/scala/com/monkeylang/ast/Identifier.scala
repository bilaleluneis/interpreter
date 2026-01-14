package com.monkeylang.ast

final class Identifier(ident: String) extends ExpressionNode:
  override def toString: String = s"ast.Identifier($ident)"
