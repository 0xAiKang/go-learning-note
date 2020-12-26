package main

import "fmt"

// map 是一种无序的键值对的集合
func main()  {
	// 声明一个map
	var dataMap map[string]string
	// 上面这段代码没有初始化map，直接赋值会创建一个 nil map，而nil map 不能用来存放键值对
	// data_map["hello"] = "go"
	fmt.Println(dataMap)

	// 初始化
	dataMapInit := make(map[string]string)
	// 对于 map 类型，一定要进行初始化再赋值
	dataMapInit["hello"] = "world"
	fmt.Println(dataMapInit) //	map[hello:world]

	// 修改元素
	dataMapInit["hello"] = "go"
	fmt.Println(dataMapInit) //	map[hello:go]

	// 获取元素
	value := dataMapInit["hello"] // go
	fmt.Println(value)

	// 检测元素是否存在
	elem, ok := dataMapInit["hello"]
	// 如果 key 在 m 中，ok 为 true 。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。
	fmt.Printf("elem => %s, ok => %s \n", elem, ok)

	// 删除元素
	delete(dataMapInit, "hello")
	fmt.Println(dataMapInit) // map[]

	// 遍历
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	// 前面是键的类型，后面是值的类型
	m := map[string]int{"ont": 1, "two": 2, "three": 3}
	fmt.Printf("one: %d two: %d", m["one"], m["two"])

	// 在map 中，不存在的，都是零
	m2 := map[int]int{}
	fmt.Println(m2)
	m2[0] = 0
	fmt.Println(m2[0])
	// ok = false
	if v, ok:=m2[3]; ok {
		fmt.Printf("key 3's value is %d", v)
	}else {
		fmt.Println(ok)
		fmt.Println("key 3 is not existing")
	}
}
