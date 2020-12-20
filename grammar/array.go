package main

import "fmt"

// 数组
func main()  {
	// 定义一个长度为 10  string  类型的数组（在go 中声明数组是需要指定长度）
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组。
	var array [10]string
	fmt.Println(array)

	// 数组赋值
	array[0] = "php"
	array[1] = "c++"
	array[3] = "js"
	array[5] = "go"
	fmt.Println(array)

	// 声明数组时初始化
	var arrayInit = [5]string{"go", "c++", "php", "js", "ruby",}
	fmt.Println(arrayInit)

	// 数组遍历
	for i:=1; i<5; i++{
		fmt.Printf("array_init[%d] = %s \n", i , arrayInit[i])
	}
}
