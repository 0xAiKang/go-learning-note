package main

import "fmt"

// 带缓冲的信道
func main()  {
	// 创建一个带缓冲的信道
	ch := make(chan int, 3)

	// 给信道赋值
	ch <- 1
	ch <- 2
	ch <- 3

	// 从信道中获取数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}