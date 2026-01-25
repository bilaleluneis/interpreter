package com.monkeylang.ast.fixtures.types

import com.monkeylang.ast.Type
import com.monkeylang.ast.ValueType

final case class IntType(val value: Int) extends ValueType:
  override def toString: String = s"ast.IntType($value)"
