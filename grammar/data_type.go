package main

import (
	"fmt"
)

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

//  同导入语句一样，变量声明也可以“分组”成一个语法块。
// 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型
var (
	ToBe bool = false
	MacInt int	= 100
)

func main()  {
	// 先声明，后赋值
	var name string
	name = "boo"

	// 申明且赋值
	var age = 21

	// 使用 := 申明并赋值
	// 这里省略 var 关键字 ，效果和上面是一样的
	// 需要注意的是：函数外的每个语句都必须以关键字开始（var、func）等，因此 := 结构不能在函数外使用
	gender := "male"

	fmt.Println("我的姓名是：", name, "年龄是：", age, "性别是：", gender)

	a := 1
	// 通过取指符获得变量a 的指针
	aPtr := &a
	// 指针不能进行算术运算
	// aPtr = aPtr + 1
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)

	var s string
	fmt.Println("*" + s + "*")
	fmt.Println(len(s))
	if s == "" {
		fmt.Println("s 是一个空")
	}
}