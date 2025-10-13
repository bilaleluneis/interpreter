package com.monkeylang.parser

import com.monkeylang.lexer.Lexer
import com.monkeylang.ast.Expression.*
import com.monkeylang.ast.Error

protected val booleanLiteralParser: Lexer -> (BooleanLiteral | Error, Lexer) =
  (l: Lexer) =>
    l.iterator.next match
      case token if token.toString == "false" => (BooleanLiteral("false"), l)
      case token if token.toString == "true"  => (BooleanLiteral("true"), l)
      case _                                  => (Error(""), l)
