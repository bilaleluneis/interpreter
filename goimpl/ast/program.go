package ast

import (
	"fmt"
	"goimpl/debug/treedrawer/tree"
)

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
	out := "\n---- Program Start ----\n"
	for _, s := range p.Statements {
		out += "\n"
		out += s.String()
		out += "\n"
	}
	out += "\n---- Program End ----\n"
	return out
}

func (p *Program) Visualize() {
	t := tree.NewTree(tree.NodeString("Program"))
	for _, s := range p.Statements {
		switch stmtType := s.(type) {
		case *Let:
			visualizeLet(*stmtType, t)
		case *Return:
			visualizeReturn(*stmtType, t)
		case *ExpressionStatement:
			visualizeExpressionStatement(*stmtType, t)
		}
	}
	fmt.Println(t)
}
