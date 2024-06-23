package main

import (
	"fmt"
	"goimpl/repl"
	"os"
	"os/user"
)

func main() {
	if usr, err := user.Current(); err == nil {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", usr.Username)
		repl.Start(os.Stdin, os.Stdout)
	}
}
