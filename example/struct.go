package main

import "fmt"

type Human struct {
	name string
}

func (man Human) getName() string {
	return man.name
}

func main() {
	var people = Human{name: "leison"}
	fmt.Printf("Hello, 类型:%T,字符串:%s\n", people, people.name)
	fmt.Printf("指针:%p\n", people)
	fmt.Printf("函数:%s\n", people.getName())
	var man = new(Human)
	fmt.Printf("%T-%s", man, man)
}
