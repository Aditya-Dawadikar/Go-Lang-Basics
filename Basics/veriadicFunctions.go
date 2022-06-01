package main

import "fmt"

func veriadicFunc(nums ...int) {
	fmt.Println(nums)

	total := 0
	for a, num := range nums {
		fmt.Println(a, num)
		total += num
	}
	fmt.Println(total)
}

func main() {
	nums := []int{1, 2, 3, 4}
	veriadicFunc(nums...)
}
