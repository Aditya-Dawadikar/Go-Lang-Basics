package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func read() {
	dat, err := os.ReadFile("./data.txt")
	checkError(err)
	fmt.Println(string(dat))

	// f: file object reference
	f, err := os.Open("./data.txt")
	checkError(err)
	fmt.Println(f, reflect.TypeOf(f))

	// reading first 5 bytes
	a := make([]byte, 5)
	b, err := f.Read(a)
	checkError(err)
	fmt.Println(string(a[:b]))

	// seek to a particular location eg: 6th character
	c, err := f.Seek(6, 0)
	checkError(err)
	fmt.Println(c)
	d := make([]byte, 10)
	e, err := f.Read(d)
	checkError(err)
	fmt.Println(string(d[:e]))

	// rewinding back to start of the file
	_, err1 := f.Seek(0, 0)
	checkError(err1)

	// reading with buffer
	h := bufio.NewReader(f)
	i, err := h.Peek(5)
	checkError(err)
	fmt.Println(string(i))
}

func write() {
	a := []byte("hello world\nI love go\n")
	err := os.WriteFile("./output1.txt", a, 0644)
	checkError(err)

	// writing to file
	f, err := os.Create("output2.txt")
	checkError(err)
	b := []byte{115, 111, 109, 101, 10}
	c, err := f.Write(b)
	checkError(err)
	fmt.Println(c)
}

func main() {
	// read()
	write()
}
