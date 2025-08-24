package parser

import (
	"fmt"
	"goimpl/parser/internal"
	"goimpl/token"
	"strings"
)

type blockTestCase struct {
	tokens          []token.Token
	expectedBlock   string
	expectedErrrMsg string
}

var blockTests = map[string]blockTestCase{
	"invalid_block": {
		tokens: []token.Token{
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrrMsg: fmt.Sprintf(internal.BlockErrExpectedRBrace, token.EOF),
	},
	"empty_block": {
		tokens: []token.Token{
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedBlock: "{}",
	},
	"block_with_return_statement": {
		tokens: []token.Token{
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RETURN, Literal: "return"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedBlock: strings.ReplaceAll(
			`{
				return x;
			}`, "\t", ""),
	},
	"block_with_multiple_statements": {
		tokens: []token.Token{
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "1"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "y"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "2"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RETURN, Literal: "return"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.IDENTIFIER, Literal: "y"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedBlock: strings.ReplaceAll(`{
		let x = 1;
		let y = 2;
		return (x + y);
		}`, "\t", ""),
	},
}
