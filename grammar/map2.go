package main

import "fmt"

func main()  {
	// 声明一个map， value 是一个函数
	m := map[int]func(op int)int{}
	
	m[1] = func(op int) int {
		return op
	}
	
	m[2] = func(op int) int {
		return op*op
	}

	m[3] = func(op int) int {
		return op*op*op
	}

	// 通过key 来访问所对应不同的函数
	fmt.Println(m[1](2), m[2](2), m[3](2))

	/**go 语言内置集合没有Set 实现，可以 map[type]bool
	1. 元素的唯一性
	2. 基本操作
		1. 添加元素
		2. 判断元素是否存在
		3. 删除元素
		4. 获取元素个数
	 */
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
	// 添加/修改元素
	mySet[3] = false
	// 获取元素个数
	fmt.Println(len(mySet))
	// 删除指定的 key
	delete(mySet, 1)
}