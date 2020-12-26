package main

import "fmt"

// 定义一个结构
type Pet struct {

}

// 给结构定义一个行为/方法
func (p Pet) speak() {
	fmt.Println("...")
}

// 给结构定义一个行为/方法
func (p Pet) speakTo(host string) {
	p.speak()
	fmt.Println(host)
}

// 通过复合去扩展
type Dog struct {
	Pet
	//p *Pet
}

/*func (d Dog) speak() {
	d.p.speak()
}

func (d Dog) speakTo(host string) {
	d.p.speakTo(host)
}*/

func main()  {
	dog := new(Dog)
 	dog.speakTo("Wang")
}
