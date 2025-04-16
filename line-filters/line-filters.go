package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// convert every line to upper case
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// check errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(0)
	}
}
