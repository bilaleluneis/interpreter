package ast

import (
	"goimpl/token"
	"testing"
)

func TestProgram(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			simpleLet,
			letSimpleExpression,
			letComplexExpression,
			expressionStatement,
		},
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got=%d", len(program.Statements))
	}
	t.Logf("Program: %s", program)
	t.Logf("Program Dump: %s", program.Dump(1))
}

// let myVar = 5;
var simpleLet Statement = &Let{
	Tok: token.Token{Type: token.LET, Literal: "let"},
	Name: &Identifier{
		Tok:   token.Token{Type: token.IDENTIFIER, Literal: "x"},
		Value: "myVar",
	},
	Value: &IntegerLiteral{
		Tok:   token.Token{Type: token.INT, Literal: "5"},
		Value: 5,
	},
}

// let myVar = 5 + 10;
var letSimpleExpression Statement = &Let{
	Tok: token.Token{Type: token.LET, Literal: "let"},
	Name: &Identifier{
		Tok:   token.Token{Type: token.IDENTIFIER, Literal: "z"},
		Value: "z",
	},
	Value: &InfixExpression{
		Left: &IntegerLiteral{
			Tok:   token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
		Tok:      token.Token{Type: token.PLUS, Literal: "+"},
		Operator: "+",
		Right: &IntegerLiteral{
			Tok: token.Token{Type: token.INT, Literal: "10"},
		},
	},
}

// let y = 5 + 10 * 10;
var letComplexExpression Statement = &Let{
	Tok: token.Token{Type: token.LET, Literal: "let"},
	Name: &Identifier{
		Tok:   token.Token{Type: token.IDENTIFIER, Literal: "y"},
		Value: "y",
	},
	Value: &InfixExpression{
		Left: &IntegerLiteral{
			Tok:   token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
		Tok:      token.Token{Type: token.PLUS, Literal: "+"},
		Operator: "+",
		Right: &InfixExpression{
			Left: &IntegerLiteral{
				Tok:   token.Token{Type: token.INT, Literal: "10"},
				Value: 10,
			},
			Tok:      token.Token{Type: token.ASTER, Literal: "*"},
			Operator: "*",
			Right: &IntegerLiteral{
				Tok:   token.Token{Type: token.INT, Literal: "10"},
				Value: 10,
			},
		},
	},
}

// 5 + 10;
var expressionStatement Statement = &ExpressionStatement{
	Tok: token.Token{Type: token.INT, Literal: "5"},
	Exprssn: &InfixExpression{
		Left: &IntegerLiteral{
			Tok:   token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
		Tok:      token.Token{Type: token.PLUS, Literal: "+"},
		Operator: "+",
		Right: &IntegerLiteral{
			Tok: token.Token{Type: token.INT, Literal: "10"},
		},
	},
}
