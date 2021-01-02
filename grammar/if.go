package main

import "fmt"

// if 语句
func main()  {
	var num = 10

	if num > 1 {
		// 如果条件为 true 则执行以下语句
		fmt.Println(num)
		num--
	}

	if i:=1; i<10 {
		// 如果条件为 true 则执行以下语句
		fmt.Println(i)
	}

	// 同for 一样，if 语句可以在语句表达式前执行一个简单的语句
	// 不过需要注意的是：改语句声明的变量作用域仅在if 内
	if v := 1; v>0 {
		fmt.Println(v)
	}
	// fmt.Println(v)
}