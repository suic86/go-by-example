package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	pf := fmt.Printf
	// print struct
	pf("struct1: %v\n", p)
	// print struct with field names
	pf("struct2: %+v\n", p)
	// print struct's representation
	// i.e. the source code snipped that would produce that value
	pf("struct3: %#v\n", p)

	// print the type of the value
	pf("type: %T\n", p)

	pf("bool: %t\n", true)

	// print integers

	// decimal format
	pf("int: %d\n", 123)
	// binary format
	pf("bin: %b\n", 14)
	// character corresponding to integer
	pf("char %c\n", 33)
	// hexadecimal representation
	pf("hex: %x\n", 456)

	// print floats

	// basic formatting
	pf("float1: %f\n", 78.9)
	// scientific notations
	pf("float2: %e\n", 123400000.0)
	pf("float3: %E\n", 123400000.0)

	// string formatting

	// basic
	pf("str1: %s\n", "\"string\"")
	// quoted
	pf("str2: %q\n", "\"string\"")
	// hexadecimal representation
	pf("str3: %x\n", "hex this")

	// print pointer
	pf("pointer: %p\n", &p)

	// print right (w/o -) and left (with -) text
	pf("width1: |%6d|%6d|\n", 12, 345)
	pf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	pf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	pf("width4: |%6s|%6s|\n", "foo", "b")
	pf("width5: |%-6s|%-6s|\n", "foo", "b")

	// returns a string with the formatted output without printing it
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// prints the results into an instance of `io.Writers`
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
