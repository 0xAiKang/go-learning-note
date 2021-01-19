package main

import (
	"fmt"
	"math"
)

// 方法与指针重定向
type Vertex struct {
	X, Y float64
}

// 定义两个函数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 参数类型为 Vertex
// 函数参数是一个指针，表示会对原始的值进行修改，相应地传参时，需要传入一个内存地址
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	v := Vertex{3, 4}
	// Scale(v, 10)
	// 因为上面的函数参数使用了指针，所以下面也必须接收一个指针
	Scale(&v, 10)
	fmt.Println(Abs(v))
}