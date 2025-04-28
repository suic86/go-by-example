package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	// standard logger - unstructured output
	// it outputs to `os.Stderr` by default
	log.Println("standard logger")

	// show microsecond
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// show file
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// custom logger with custom prefix and target
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// method to change prefix of an existing logger (including the default one)
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// log can write into any `io.Writer`
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Println("from buflog:", buf.String())

	// log/slog stands for structured logging
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
}
