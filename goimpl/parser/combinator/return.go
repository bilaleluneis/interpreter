package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

// Retrn FIXME: at the moment value is not captured, just skipping until we find a semicolon
func Retrn[L lexer.LexerConstraint[L]](l L) Result[L] {
	if retrnTok := l.NextToken(); retrnTok.Type == token.RETURN {
		retrnStmnt := ast.Return{Tok: retrnTok}
		for l.NextToken().Type != token.SEMICOLON {
			// FIXME: skip expression for now and loop until we find a semicolon
		}
		return Result[L]{l, retrnStmnt}
	}
	return Result[L]{l, ast.Error{Message: "Failed to parse return statement"}}
}
