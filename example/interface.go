package main

import "fmt"

type Phone interface {
	call()
}
type IPhone struct {
}

func (phone IPhone) call() {
	fmt.Println("IPhone")
}

func main() {
	fmt.Println("Hello World!")
	var phone Phone
	phone = new(IPhone)
	phone.call()
}
