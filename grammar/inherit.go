package main

import "fmt"

// go 语言中没有继承
// 在 struct 中可以包含其他的struct，继承内部 struct 的方法和变量，同时可以重写
type Person2 struct {
	inner
	Name string
	Age int
	Gender string
}

type inner struct {
	Hobby string
}

func (i * inner) say() {
	fmt.Println("Hi")
}

func main()  {
	boo := new(Person2)
	fmt.Println("暂时还没有名字：" + boo.Name)
	boo.Name = "boo"
	fmt.Println("我的名字是：" + boo.Name)
	boo.say()			//继承调用
	boo.inner.say()		//继承调用 这里也可以重写
}