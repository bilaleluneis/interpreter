package parser

import (
	"fmt"
	"goimpl/parser/internal"
	"goimpl/token"
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
		expectedBlock: "{return x;}",
	},
}
