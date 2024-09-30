package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"strconv"
)

func IntExpr[L lexer.LexerConstraint[L]](l L) Result[L] {
	if intTok := l.NextToken(); intTok.Type == token.INT {
		intLiteral := ast.IntegerLiteral{Tok: intTok}
		if v, e := strconv.ParseInt(intTok.Literal, 0, 64); e == nil {
			intLiteral.Value = v
			return Result[L]{l, ast.ExpressionStatement{Exprssn: &intLiteral}}
		}
	}
	return Result[L]{l, ast.Error{Message: "Failed to parse integer expression"}}
}

func IdentifierExpr[L lexer.LexerConstraint[L]](l L) Result[L] {
	if identTok := l.NextToken(); identTok.Type == token.IDENTIFIER {
		ident := ast.Identifier{Tok: identTok, Value: identTok.Literal}
		return Result[L]{l, ast.ExpressionStatement{Exprssn: &ident}}
	}
	return Result[L]{l, ast.Error{Message: "Failed to parse identifier expression"}}
}
