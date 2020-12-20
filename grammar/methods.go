package main

import (
	"fmt"
	"math"
)

// 认识方法
type Vertex struct {
	X, Y float64
}

/*
	1. 因为在 Go 语言中，并没有类的概念
	2. 方法就是一类特殊的 接收者 参数的函数。在功能上，方法和函数并没有什么区别
*/
func (v Vertex) Abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	// 直接调用 Abs方法。Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
	fmt.Println(v.Abs())
}