package main

import "fmt"

const NAME string = "Leison"
const (
	a = iota
	b
	c = "c"
	d //c
	e = iota
	f
)

func getName() string {
	return NAME
}
func main() {
	fmt.Println("Hello World!")
	fmt.Println(getName())
	fmt.Printf("%d,%d,%s,%s,%d,%d\n", a, b, c, d, e, f)
	fmt.Printf("%d,%d\n", 3<<1, 3<<2)
}
