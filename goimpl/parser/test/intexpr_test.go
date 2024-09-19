package test

import (
	"fmt"
	"goimpl/ast"
	"goimpl/lexer"
	"goimpl/parser"
	"goimpl/parser/pratt"
	"goimpl/token"
	"testing"
)

func TestIntLiteralExpr(t *testing.T) {
	// 5;
	l := lexer.NewStubLexer([]token.Token{
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	})

	lxr := l.GetCopy()
	prattParser := pratt.New(&lxr)
	if err := parseIntLiteral(prattParser); err != nil {
		t.Error(err)
		t.FailNow() // stop the test here
	}

	//FIXME - implement combinator parser
	//lxr = l.GetCopy() // reset the lexer
	//combParser := combinator.New(lxr)
	//if err := parseIntLiteral(combParser); err != nil {
	//	t.Error(err)
	//	t.FailNow() // stop the test here
	//}
}

func parseIntLiteral(p parser.ParserType) error {
	program, ok := p.ParseProgram()
	if !ok {
		return fmt.Errorf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement) // stmt is of type *ast.ExpressionStatement
	if !ok {
		return fmt.Errorf("statement is not ExpressionStatement. got=%T", program.Statements[0])
	}
	integ, ok := stmt.Exprssn.(*ast.IntegerLiteral) // expression statement is of type *ast.IntegerLiteral
	if !ok {
		return fmt.Errorf("exp not IntegerLiteral. got=%T", stmt.Exprssn)
	}
	if integ.Value != 5 {
		return fmt.Errorf("integ.Value not %d. got=%d", 5, integ.Value)
	}
	tokenLit := integ.TokenLiteral()
	if tokenLit != "5" {
		return fmt.Errorf("tok Literal not %s. got=%s", "5", tokenLit)
	}
	return nil
}
