package main

import "fmt"

func func1() {
	fmt.Println("hello world")
}

func func2() string {
	return "hello"
}

func func3() (string, string) {
	return "hello", "world"
}

func main() {
	func1()
	fmt.Println(func2())
	fmt.Println(func3())
}
