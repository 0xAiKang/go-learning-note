package main

import "fmt"

func main()  {
	// map 是一种无序的键值对的集合
	// 声明一个map
	var data_map map[string]string
	// 上面这段代码没有初始化map，直接赋值会创建一个 nil map，而nil map 不能用来存放键值对
	// data_map["hello"] = "go"
	fmt.Println(data_map)

	// 初始化
	data_map_init := make(map[string]string)
	// 对于 map 类型，一定要进行初始化再赋值
	data_map_init["hello"] = "world"
	fmt.Println(data_map_init)				//	map[hello:world]

	// 修改元素
	data_map_init["hello"] = "go"
	fmt.Println(data_map_init)				//	map[hello:go]

	// 获取元素
	value := data_map_init["hello"]	    	// go
	fmt.Println(value)

	// 检测元素是否存在
	elem, ok := data_map_init["hello"]
	// 如果 key 在 m 中，ok 为 true 。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。
	fmt.Printf("elem => %s, ok => %s \n", elem, ok)

	// 删除元素
	delete(data_map_init, "hello")
	fmt.Println(data_map_init)				// map[]

	// 遍历
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}
}
