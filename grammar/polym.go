package main

import "fmt"

// 在 go 语言中，并没有 extends，那么它是如何实现多态的呢？
type Person interface {
	sayHi()
}

type Girl struct {
	Name string
}

type Boy struct {
	Name string
}

func (this * Girl)  sayHi(){
	fmt.Println("Hi, I am  " + this.Name)
}

func (this * Boy) sayHi(){
	fmt.Println("Hi, I am  " + this.Name)
}

func main()  {
	girl := &Girl{"Yui"}
	boy := &Boy{"Boo"}

	m := map[int]Person{}
	m[0] = girl
	m[1] = boy

	for _, v := range m {
		v.sayHi()
	}

}
