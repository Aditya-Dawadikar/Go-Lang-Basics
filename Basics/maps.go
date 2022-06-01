package main

import "fmt"

func main() {
	var a = make(map[string]string)

	a["name"] = "aditya"
	a["age"] = "22"
	a["phone"] = "98xxxxxxxx"

	fmt.Println(a)
	v1, v2 := a["name"]
	v3, v4 := a["name1"]
	fmt.Println(v1, v2)
	fmt.Println(v3, v4)

	delete(a, "phone")
	fmt.Println(a)

	b := map[string]int{"abc": 1, "def": 2}
	fmt.Println(b)
}
