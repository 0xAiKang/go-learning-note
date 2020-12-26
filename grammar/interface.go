package main

import "fmt"

// 接口定义
type Programmer interface {
	writeHelloWorld() string
}

// 接口实例
type GoProgrammer struct {

}

// 这里使用鸭子类型将 writeHelloWorld 方法绑定在 GoProgrammer 结构上，并没有完全依赖与接口
func (g *GoProgrammer) writeHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func main()  {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.writeHelloWorld())
}