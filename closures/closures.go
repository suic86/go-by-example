package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()
	for i := 0; i < 3; i++ {
		fmt.Println(nextInt())
	}
	newInts := intSeq()
	fmt.Println(newInts())
}
