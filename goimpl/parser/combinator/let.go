package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

func Let[L any, CL lexer.CopyableLexer[L]](l L) (ast.Statement, L) {
	failure, currLexer := Fail[L, CL](CL(&l).GetCopy()) // on failure return this
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
		return letStmnt, letLexr.GetCopy()
	}
	return failure, currLexer
}
