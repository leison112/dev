package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	for i := 0; i < 10; i++ {
		go func(_i int) {
			fmt.Printf("%d\n", _i)
		}(i)
	}
	runtime.Gosched()
	fmt.Println("Bye Bye!")
	time.Sleep(time.Second)
}
