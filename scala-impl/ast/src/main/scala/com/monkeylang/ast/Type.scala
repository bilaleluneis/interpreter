package com.monkeylang.ast

sealed abstract class Type extends ExpressionNode:
	override def toString: String = "ast.Type"

class ValueType extends Type:
	override def toString: String = "ast.ValueType"

class referenceType extends Type:
	override def toString: String = "ast.ReferenceType"

class FunctionType(paramTypes: List[Type], returnType: Type) extends Type:
	override def toString: String =
		val params = paramTypes.map(_.toString).mkString(", ")
		s"ast.FunctionType([$params] => $returnType)"
