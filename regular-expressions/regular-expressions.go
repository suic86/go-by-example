package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	pl := fmt.Println
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	pl(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	pl(r.MatchString("peach"))

	pl(r.FindString("peach punch"))
	pl("idx:", r.FindStringIndex("peach punch"))
	pl(r.FindStringSubmatch("peach punch"))
	pl(r.FindStringSubmatchIndex("peach punch"))
	pl(r.FindAllString("peach punch pinch", -1))
	pl(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	pl("regexp:", r)

	pl(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	pl(string(out))

}
