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
			assignmetTok := letLexr.NextToken()
			if assignmetTok.Type == token.ASSIGN {
				for letLexr.NextToken().Type != token.SEMICOLON {
					// skip value for now and loop until we find a semicolon
				}
			}
		}
		return Result[L, CL]{*letLexr, letStmnt}
	}
	return Result[L, CL]{l, &ast.Error{Message: "Failed to parse let statement"}}
}
