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

func ParsePrefixInt[L lexer.LexerConstraint[L]](l L) Result[L] {
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
