package main

import "fmt"

func main()  {
	// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

	// 第一种定义方式
	type Person struct {
		name string  	// 姓名
		age uint32   	// 年龄
		birth string  	// 出生日期
		height float32  // 身高
		weight float32  // 体重
	}

	boo := new(Person)
	fmt.Println(boo)

	// 第二种定义方式
	mac := &Person{name: "mac", age: 18,}
	fmt.Println(mac)

	// 访问结构体
	fmt.Printf("姓名: %s, 年龄: %d", mac.name, mac.age)
}