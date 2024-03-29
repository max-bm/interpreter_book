package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		// panic stops execution of the current goroutine
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	// Take input to the repl from stdin, and write output to stdout
	repl.Start(os.Stdin, os.Stdout)
}
