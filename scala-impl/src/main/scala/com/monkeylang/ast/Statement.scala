package com.monkeylang.ast

enum Statement extends AstNode:
  case Return(value: Expression)
  case Error(err: String, stmt: Statement)
