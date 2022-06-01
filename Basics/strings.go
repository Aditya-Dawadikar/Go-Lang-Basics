package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello"

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// returns index and rune value
	// it will not print the character, hence we typecase it to string
	for index, runeval := range s {
		fmt.Println(index, runeval, string(runeval))
	}

}
