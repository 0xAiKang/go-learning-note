package main

import (
	"fmt"
)

//定义除法运算函数
func Devide(num1, num2 int) int {
	if num2 == 0 {
		panic("num cannot be 0")
	} else {
		return num1 / num2
	}
}
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic的内容%v\n", r)
		}
	}()

	rs := Devide(a, b)
	fmt.Println("结果是：", rs)
}