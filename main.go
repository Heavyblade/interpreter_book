package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey programming language!\n", user.Name)
	fmt.Println("Feel free to type some commands")

	repl.Start(os.Stdin, os.Stdout)
}
