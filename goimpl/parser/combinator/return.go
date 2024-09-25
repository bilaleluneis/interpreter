package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

// Retrn FIXME: at the moment value is not captured, just skipping until we find a semicolon
func Retrn[L any, CL lexer.CopyableLexer[L]](l L) Result[L, CL] {
	retrnLexr := CL(&l)
	if retrnTok := retrnLexr.NextToken(); retrnTok.Type == token.RETURN {
		retrnStmnt := ast.Return{Tok: retrnTok}
		for retrnLexr.NextToken().Type != token.SEMICOLON {
			// FIXME: skip expression for now and loop until we find a semicolon
		}
		return Result[L, CL]{*retrnLexr, retrnStmnt}
	}
	return Result[L, CL]{l, ast.Error{Message: "Failed to parse return statement"}}
}
