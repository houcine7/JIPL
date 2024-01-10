package main

import (
	"fmt"
	"os"
	"os/user"

	repl "github.com/houcine7/JIPL/cmd/REPL"
)

func main() {
	currUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!, welcome to JIPL happy coding :)", currUser.Username)
	fmt.Printf("Start typing JIPL code ...\n")
	repl.Start(os.Stdin, os.Stdout)
}
