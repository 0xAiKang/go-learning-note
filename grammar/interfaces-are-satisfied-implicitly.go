package main

import "fmt"

// 接口与隐式实现
type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T 实现了接口I，但我们无需显示声明
func (t T) M() {
	fmt.Println(t.S)
}

func main()  {
	var i I = T{"hello"}
	i.M()
}