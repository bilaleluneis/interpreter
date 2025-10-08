package com.monkeylang.parser

import com.monkeylang.lexer.Lexer
import com.monkeylang.ast.Statement.*
import com.monkeylang.ast.Expression.*

protected def returnParser(lexer: Lexer): (Return, Lexer) =
  (Return(BooleanLiteral("true")), lexer)
