package main

import (
	"fmt"
	"reflect"
	str "strings"
)

func main() {
	s1 := "hello"
	s2 := "world"

	fmt.Println(s1 + " " + s2)

	// frequency count works on a substring
	fmt.Println(str.Count(s1, "l"))
	fmt.Println(str.Count(s1, "he"))

	// prefix and suffix
	fmt.Println(str.HasPrefix(s1, "he"), str.HasSuffix(s1, "ol"))

	// generating tokens using string split
	tokens := str.Split("hello world go lang", " ")
	fmt.Println(tokens, reflect.TypeOf(tokens))

	// string repeatition
	fmt.Println(str.Repeat("a", 11))

	// string join using a substring
	fmt.Println(str.Join([]string{s1, s2}, " "))

	// index of substring
	fmt.Println(str.Index(s1, "el"))

	// check if substring exists
	fmt.Println(str.Contains(s1, "wallah"))

	// replacing substring
	fmt.Println(str.Replace("helelelello", "el", s2, 1))
	fmt.Println(str.Replace("helelelello", "el", s2, 0))
	fmt.Println(str.Replace("helelelello", "el", s2, -1))
}
