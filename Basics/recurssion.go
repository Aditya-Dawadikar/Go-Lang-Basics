package main

import "fmt"

func f(i int) {
	fmt.Println(i)
	if i < 5 {
		i += 1
		f(i)
	}
}

func main() {
	f(0)
}
