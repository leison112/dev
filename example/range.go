package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for i, v := range nums {
		fmt.Printf("%d->%d\n", i, v)
	}
	maps := map[string]string{"a": "A", "b": "B", "c": "C"}
	for k, v := range maps {
		fmt.Printf("%s->%s\n", k, v)
	}
}
