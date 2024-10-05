package test

import "goimpl/ast"

func isExpressionStmt(stmt ast.Statement) bool {
	_, isExprPtr := stmt.(*ast.ExpressionStatement)
	_, isExpr := stmt.(ast.ExpressionStatement)
	return isExpr || isExprPtr
}

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

func isReturn(stmt ast.Statement) (ast.Expression, bool) {
	if ptr, ok := stmt.(*ast.Return); ok {
		return ptr.Value, true
	} else if value, ok := stmt.(ast.Return); ok {
		return value.Value, true
	}
	return nil, false
}

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
