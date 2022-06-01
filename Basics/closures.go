package main

import "fmt"

func func1() func() int {
	a := 1
	return func() int {
		a += 1
		return a
	}
}

func main() {
	func1Proxy := func1()
	fmt.Println(func1Proxy())
	fmt.Println(func1Proxy())
	fmt.Println(func1Proxy())

	func2Proxy := func1()
	fmt.Println(func2Proxy())
	fmt.Println(func2Proxy())
	fmt.Println(func2Proxy())
}
