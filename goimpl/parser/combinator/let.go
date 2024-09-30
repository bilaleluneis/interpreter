package combinator

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
)

func Let[L lexer.LexerConstraint[L]](l L) Result[L] {
	if letTok := l.NextToken(); letTok.Type == token.LET {
		letStmnt := ast.Let{Tok: letTok}
		if identifier := l.NextToken(); identifier.Type == token.IDENTIFIER {
			letStmnt.Name = &ast.Identifier{Tok: identifier, Value: identifier.Literal}
			if l.NextToken().Type == token.ASSIGN {
				for l.NextToken().Type != token.SEMICOLON {
					// skip value for now and loop until we find a semicolon
					//FIXME: do I need to check for EOF here?
				}
				// got here then we have found a semicolon, return the let statement
				return Result[L]{l, letStmnt}
			}
		}
	}
	return Result[L]{l, ast.Error{Message: "Failed to parse let statement"}}
}
