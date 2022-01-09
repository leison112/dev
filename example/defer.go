package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("%d\n", funcReturn(0))
	fmt.Printf("%d\n", deferFuncReturn(0))
	fmt.Printf("%d\n", deferFuncReturn2(0))
	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
}

func funcReturn(result int) int {
	result = 1
	func() { result++ }()
	return result
}

//返回参数 返回值改变
func deferFuncReturn(result int) int {
	result = 1
	defer func() { result++ }()
	return result
}

//返回局部变量 返回值不变
func deferFuncReturn2(result int) int {
	var i = 1
	defer func() { i++ }()
	return i
}
