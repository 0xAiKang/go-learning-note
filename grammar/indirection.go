package main

import "fmt"

type Vertex struct {
	X, Y float64
}

//  定义一个带指针的方法
func (v *Vertex) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 定义一个带指针的函数
func ScaleFunc(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3,}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}