package pratt

import (
	"fmt"
)

func printErrs(p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	fmt.Printf("parser has %d error(s)\n", len(errors))
	for _, msg := range errors {
		fmt.Printf("parser error: %q\n", msg)
	}
}
