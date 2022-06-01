package main

import "fmt"

type rect struct {
	width, height int
}

// the function needs a receiver of type rect
func (r rect) area() int {
	return r.width * r.height
}

// receiver can be a pointer
func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	rptr := &r
	fmt.Println(rptr.area())
	fmt.Println(rptr.perimeter())
}
