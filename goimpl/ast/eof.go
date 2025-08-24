package ast

type Eof struct{}

func (Eof) statmentNode()        {}
func (Eof) TokenLiteral() string { return "EOF" }
func (Eof) String() string       { return "EOF" }
