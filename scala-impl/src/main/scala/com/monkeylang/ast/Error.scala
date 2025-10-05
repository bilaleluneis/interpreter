package com.monkeylang.ast

enum Error extends AstNode:
  case Generic(message: String)
