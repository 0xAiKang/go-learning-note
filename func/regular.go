package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "[0-9]{4}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}"
	// 解析正则表达式
	reg, _:= regexp.Compile(pattern)
	fmt.Println(reg)
	// 检测是否匹配
	ok := reg.MatchString("<p>GoCN 每日新闻 (2020-04-02)</p>")
	fmt.Println(ok)
	// 返回匹配内容
	str := reg.FindString("<p>GoCN 每日新闻 (2020-04-02)</p>")
	fmt.Println(str)	// 2020-04-02

	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 通过Compile 得到一个优化过的 regexp 结构体
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 查找匹配的字符串
	fmt.Println(r.FindString("peach punch"))
}