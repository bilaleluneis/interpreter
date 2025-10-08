package com.monkeylang.parser

import com.monkeylang.lexer.Lexer
import com.monkeylang.ast.{Statement, Expression}

enum ParserType(val parser: (Lexer) -> (Statement | Expression, Lexer)):
  case ReturnParserType extends ParserType(returnParser)
