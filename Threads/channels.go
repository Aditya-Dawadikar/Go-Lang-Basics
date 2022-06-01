package main

import (
	"fmt"
)

func assign(message chan string) {
	message <- "hello world"
}

func assign2(message chan string) {
	message <- "hello duniya"
}

func main() {
	message := make(chan string)

	go assign(message)
	msg := <-message
	fmt.Println(msg)

	go assign2(message)
	updated_msg := <-message
	fmt.Println(updated_msg)
}
