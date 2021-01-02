package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

// 定义一个函数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 方法只是一个带接收者的函数
func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 方法即函数
func main()  {
	v := Vertex{3, 4}

	// 通过函数调用
	fmt.Println(Abs(v))

	// 通过方法调用
	fmt.Println(v.Abs2())
}