package main

import "fmt"

func main()  {
	i, j := 42, 2048

	// & 操作符会生成一个指向其操作数的指针
	p := &i		// 变量 p 指向变量 i
	fmt.Println(*p)		// 通过指针 p 读取 i
	*p = 21				// 通过指针 p 设置 i
	fmt.Println(i)

	p = &j		// 变量 p 指向变量 j
	*p = *p /2  // 通过指针对j 进行除法运算
	fmt.Println(j)
}