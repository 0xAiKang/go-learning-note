package main

import (
	"fmt"
)

type Employee struct {
	Id string
	Name string
	Age int
}

// 给这个实例创建一个行为（方法）
// 和传统方法的区别在于方法名前面多了一个实例类型
func (e Employee) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", e.Name, e.Age)
}

// 为了避免内存拷贝，通过情况下会使用第二种定义方式
func (e *Employee) String2() string {
	return fmt.Sprintf("Name: %s, Age: %d", e.Name, e.Age)
}

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
func main()  {
	e1 := Employee{"1", "boo", 21}
	e2 := Employee{Name: "mac", Age: 22}
	e3 := new(Employee)

	e3.Id = "2"
	e3.Name = "lee"
	e3.Age = 18

	fmt.Println(e2)
	fmt.Printf("e is %T\n", e1)
	// 这里返回的是指针的引用，相当于 &Employee{}
	fmt.Printf("e is %T\n", e3)

	e4 := Employee{"4","boo", 21}
	fmt.Println(e4.String())

}