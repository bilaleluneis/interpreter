package com.monkeylang.ast

sealed trait AstNode

final case class Error(msg: String) extends AstNode

enum Expression extends AstNode:
  case BooleanLiteral(value: "true" | "false")

enum Statement extends AstNode:
  case Return(value: Option[Expression])
