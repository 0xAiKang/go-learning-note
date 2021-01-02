package main

import "fmt"

func main()  {
	/** 在Go 语言中， 没有明确初始化的变量声明会被赋予零值：
		数值类型为 0，
		布尔类型为 false，
		字符串为 ""（空字符串）
	 */
	var i int
	var s string
	var f float64
	var b bool
	fmt.Printf("%v %v %v %v\n", i, f, b, s)
}