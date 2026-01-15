package com.monkeylang.ast.fixtures.types

import com.monkeylang.ast.Type

final case class IntType(val value: Int) extends Type:
  override def toString: String = s"ast.IntType($value)"
