package main

import (
	"fmt"
	"math"
)

// 指针与函数。可以为指针接受者声明方法
type Vertex struct {
	X, Y float64
}

// float64 的作用是用来定义返回值的类型
func (v Vertex) Abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/*
	1. float64 的作用是定义函数参数的类型
	2. 通过 *Vertex 定义了 Scale 方法，那么 Scale 方法会对原始 Vertex 值的副本进行操作
 */
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}