package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 通过Compile 得到一个优化过的 regexp 结构体
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 查找匹配的字符串
	fmt.Println(r.FindString("peach punch"))
}