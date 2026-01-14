package com.monkeylang.ast

sealed trait AstNode

trait ExpressionNode extends AstNode
trait StatementNode extends AstNode
trait ErrorNode extends AstNode

enum AstNodeType:
  case ErrorType(err: ErrorNode)
  case ExpressionType(expr: ExpressionNode)
  case StatementType(stmt: StatementNode)

  override def toString: String =
    this match
      case ErrorType(err)       => s"ErrorType($err)"
      case ExpressionType(expr) => s"ExpressionType($expr)"
      case StatementType(stmt)  => s"StatementType($stmt)"
