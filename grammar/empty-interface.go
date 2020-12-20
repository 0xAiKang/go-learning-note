package main

import "fmt"

// 空接口
func main()  {
	//  定义一个空接口
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// 实现接口
func describe(i interface{})  {
	fmt.Printf("(%v, %T)\n", i, i)
}