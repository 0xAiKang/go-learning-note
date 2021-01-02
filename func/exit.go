package main

import (
	"fmt"
	"os"
)

func run()  {
	fmt.Println("我被调用了")
}

func main()  {
	/**
		在Go 语言中， 退出程序有两种方式
		os.Exit vs. panic
		1. os.Exit 退出时 不会调用 defer  函数
		2. panic 退出前会执行 defer 函数
		3. os.Exit 退出时不会输出当前调用栈的信息
		4. panic 用于不可恢复的错误
	 */
	fmt.Println("start")
	defer run()
	// panic("something went wrong")
	os.Exit(3)
}