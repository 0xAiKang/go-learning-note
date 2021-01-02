package main

import "fmt"

func main()  {
	// 表达式  T(v) 将值  v  转换为 类型 T
	i := 1024
	s := string(i)
	fmt.Println(i, s)
	fmt.Printf("%T %T", i, s)
}