package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 10
	var b = "hello"
	var c = 'c'
	var d = 10.10
	var e = true

	const f = 50000000000
	const g = 3e20 / f

	fmt.Println(a, b, c, d, e)

	fmt.Println(g, int64(g), math.Sin(f))
}
