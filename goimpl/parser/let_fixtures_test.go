package parser

import "goimpl/token"

type letTestCase struct {
	tokens   []token.Token
	wantErr  bool
	checkVal bool
	expected string
}

var letStatementTests = map[string]letTestCase{
	"simple_let": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		wantErr:  false,
		checkVal: true,
		expected: "let x = 5;",
	},
	"missing_identifier": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		wantErr:  true,
		checkVal: false,
	},
	"missing_assign": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		wantErr:  true,
		checkVal: false,
	},
	"missing_semicolon": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.EOF, Literal: ""},
		},
		wantErr:  true,
		checkVal: false,
	},
}
