package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "aditya", age: 22}
	fmt.Println(p)

	ptr := &p
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// omitted fields are zero valued
	p2 := person{age: 22}
	fmt.Println(p2)

	fmt.Println(p.name, reflect.TypeOf(p.name), p.age, reflect.TypeOf(p.age))
	fmt.Println(p2.name, reflect.TypeOf(p2.name), p2.age, reflect.TypeOf(p2.age))
}
