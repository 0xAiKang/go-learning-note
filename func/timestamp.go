package main

import (
	"fmt"
	"time"
)

func main()  {
	p := fmt.Println

	// 获取当前时间
	now := time.Now()
	// 获取当前时间对应对应时间戳
	secs := now.Unix()
	// 获取当前时间对应纳秒时间戳
	nanos := now.UnixNano()
	// 需要注意的是，毫秒时间戳是不存在的，如果要得到毫秒时间戳，需要手动转一下
	millis := nanos / 1000000
	p(secs)
	p(nanos)
	p(millis)

	// 将时间戳转换为对应的时间
	p(time.Unix(secs, 0))
	// 将纳秒时间戳转换为对应的时间
	p(time.Unix(0, nanos))
}