package main

import (
	"fmt"
	"runtime"
)

func main()  {
	/**
		go 语言的Switch 与 其他编程语言的差异：
		1. 条件表达式不限制常量或者整数
		2. 单个case 中，可以出现多个结果选项，用逗号分隔
		3. 不需要手动 break
		4. 可以不设定switch 后的条件表达式，此时，整个结构与if...else的逻辑等同
	 */
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// 在go 语言中，不需要手动break
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s", os)
	}

	for i:=0; i<5; i++ {
		switch i {
		case 0,2:
			fmt.Println("Even")
		case 1,3:
			fmt.Println("Odd")
		default:
			fmt.Println(i)
		}
	}
}