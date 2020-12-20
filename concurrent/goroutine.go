package main

import (
	"fmt"
	"time"
)

// go 协程
func sayHi(s string) {
	for i := 0; i<5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, i)
	}
}

func main()  {
	// 创建一个  go 协程
	go sayHi("world")		// 因为使用协程，每次调用结果都不一样
	sayHi("hello ～")		// 传统方式调用，顺序执行
}