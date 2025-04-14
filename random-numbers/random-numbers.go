package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	p := fmt.Print
	pl := fmt.Println

	p(rand.IntN(100), ",")
	p(rand.IntN(100))
	pl()

	pl(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	pl()

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)

	p(r2.IntN(100), ",")
	p(r2.IntN(100))
	pl()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)

	p(r3.IntN(100), ",")
	p(r3.IntN(100))
	pl()

}
