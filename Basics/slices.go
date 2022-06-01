package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s, len(s))

	s[0] = "aditya"
	s[1] = "prashant"
	s[2] = "dawadikar"
	fmt.Println(s)
	fmt.Println(len(s[1]))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	d := s[1:3]
	e := s[:3]
	fmt.Println(d, e)

	f := make([][]int, 3)
	for i := 0; i < 3; i++ {
		j := i + 1
		f[i] = make([]int, j)
		for j := 1; j <= i; j++ {
			f[i][j] = j
		}
	}
	fmt.Println(f)
}
