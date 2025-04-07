package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func factNonRec(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(factNonRec(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
