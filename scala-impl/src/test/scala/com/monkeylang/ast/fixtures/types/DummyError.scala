package com.monkeylang.ast.fixtures.types

import com.monkeylang.ast.ErrorNode

final case class DummyError(val msg: String) extends ErrorNode:

  override def toString: String = s"DummyError($msg)"
