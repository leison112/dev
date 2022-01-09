package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j && j == 5 {
				goto breakHere
			} else {
				fmt.Printf("%d-%d\n", i, j)
			}
		}
	}
	return
breakHere:
	fmt.Println("Break Here")
}
