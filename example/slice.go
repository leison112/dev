package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("arr len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)
}
func printSlice2(arrPt *[]int) {
	for i, v := range *arrPt {
		fmt.Printf("%d,%v\n",i,v)
	}
}
func main() {
	fmt.Println("Hello World!")
	arr := []int{1, 2, 3}
	fmt.Printf("%T\n",arr)
	printSlice(arr)
	printSlice(arr[1:2])
	printSlice(arr[:2])
	printSlice(arr)
	printSlice(append(arr, 1))

	var arr1 = make([]int, len(arr), cap(arr))
	fmt.Printf("%T\n",arr1)
	copy(arr1, arr)
	printSlice(arr1)
	printSlice2(&arr1)

}
