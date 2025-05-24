package main

import (
	"divakaivan/lang-interpreter-go/repl"
	"flag"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")

	showAST := flag.Bool("ast", false, "print the AST")
	flag.Parse()

	repl.Start(os.Stdin, os.Stdout, *showAST)
}
