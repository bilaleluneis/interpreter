package test

import (
	"goimpl/ast"
	"testing"
)

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

//lint:ignore U1000 will be used soon
func toPrefixExpression(expr ast.Expression) *ast.PrefixExpression {
	if ptr, ok := expr.(*ast.PrefixExpression); ok {
		return ptr
	} else if value, ok := expr.(ast.PrefixExpression); ok {
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

func fail(tag string, t *testing.T, format string, args ...any) {
	t.Errorf(tag+"::"+format, args...)
}

//lint:ignore U1000 false positive
func isExpressionStmt(stmt ast.Statement) bool {
	_, isExprPtr := stmt.(*ast.ExpressionStatement)
	_, isExpr := stmt.(ast.ExpressionStatement)
	return isExpr || isExprPtr
}

//lint:ignore U1000 false positive
func isIntLiteral(stmt ast.Statement) (string, bool) {
	var isIntLit bool
	var literal string
	var intAst *ast.IntegerLiteral
	if ptr, ok := stmt.(*ast.ExpressionStatement); ok {
		intAst, isIntLit = ptr.Exprssn.(*ast.IntegerLiteral)
		literal = intAst.TokenLiteral()
	} else if value, ok := stmt.(ast.ExpressionStatement); ok {
		intAst, isIntLit = value.Exprssn.(*ast.IntegerLiteral)
		literal = intAst.TokenLiteral()
	}
	return literal, isIntLit
}

//lint:ignore U1000 false positive
func isIdentifier(stmt ast.Statement) (string, bool) {
	var isIdent bool
	var literal string
	var ident *ast.Identifier
	if ptr, ok := stmt.(*ast.ExpressionStatement); ok {
		ident, isIdent = ptr.Exprssn.(*ast.Identifier)
		literal = ident.TokenLiteral()
	} else if value, ok := stmt.(ast.ExpressionStatement); ok {
		ident, isIdent = value.Exprssn.(*ast.Identifier)
		literal = ident.TokenLiteral()
	}
	return literal, isIdent
}

//lint:ignore U1000 false positive
func isReturn(stmt ast.Statement) (ast.Expression, bool) {
	if ptr, ok := stmt.(*ast.Return); ok {
		return ptr.Value, true
	} else if value, ok := stmt.(ast.Return); ok {
		return value.Value, true
	}
	return nil, false
}

//lint:ignore U1000 false positive
func isPrefixExpression(stmt ast.Statement) (*ast.PrefixExpression, bool) {
	if ptr, ok := stmt.(*ast.ExpressionStatement); ok {
		prefix, isPrefix := ptr.Exprssn.(*ast.PrefixExpression)
		return prefix, isPrefix
	} else if value, ok := stmt.(ast.ExpressionStatement); ok {
		prefix, isPrefix := value.Exprssn.(ast.PrefixExpression)
		return &prefix, isPrefix
	}
	return nil, false
}
