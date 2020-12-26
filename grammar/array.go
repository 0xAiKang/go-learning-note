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

	// 数组遍历（方式一）
	for i:=1; i<5; i++{
		fmt.Printf("array_init[%d] = %s \n", i , arrayInit[i])
	}

	// 数组遍历（方式二）使用 _ 可以占位
	for _, v := range arrayInit {
		fmt.Println(v)
	}

	// 声明一个数组
	var arr [3]int
	fmt.Println(arr)			// [0 0 0]
	// 自动获取数组长度
	arr2 := [...]int{1,3,5, 7}
	fmt.Println(arr2)			// [1 3 5 7]

	// 数组索引
	// a[开始索引（包含），结束索引（不包含）]
	arr3 := [...]int{1,2,5,7,9}
	// 取前三个元素（go 语言不支持负的索引）
	fmt.Println(arr3[:3])
}
