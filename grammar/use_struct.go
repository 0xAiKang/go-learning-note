package main

import "fmt"

// 一个结构体（struct）就是一组字段（field）
type Vertex struct {
	name string
	age int
}

func main()  {
	fmt.Println(Vertex{"boo", 21}) 		// {boo 21}

	// 结构体字段使用 . 来访问
	v := Vertex{"boo", 21}
	fmt.Println(v.name)

	// 结构体字段可以通过结构体指针来访问
	p := &v
	p.age = 22
	fmt.Println(v)		// {boo 22}

	// 结构体文法
	var (
		v1 = Vertex{"mac", 25}
		v2 = Vertex{age:18}						// name: "" 被隐式地赋予
		v3 = Vertex{}
		p2  = &Vertex{"foo", 2}
	)

	fmt.Println(v1, p2, v2, v3)
}