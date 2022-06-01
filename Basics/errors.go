package main

import (
	"errors"
	"fmt"
)

func errorFunc() error {
	return errors.New("Intentionally raised")
}

type ErrorType struct {
	errorCode   int
	errorString string
}

func (e *ErrorType) RandomError() string {
	return fmt.Sprintf("Error code: %d\n Error message:%s", e.errorCode, e.errorString)
}

func main() {
	err := errorFunc()
	fmt.Println(err)

	err_1 := ErrorType{errorCode: -1, errorString: "random error"}
	fmt.Println(err_1.RandomError())
}
