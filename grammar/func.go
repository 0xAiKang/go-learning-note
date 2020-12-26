package main

import (
	"fmt"
)

// 认识函数
func main()  {
	fmt.Println(getMaxAndMin(2, 3))

	fmt.Println(sum(10, 89, 1))

	testDefer()
}

/**
	1. go 语言中的函数是支持多返回值，变量名后面的参数用户规定返回值类型
	2. 当连续两个或多个函数的以命名形参类型相同时，除最后一个类型外，其他都可省略
	3. 所有参数都是值传递：slice、map、channel 会有引用传递的错觉
	4. 函数可以作为变量的值
	5. 函数可以作为参数和返回值
 */
func getMaxAndMin(a , b int32) (int32, int32) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

// 可变长参数
func sum(ops ...int) int {
	 res := 0
	for _, op := range ops{
		res += op
	}
	return res
}

// 延迟执行函数 defer
func testDefer()  {
	// defer
	/*defer func() {
		fmt.Println("Clear resources")
	}()*/

	// recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	panic("Something wrong") // defer 仍会执行
}