package main

import (
	"fmt"
	"math"
)

// `const` statement can appear anywhere where a `var` statemet can appear
const s string = "constant"

func main() {
	fmt.Println(s)

	// A numeric constant has no type unitl it's given one, such as by explicit conversion
	const n = 500_000_000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
