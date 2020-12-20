package main

import "fmt"

// 认识函数
func main()  {
	fmt.Println(getMaxAndMin(2, 3))
}

/**
	1. go 语言中的函数是支持多返回值，变量名后面的参数用户规定返回值类型
	2. 当连续两个或多个函数的以命名形参类型相同时，除最后一个类型外，其他都可省略
 */
func getMaxAndMin(a , b int32) (int32, int32) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}