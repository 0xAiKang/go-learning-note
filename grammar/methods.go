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
	2. 方法就是一类带有 接收者参数 的函数。在功能上，方法和函数并没有什么区别
	3. 方法的接收者参数在它自己的参数列表内，位于 func 关键字和方法名之间
	4. 在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接受者
*/
func (v Vertex) Abs() float64  {
	// v 可以简单理解成 Vertex 这个结构体的实例
	// 所以通过这个实例可以直接访问其成员及方法
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	// 直接调用 Abs方法。Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
	fmt.Println(v.Abs())
}