package main

import "fmt"

/**
*自带break
**/
func switchFn(c string) {
	switch c {
	case "S":
		fmt.Println("S")
	case "C":
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}

/**
*fallthrough
 */
func switchFnFall(c string) {
	switch c {
	case "S":
		fmt.Println("S")
		fallthrough
	case "C":
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}

func main() {
	fmt.Println("Hello World!")

	switchFn("S")
	switchFnFall("S")
}
