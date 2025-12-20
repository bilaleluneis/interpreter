package ast

import "strings"

// Program is the top most node that will contain child statments
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out strings.Builder
	out.WriteString("\n---- Program Start ----\n")
	for _, s := range p.Statements {
		out.WriteString("\n")
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	out.WriteString("\n---- Program End ----\n")
	return out.String()
}
