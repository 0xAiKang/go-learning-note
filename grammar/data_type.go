package main

import "fmt"

// 数据类型
/**
	常见的数据类型有：
	1. 指针类型（Pointer）
	2. 数组类型
	3. 结构化类型(struct)
	4. Channel 类型
	5. 函数类型
	6. 切片类型
	7. 接口类型（interface）
	8. Map 类型
 */
func main()  {
	// 先声明，后赋值
	var name string
	name = "boo"

	// 申明且赋值
	var age = 21

	// 使用 := 申明并赋值
	// 这里省略 var 关键字 ，效果和上面是一样的
	gender := "male"

	fmt.Println("我的姓名是：", name, "年龄是：", age, "性别是：", gender)
}