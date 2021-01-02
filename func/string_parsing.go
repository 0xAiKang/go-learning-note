package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// Go 语言内建的 strconv 包提供了数字解析能力
	var p = fmt.Println

	// 使用ParseFloat 将string 类型转换为float 类型，这里的64 表示解析的数的位数
	f, _ := strconv.ParseFloat("10.24", 64)
	p(f)

	// 使用ParseInt 将string 类型转换为int 类型，0 表示自动推断字符串所表示的数字进制，64表示所返回的整数是以64 位存储的
	i, _:= strconv.ParseInt("1024", 0, 64)
	p(i)

	// ParseInt 会自动识别出字符串是十六进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	// Atoi 是一个基础的 10 进制整型数转换函数
	k, _ := strconv.Atoi("135")
	p(k)
}