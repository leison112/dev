package main

import "fmt"

//定义指针变量
//为指针变量赋值
//访问指针变量中指向地址的值

func main() {
	i := 1
	fmt.Println(&i)
	fmt.Printf("%x\n", &i)
	//一个指针指向一个值的内存地址
	//var var_name *var_type
	var ip *int
	ip = &i
	fmt.Println(ip)
	fmt.Println(*ip)

	pptr := &ip
	fmt.Println(**pptr)

	//nil
	var ptr *int
	fmt.Println(ptr)
	fmt.Printf("%x\n", ptr)
}
