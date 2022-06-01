package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	a := 10
	if a < 10 {
		fmt.Println("a < 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a = 10")
	}

	switch a {
	case 9:
		fmt.Println("9")
	case 10:
		fmt.Println("10")
	case 11:
		fmt.Println("11")
	default:
		fmt.Println("default")
	}
}
