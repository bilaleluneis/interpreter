package combinator

import (
	"goimpl/ast"
	"goimpl/parser"
)

type CombinatorParser func(...parser.ParserType) parser.ParserType

func (fp CombinatorParser) ParseProgram() *ast.Program {
	return fp.ParseProgram()
}
