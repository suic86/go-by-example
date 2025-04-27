package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Printf("argsWithProg: %v\n", argsWithProg)
	fmt.Printf("argsWithoutProg: %v\n", argsWithoutProg)
	fmt.Println("Command: `./command-line-arguments a b c d`")
	fmt.Printf("3rd argument: %v\n", arg)
}
