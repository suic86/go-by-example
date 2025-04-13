package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// printm marshals the give object `d` and ant prints the marshalled data as string.
func printm(d any) {
	m, _ := json.Marshal(d)
	fmt.Println(string(m))
}

func main() {
	// marshall atomic values
	printm(true)
	printm(1)
	printm(2.34)
	printm("gopher")

	// marshall collections
	printm([]string{"apple", "peach", "pear"})
	printm(map[string]int{"apple": 5, "lettuce": 7})

	// `json` package can automacitally encode custom data types
	// it will include only exported fields
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	printm(res1D)

	// tags on struct can customize the encoded JSON key names
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	printm(res2D)

	// unmarshall data from a byte string
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]any
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	// json can be decoded into custom data types
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// json representations can be directly streamed
	// into instances of `os.Writers`.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// json representations can be directly read
	// from instances of `os.Readers`
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
