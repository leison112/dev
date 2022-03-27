package main

import (
	"github.com/cockroachdb/errors"
	"fmt"
)

func errorFunc(i int) (int, error) {
	return i, errors.New("Custom Error")
}

type MyError struct {
	code int
	msg  string
}

func (myError MyError) Error() string {
	return fmt.Sprintf("[%d]%s", myError.code, myError.msg)
}

func main() {
	fmt.Println("Hello World!")
	i, err := errorFunc(1)
	fmt.Printf("%d,%+v\n", i, err)
	myError := MyError{code: 500, msg: "异常"}
	var myErr error
	myErr = myError
	fmt.Printf("%s\n", myErr)
}
