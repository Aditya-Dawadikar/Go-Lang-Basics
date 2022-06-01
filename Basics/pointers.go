package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)

	iptr := &i
	fmt.Println(iptr, *iptr)
}
