package com.monkeylang.ast

enum Statement extends AstNode:
  case Return(value: Expression)
