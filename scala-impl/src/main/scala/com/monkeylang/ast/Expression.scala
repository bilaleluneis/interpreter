package com.monkeylang.ast

enum Expression extends AstNode:
  case BooleanLiteral(value: String) // "true" or "false"
  case Error(err: String, exp: Expression)
