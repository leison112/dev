package main

import "fmt"
import "container/list"

func main()  {
    l := list.New()
    l.PushBack("second")
    l.PushFront("first")
    for i := l.Front(); i != nil; i=i.Next() {
    	fmt.Println(i.Value)
    }
}