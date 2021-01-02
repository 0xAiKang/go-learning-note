package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 结构体指针
func main()  {
	// 结构体字段可以通过结构体指针来方法
	v := Vertex{10, 24}
	p := &v
	fmt.Println(p)
	// 通过结构体指针 p，访问 X
	fmt.Println((*p).X)
	// 使用隐式间接引用，可以简写为：
	fmt.Println(p.X)
	fmt.Println(v)
}