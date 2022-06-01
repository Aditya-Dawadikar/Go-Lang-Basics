package main

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height int
}

type circle struct {
	radius int
}

func (r rect) area() float64 {
	return float64(r.height * r.width)
}

func (r circle) area() float64 {
	return math.Pi * math.Pow(float64(r.radius), 2)
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}
	fmt.Println(r.area(), reflect.TypeOf(r.area()))
	fmt.Println(c.area(), reflect.TypeOf(c.area()))
}
