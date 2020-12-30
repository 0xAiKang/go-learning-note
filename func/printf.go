package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main()  {
	var f = fmt.Printf
	p := point{10, 24}

	// %v 提供打印“动词”。 例如，这样打印 point 结构体的实例。
	f("%v \n", p)		// {10 24}

	// %+v 输出结构体字段名
	f("%+v \n", p)		// {x:10 y:24}

	// %T 打印值的类型
	f("%T \n", p)		// main.point

	// %d 打印标准的十进制格式化
	f("%d \n", p.x)		// 10

	// %x 提供十六进制编码
	f("%x \n", p.y)		// 10

	// %f 进行基本的浮点数输出
	f("%f \n", 10.24)	// 10.240000

	// %s 进行基本的字符串输出
	f("%s \n", "hello go")	// hello go
}