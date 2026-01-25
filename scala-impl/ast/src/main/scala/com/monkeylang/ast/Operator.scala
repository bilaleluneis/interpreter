package com.monkeylang.ast

enum OperatorPrecedence:
	case LOWEST
	case EQUALS // ==
	case LESSGREATER // > or <
	case SUM // +
	case PRODUCT // *
	case PREFIX // -X or !X
	case CALL // myFunction(X)

sealed trait Operator extends ExpressionNode:
	val symbl: String
	val precd: OperatorPrecedence

trait UnaryOperator(val Operand: ExpressionNode) extends Operator: // e.g., -X or !X
	final override def toString: String = s"ast.UnaryOperator($symbl $Operand)"

trait BinaryOperator( // e.g., X + Y
		val LeftOperand: ExpressionNode,
		val RightOperand: ExpressionNode
) extends Operator:
	final override def toString: String =
		s"ast.BinaryOperator($LeftOperand $symbl $RightOperand)"

// the only operator that will be implemented
// as part of the Monkey language interpreter
// assignment operator "="
final case class Assign(
		leftOperand: ExpressionNode,
		rightOperand: ExpressionNode
) extends BinaryOperator(leftOperand, rightOperand):
	override val symbl = "="
	override val precd = OperatorPrecedence.LOWEST
