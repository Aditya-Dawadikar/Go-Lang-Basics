package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	for i := 1; i <= 5; i++ {
		a[i-1] = i
	}
	fmt.Println(a)

	var b = [5]int{6, 7, 8, 9, 10}
	fmt.Println(b)

	var c [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			c[i][j] = 10
		}
	}
	fmt.Println(c)
	fmt.Println(len(c), len(c[1]))
}
