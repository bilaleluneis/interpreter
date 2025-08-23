package internal

const (
	LetErrExpectedIdentifier  = "parse let: expected IDENTIFIER, got %s"
	LetErrExpectedAssign      = "parse let: expected ASSIGN, got %s"
	LetErrExpectedExpression  = "parse let: expected expression, got %s"
	LetErrExpectedSemicolon   = "parse let: expected semicolon, got %s"
	ErrExpectedPrefixParseFn  = "parse expression: expected prefix parse function for token type %s"
	ErrExpectedIntegerLiteral = "invalid integer literal: %s"
	ErrExpectedOpenPren       = "parse expression: expected (, got %s"
	ErrExpectedClosePren      = "parse expression: expected ), got %s"
	ErrExpectedOpenBrace      = "parse block: expected {, got %s"
	BlockErrExpectedRBrace    = "parse block: expected }, got %s"
)
