package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"strconv"
)

func prefix[L lexer.LexerConstraint[L]](l L) (ast.PrefixExpression, bool) {
	if prfix := l.NextToken(); prfix.Type == token.MINUS || prfix.Type == token.BANG {
		return ast.PrefixExpression{Tok: prfix, Operator: prfix.Literal}, true
	}
	return ast.PrefixExpression{}, false
}

func PrefixInt[L lexer.LexerConstraint[L]](l L) Result[L] {
	if prefix, ok := prefix(l); ok {
		if intLiteral := l.NextToken(); intLiteral.Type == token.INT {
			if intVal, err := strconv.ParseInt(intLiteral.Literal, 0, 64); err == nil {
				intExprssion := ast.IntegerLiteral{Tok: intLiteral, Value: intVal}
				prefix.Right = intExprssion
				return Result[L]{l, ast.ExpressionStatement{Exprssn: prefix}}
			}
		}
	}
	return Result[L]{l, ast.Error{Message: "Failed to parse integer prefix expression"}}
}

func InfixInt[L lexer.LexerConstraint[L]](l L) Result[L] {
	if leftExpr := IntExpr(l); !leftExpr.IsError() {
		infixExpr := ast.InfixExpression{Left: toExpression(leftExpr.stmnt)}
		if infix := l.NextToken(); infix.Type == token.PLUS || infix.Type == token.MINUS {
			infixExpr.Operator = infix.Literal
			if rightExpr := IntExpr(l); !rightExpr.IsError() {
				infixExpr.Right = toExpression(rightExpr.stmnt)
				return Result[L]{l, ast.ExpressionStatement{Exprssn: infixExpr}}
			}
		}
	}
	return Result[L]{l, ast.Error{Message: "Failed to parse infix expression"}}
}

// TODO: make part of Result struct , to convert to ast.ExpressionStatement
func toExpression(stmnt ast.Statement) ast.Expression {
	if expr, ok := stmnt.(ast.ExpressionStatement); ok {
		return expr.Exprssn
	}
	return nil
}
