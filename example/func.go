package main

import "fmt"

func innerFn(f func(string)) {
	f("innerFn")
}
func getSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	f1 := func(s string) {
		fmt.Println(s)
	}
	innerFn(f1)

	getSeq1 := getSeq()
	fmt.Println(getSeq1())
	fmt.Println(getSeq1())

	getSeq2 := getSeq()
	fmt.Println(getSeq2())
}
