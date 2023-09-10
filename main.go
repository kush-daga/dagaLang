package main

import (
	"fmt"
	"github.com/kush-daga/dagalang/repl"
	"os"
	"os/user"
)

func main() {
	currUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the DagaLang programming language!\n", currUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
