package main

import (
	"fmt"
)

// 切片是一个类似动态数组的数据类型
func main()  {
	// 声明一个为 int64 类型的切片，
	// 切片和数组的区别在于，切片在声明时，不需要指定长度，因为它是动态的
	var sliceEmpty []int64

	fmt.Println(sliceEmpty) // []

	// 声明一个为 int64 类型长度为10 的数组
	var arrayEmpty [10]int64
	fmt.Println(arrayEmpty) // [0 0 0 0 0 0 0 0 0 0]

	// 初始化一个 int 64 类型的切片
	slice := make ([]int64, 10)
	/**[]type, len, cap
		其中len 个元素会被初始化为默认零值，未初始化的元素不能访问
	 */
	fmt.Println(slice)			// [0 0 0 0 0 0 0 0 0 0]

	// 初始化一个数组
	array := [10] int64 {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)			// [0 1 2 3 4 5 6 7 8 9]

	// append : 实现对 slice 元素的添加
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)			// [0 0 0 0 0 0 0 0 0 0 1 2]

	// 截取：可以通过设置下限及上限来设置截取切片
	// 截取索引 1 - 4
	fmt.Println(slice[1:4])
	// 默认下限为 0
	fmt.Println(slice[:3])
	// 默认上限为 len(s)
	fmt.Println(slice[4:])

	// 获取切片长度
	fmt.Println(len(slice))

	// 遍历
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	/*
	声明一个切片的四种方式：
	var s0 []int
	s1 := []int{}
	s2 := []int{1,3,4}
	s4 := make([]int, 3, 3)
	 */
	s := []int{}
	for i:=10; i>0; i-- {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}

	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dev"}
	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2))
}