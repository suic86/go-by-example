package main

import (
	"fmt"
)

func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` must be called with a deferred function
	// the return value of `recover` is the error raised in the call to `panic`
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
