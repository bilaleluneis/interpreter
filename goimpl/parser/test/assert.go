package test

import (
	"goimpl/ast"
	"testing"
)

//FIXME: might not need this file and utils

func toExpression(stmt ast.Statement) ast.Expression {
	if ptr, ok := stmt.(*ast.ExpressionStatement); ok {
		return ptr.Exprssn
	} else if value, ok := stmt.(ast.ExpressionStatement); ok {
		return value.Exprssn
	}
	return nil
}

func toInfixExpression(expr ast.Expression) *ast.InfixExpression {
	if ptr, ok := expr.(*ast.InfixExpression); ok {
		return ptr
	} else if value, ok := expr.(ast.InfixExpression); ok {
		return &value
	}
	return nil
}

func toPrefixExpression(expr ast.Expression) *ast.PrefixExpression {
	if ptr, ok := expr.(*ast.PrefixExpression); ok {
		return ptr
	} else if value, ok := expr.(ast.PrefixExpression); ok {
		return &value
	}
	return nil
}

func toBoolean(expr ast.Expression) *ast.Boolean {
	if ptr, ok := expr.(*ast.Boolean); ok {
		return ptr
	} else if value, ok := expr.(ast.Boolean); ok {
		return &value
	}
	return nil
}

func toIntegerLiteral(expr ast.Expression) *ast.IntegerLiteral {
	if ptr, ok := expr.(*ast.IntegerLiteral); ok {
		return ptr
	} else if value, ok := expr.(ast.IntegerLiteral); ok {
		return &value
	}
	return nil
}

func toIdentifier(expr ast.Expression) *ast.Identifier {
	if ptr, ok := expr.(*ast.Identifier); ok {
		return ptr
	} else if value, ok := expr.(ast.Identifier); ok {
		return &value
	}
	return nil
}

func toReturn(stmt ast.Statement) *ast.Return {
	if ptr, ok := stmt.(*ast.Return); ok {
		return ptr
	} else if value, ok := stmt.(ast.Return); ok {
		return &value
	}
	return nil
}

func fail(tag string, t *testing.T, format string, args ...any) {
	t.Errorf(tag+"::"+format, args...)
}

func success(tag string, t *testing.T, format string, args ...any) {
	t.Logf(tag+"::"+format, args...)
}
