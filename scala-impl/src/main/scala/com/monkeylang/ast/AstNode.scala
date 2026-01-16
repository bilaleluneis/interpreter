package com.monkeylang.ast

sealed trait AstNode

trait ExpressionNode extends AstNode
trait StatementNode extends AstNode
trait ErrorNode extends AstNode

// enum to represent different AST node types
// A is a covariant type parameter bounded by AstNode
// This allows AstNodeType to work with any subtype of AstNode sealed trait
enum AstNodeType[+A <: AstNode]:
  case ErrorType(err: ErrorNode)
  case ExpressionType(expr: ExpressionNode)
  case StatementType(stmt: StatementNode)

  override def toString: String =
    this match
      case ErrorType(err)       => s"ErrorType($err)"
      case ExpressionType(expr) => s"ExpressionType($expr)"
      case StatementType(stmt)  => s"StatementType($stmt)"

type AstType = AstNodeType[ExpressionNode | StatementNode | ErrorNode]
