package lexer

import "testing"

func TestLexerCopy(t *testing.T) {
	lexer := NewLazyLexer("let five = 5;")
	lexerCopy := CopyOf(&lexer)
	lexerCopy.NextToken() // let
	if lexerCopy.NextToken().Literal != "five" {
		t.Fatal("lexerCopy shuld have returned 'five'")
	}
	if lexer.NextToken().Literal != "let" {
		t.Fatal("lexerCopy should not affect original lexer")
	}
	if &lexer == lexerCopy {
		t.Fatal("lexer and lexerCopy should be different instances")
	}
}
