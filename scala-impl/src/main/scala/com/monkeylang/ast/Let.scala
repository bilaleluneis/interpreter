package com.monkeylang.ast

final class Let(ident: Identifier, value: ExpressionNode) extends StatementNode:
  override def toString: String = s"ast.Let($ident, $value)"
