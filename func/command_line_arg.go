package main

import (
	"fmt"
	"os"
)

func main()  {
	// 要实验命令行参数，最好先编译一个二进制文件
	var p = fmt.Println

	// os.Args 提供原始命令行参数的访问，
	argsWithProg := os.Args
	// 而os.Args[1:] 保存了程序的全部参数
	argsWithoutProg := os.Args[1:]
	// 使用标准的下标方式取得单个参数的值
	arg := os.Args[3]

	p(argsWithProg)
	p(argsWithoutProg)
	p(arg)
}