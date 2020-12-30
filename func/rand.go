package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	/**
		默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。
		要产生不同的数字序列，需要给定一个不同的种子。
		如果不定义seed值，则程序每次运行默认使用种子 1，所以得到的随机数都是一样的
	 */

	// 通过当前unix 时间戳，生成不同的随机数数字序列
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	// 生成一个 [0, 100) 的随机数
	fmt.Println(r.Int63n(100))
}
