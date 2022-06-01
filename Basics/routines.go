package main

import (
	"fmt"
	"sync"
)

func one(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println(1)
}
func two(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println(2)
}
func three(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println(3)
}

func normalOrder() {
	one(nil)
	two(nil)
	three(nil)
}

func parallelOrder(wg *sync.WaitGroup) {
	go one(wg)
	go two(wg)
	go three(wg)
}

func main() {
	var wg = new(sync.WaitGroup)
	wg.Add(3)

	fmt.Println("Function call without go routine")
	normalOrder()

	fmt.Println("Function call with go routine")
	parallelOrder(wg)
	wg.Wait()
}
