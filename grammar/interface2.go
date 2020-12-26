package main

import "fmt"

// 定义一个别名
type Code string

// 定义接口
type Programmer interface {
	WriteHelloWorld() Code
}

// 实现接口
type GoProgrammer struct {

}

// 使用鸭子类型将 WriteHelloWorld 方法绑定在 GoProgrammer 结构上
func (g GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

// 实现接口
type JavaProgrammer struct {

}

// 使用鸭子类型将 WriteHelloWorld 方法绑定在 JavaProgrammer 结构上
func (j JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

// 定义一个方法，接收实例对象为参数
func writeFirstHelloWorld(p Programmer)  {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

// 通过接口实现多态
func main()  {
	// 创建一个go 的实例
	goProg := new(GoProgrammer)
	// 等同于以下写法
	//goProg := &GoProgrammer{}
	javaProg := new(GoProgrammer)

	writeFirstHelloWorld(goProg)
	writeFirstHelloWorld(javaProg)
}