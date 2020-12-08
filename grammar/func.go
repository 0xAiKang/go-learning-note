package main

import "fmt"

func main()  {
	fmt.Println(getMaxAndMin(2, 3))
}

// go 语言中的函数是支持多返回值
func getMaxAndMin(a int32, b int32) (int32, int32) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}