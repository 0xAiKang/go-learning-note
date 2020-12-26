package main

import "fmt"

//  空接口可以接口任意类型
func emptyInterface(p interface{})  {
	// 使用断言获取第二个返回值
	/*if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}

	if s,ok := p.(string); ok {
		fmt.Println("String", s)
		return
	}

	fmt.Println("Unknow type")*/

	// 使用Switch 简写
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
		// 自动break 了
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow type")
	}
}

func main()  {
	emptyInterface(10)
	emptyInterface("10")
}