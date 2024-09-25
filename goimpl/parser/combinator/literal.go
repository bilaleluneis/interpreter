package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"strconv"
)

func IntExpr[L any, CL lexer.CopyableLexer[L]](l L) Result[L, CL] {
	lxr := CL(&l)
	if intTok := lxr.NextToken(); intTok.Type == token.INT {
		intLiteral := ast.IntegerLiteral{Tok: intTok}
		if v, e := strconv.ParseInt(intTok.Literal, 0, 64); e == nil {
			intLiteral.Value = v
			return Result[L, CL]{*lxr, ast.ExpressionStatement{Exprssn: &intLiteral}}
		}
	}
	return Result[L, CL]{l, ast.Error{Message: "Failed to parse integer expression"}}
}

func IdentifierExpr[L any, CL lexer.CopyableLexer[L]](l L) Result[L, CL] {
	lxr := CL(&l)
	if identTok := lxr.NextToken(); identTok.Type == token.IDENTIFIER {
		ident := ast.Identifier{Tok: identTok, Value: identTok.Literal}
		return Result[L, CL]{*lxr, ast.ExpressionStatement{Exprssn: &ident}}
	}
	return Result[L, CL]{l, ast.Error{Message: "Failed to parse identifier expression"}}
}
