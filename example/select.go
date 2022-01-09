package main

import "fmt"

func send(c chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("Quit")
			return
		default:
			fmt.Print(".")
		}
	}
}
func receive(c <-chan int, quit chan<- int) {
	//死循环
	/*
		for v:=range c{
			fmt.Println(v)
		}
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go receive(c, quit)
	send(c, quit)
}
