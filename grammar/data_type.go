package main

import "fmt"

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