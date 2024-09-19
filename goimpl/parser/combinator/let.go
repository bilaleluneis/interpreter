package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

func Let[L any, CL lexer.CopyableLexer[L]](l L) Result[L, CL] {
	letLexr := CL(&l)
	if letTok := letLexr.NextToken(); letTok.Type == token.LET {
		letStmnt := ast.Let{Tok: letTok}
		if identifier := letLexr.NextToken(); identifier.Type == token.IDENTIFIER {
			letStmnt.Name = &ast.Identifier{Tok: identifier, Value: identifier.Literal}
			if letLexr.NextToken().Type == token.ASSIGN {
				for letLexr.NextToken().Type != token.SEMICOLON {
					// skip value for now and loop until we find a semicolon
					//FIXME: do I need to check for EOF here?
				}
				// got here then we have found a semicolon, return the let statement
				return Result[L, CL]{*letLexr, letStmnt}
			}
		}
	}
	return Result[L, CL]{l, ast.Error{Message: "Failed to parse let statement"}}
}
