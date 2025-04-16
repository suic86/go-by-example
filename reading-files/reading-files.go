package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) == 2 {
		fmt.Errorf("invalid number of arguments %d instead of 2", len(os.Args))
	}
	fName := os.Args[1]

	// slurping file's entire content into memory
	dat, err := os.ReadFile(fName)
	check(err)
	fmt.Print(string(dat))

	// opens a file but does not read it
	f, err := os.Open(fName)
	check(err)
	defer f.Close()

	// read 5 bytes from the file into a byte slice
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek to a known location and read from there
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// seek relative to the cursor
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	_, err = f.Seek(-4, io.SeekCurrent)
	check(err)

	// read at least n bytes from a location
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// rewind to the start (no built in function for this)
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// buffered reader implementation in the standard library
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}
