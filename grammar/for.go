package main

import "fmt"

// for 语句
func main()  {
	// 传统for 循环
	for i := 1; i< 10; i++ {
		fmt.Println(i)
	}

	array := [5]string{"go", "c++", "php", "js", "ruby",}
	// for range 遍历（这个默认会遍历整个数组，和foreach 的使用效果很像）
	for index, value := range array {
		// 如果不需要下标，可以使用 _ 代替 index
		fmt.Printf("array[%d] = %s \n ", index, value)
	}

	// while 循环
	/*for true {
		fmt.Println("hello, go ")
	}*/

	// 死循环
	/*for  {
		// todo
	}*/
}