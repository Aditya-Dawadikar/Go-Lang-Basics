package main

import "fmt"

func forOne() {
	var i = 1
	for i < 10 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println()
}

func forTwo() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()
}

func forThree() {
	var i = 1
	for {
		fmt.Println(i)
		i += 1
		if i > 10 {
			break
		}
	}
	fmt.Println()
}

func main() {
	forOne()
	forTwo()
	forThree()
}
