package ast

type Eof struct{}

func (Eof) statementNode()       {}
func (Eof) TokenLiteral() string { return "EOF" }
func (Eof) String() string       { return "EOF" }
