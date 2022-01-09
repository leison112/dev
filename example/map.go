package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var vMap map[string]string
	vMap = make(map[string]string)
	vMap["name"] = "Leison"
	vMap["home"] = "Henan"
	fmt.Printf("My name is %s,My Home is %s\n", vMap["name"], vMap["home"])
	for key := range vMap {
		fmt.Println(key, ":", vMap[key])
	}
	{
		vMap["age"] = "30"
		value, ok := vMap["age"]
		fmt.Println(value)
		fmt.Println(ok)
	}
	{
		delete(vMap, "age")
		value, ok := vMap["age"]
		fmt.Println(value)
		fmt.Println(ok)
	}
}
