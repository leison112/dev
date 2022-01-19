// fmt project main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Fprintln(os.Stdout, "os.Stdout")
	f, e := os.OpenFile("./demo.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Println("打开文件出错", e)
		return
	}
	fmt.Println(f.Name())
	fmt.Println(filepath.Abs("./"))
	var (
		name string
		age  int
	)
	//中间用空格分隔
	// fmt.Scan(&name, &age)
	// fmt.Fprintln(f, name, age)
	//为什么不能多次Scan
	// fmt.Scanf("name=%s age=%d", &name, &age)
	// fmt.Fprintln(f, name, age)

	b := strings.NewReader("Leison 30")
	fmt.Fscan(b, &name, &age)
	fmt.Fprintln(f, name, age, "Fscan")

	e2 := fmt.Errorf("My Error")
	fmt.Fprintln(f, e2)
	fmt.Println("Bye Bye")
}
