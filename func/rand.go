package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	/**
		默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。
		要产生不同的数字序列，需要给定一个不同的种子。
		如果不定义seed值，则程序每次运行默认使用种子 1，所以得到的随机数都是一样的
	 */
	// 随机生成种子
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(s2)
	fmt.Println(r2)
	// 生成一个 [0, 100) 的随机数
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Int63n(100))
	fmt.Println(int64(r2.Int()))
}
