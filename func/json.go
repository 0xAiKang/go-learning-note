package main

import (
	"encoding/json"
	"fmt"
)

// 通过结构体演示自定义编码和解码
type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main()  {
	var p = fmt.Println
	var f = fmt.Printf
	// Go 提供内建的 JSON 编码解码（序列化反序列化）支持
	bolB, _ := json.Marshal(true)
	f("%T \n", bolB)
	p(string(bolB))

	intB, _ := json.Marshal(1)
	f("%T \n", intB)
	p(string(intB))

	fltB, _ := json.Marshal(10.24)
	p(string(fltB))

	strB, _ := json.Marshal("hello go")
	p(string(strB))

	p("----------")

	// 将切片编码成 JSON 数组
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	f("%T \n", slcD)
	f("%T \n", slcB)
	f("%T \n", string(slcB))
	p(string(slcB))

	p("----------")

	// 将Map 编码成 JSON 对象
	mapD := map[string]int{"apple": 5, "peach": 10, "pear": 7}
	mapB, _ := json.Marshal(mapD)
	f("%T \n", mapD)
	f("%T \n", mapB)
	f("%T \n", string(mapB))
	p(string(mapB))

	p("----------")

	// 自定义编码类型
	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	p(string(res1B))

	p("----------")

	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "pearch", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	p(string(res2B))

	p("----------")

	// byt := []byte("{\"num\": 10.24, \"str\": [\"a\", \"b\"]}")
	byt := []byte(`{"num": 10.24, "str": ["a", "b"]}`)

	// 注意这里不是一个表达式
	// 声明一个 键为string，值为任意类型的值（空接口）的map
	var data map[string]interface{}

	// 解码及错误检查，前面是需要解码的数据，后面是存储解码后的内容
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	p(data)
	num := data["num"].(float64)
	p(num)

	// 使用类型断言判断接口值是否保存了一个特定的类型
	// 接口断言返回两个值：其底层值以及一个报告断言是否成功的布尔值
	strs, ok:= data["str"].([]interface{})
	str1 := strs[0]
	p(str1, ok)

	// 定义一个包含数组的JSON 对象
	str := `{"page": 1, "fruits": ["apple", "pear"]}`
	res := response2{}
	// 解码
	json.Unmarshal([]byte(str), &res)
	p(res)
	p(res.Fruits[0])
}