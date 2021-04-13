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
		panic(err)
	}

	fmt.Printf("Hello %s! this is the Monkey programming language!\n", user.Username)
	fmt.Printf("Starting in REPL mode\n")
	repl.Start(os.Stdin, os.Stdout)
}
