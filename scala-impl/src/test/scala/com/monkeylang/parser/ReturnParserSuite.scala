package com.monkeylang.parser

import com.monkeylang.lexer.{Lexer, MockLexer, Token}
import com.monkeylang.ast.Statement.Return
import com.monkeylang.ast.Expression.BooleanLiteral
import scala.collection.immutable.ArraySeq

final class ReturnParserSuite extends munit.FunSuite:

  test("Return parser should parse return statements"):

    val lexer: Lexer = MockLexer(
      ArraySeq(Token.RETURN, Token.TRUE, Token.SEMICOLON, Token.EOF)
    )

    val (ast, _) = returnParser(lexer)
    val expected = Return(Some(BooleanLiteral("true")))

    assertEquals(ast, expected)
