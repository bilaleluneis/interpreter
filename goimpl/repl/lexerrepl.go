package repl

import (
	"bufio"
	"fmt"
	"goimpl/lexer"
	"goimpl/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		_, _ = fmt.Fprintf(out, PROMPT) // ignore error
		if scanned := scanner.Scan(); scanned {
			line := scanner.Text()
			if l, ok := lexer.New(line); ok {
				for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
					_, _ = fmt.Fprintf(out, "%+v\n", tok) // ignore error
				}
			}
		}
	}
}
