package test

import (
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

var booleanTests = []lexerToAstResult{
	{
		lxr: lexer.NewStubLexer([]token.Token{
			{Type: token.TRUE, Literal: "true"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		}),
		expectedAst: &ast.ExpressionStatement{
			Tok:     token.Token{Type: token.TRUE, Literal: "true"},
			Exprssn: &ast.Boolean{Tok: token.Token{Type: token.TRUE, Literal: "true"}, Value: true},
		},
		expectedProgram: "true",
	},
}

// FIXME: combinator parser not used in this test
func TestBooleanExpression(t *testing.T) {
	for _, tt := range booleanTests {
		lxrUnderTest := &tt.lxr
		for pname, parser := range testParsers(lxrUnderTest).initPratt().parsers {
			t.Logf("testing parser %s", pname)
			tt.test(t, parser)
		}
	}
}
