package main

import "fmt"
import "strings"

func main() {
	s1 := "Leison"
	s2 := "-111111"
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	fmt.Printf("Hello%s\n", build.String())
}
