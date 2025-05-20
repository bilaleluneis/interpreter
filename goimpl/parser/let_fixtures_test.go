package parser

import (
	"fmt"
	"goimpl/parser/internal"
	"goimpl/token"
)

type letTestCase struct {
	tokens   []token.Token
	wantErr  bool
	errorMsg string // expected error message when wantErr is true
	expected string // expected string representation for valid cases
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
		errorMsg: fmt.Sprintf(internal.LetErrExpectedIdentifier, token.ASSIGN),
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
		errorMsg: fmt.Sprintf(internal.LetErrExpectedAssign, token.INT),
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
		errorMsg: fmt.Sprintf(internal.LetErrExpectedSemicolon, token.EOF),
	},
}
