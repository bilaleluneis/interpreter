package com.monkeylang.parser

import com.monkeylang.lexer.Lexer
import com.monkeylang.ast.Statement.*
import com.monkeylang.ast.Expression.*

protected val returnParser: Lexer -> (Return, Lexer) = (l: Lexer) =>
  (Return(Some(BooleanLiteral("true"))), l)
