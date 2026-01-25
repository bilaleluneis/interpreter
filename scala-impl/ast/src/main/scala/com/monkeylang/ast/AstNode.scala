// ...existing code from original AstNode.scala will be moved here...
package com.monkeylang.ast

sealed trait AstNode

trait ExpressionNode extends AstNode
trait StatementNode extends AstNode
trait ErrorNode extends AstNode

enum AstNodeKind[+A <: AstNode]:
	case ERROR(err: ErrorNode)
	case EXPRESSION(expr: ExpressionNode)
	case STATEMENT(stmt: StatementNode)

	override def toString: String =
		this match
			case ERROR(err)       => s"AstNodeKind.ERROR($err)"
			case EXPRESSION(expr) => s"AstNodeKind.EXPRESSION($expr)"
			case STATEMENT(stmt)  => s"AstNodeKind.STATEMENT($stmt)"

enum OperatorKind extends ExpressionNode:
	case UNARY(up: UnaryOperator)
	case BINARY(bp: BinaryOperator)

	override def toString: String =
		this match
			case UNARY(up)  => s"OperatorKind.UNARY($up)"
			case BINARY(bp) => s"OperatorKind.BINARY($bp)"

enum TypeKind extends ExpressionNode:
	case VALUE_TYPE(vt: ValueType)
	case REFERENCE_TYPE(rt: referenceType)
	case FUNCTION_TYPE(ft: FunctionType)

	override def toString: String =
		this match
			case VALUE_TYPE(vt)     => s"TypeKind.VALUE_TYPE($vt)"
			case REFERENCE_TYPE(rt) => s"TypeKind.REFERENCE_TYPE($rt)"
			case FUNCTION_TYPE(ft)  => s"TypeKind.FUNCTION_TYPE($ft)"

type AstType = AstNodeKind[ExpressionNode | StatementNode | ErrorNode]
