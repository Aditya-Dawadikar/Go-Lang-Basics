package main

import (
	"fmt"
	"time"
)

func someFunc(done chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("still running")
	}
	done <- true
}
func main() {
	done := make(chan bool)

	go someFunc(done)
	d := <-done
	fmt.Println("final value of d: ", d)

}
