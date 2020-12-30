package main

import (
	"fmt"
	"time"
)

func main()  {
	p := fmt.Println

	now := time.Now()
	// 从获取当前时间开始
	p(now)
	y, m, d := now.Date()
	date := fmt.Sprintf("%d-%d-%d", y, m, d)
	fmt.Println(date)

	// 通过提供年月日等信息，构建一个time
	then := time.Date(2020, 12, 30, 16, 35, 00, 122234234, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 获取当前的周期
	p(then.Weekday())
}