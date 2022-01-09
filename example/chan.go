package main

import "fmt"
import "runtime"
import "time"

func send(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("v-%d", i)
	}
	close(c)
}
func receive(c <-chan string) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Hello, World!")
	var c1 chan string

	c1 = make(chan string)
	//defer close(c1)
	go func() {
		c1 <- "v1"
	}()
	v1, ok1 := <-c1
	fmt.Println(v1)
	fmt.Println(ok1)
	go send(c1)
	go receive(c1)
	runtime.Gosched()
	time.Sleep(time.Second)
}
