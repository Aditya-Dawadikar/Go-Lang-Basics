package main

import "fmt"

type abc struct {
	num int
}

// abc struct is embedded inside def struct
type def struct {
	abc
	x int
}

func main() {
	a := def{abc: abc{num: 10}, x: 10}
	fmt.Println(a)
	fmt.Println(a.abc, a.abc.num, a.x)
}
