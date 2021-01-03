package main

import (
	"fmt"
	"time"
)

func main()  {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.ANSIC))

	// 格式化日期的两种方式：
	fmt.Printf("%d-%d-%d %d:%d:%d \n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// 2021-1-3 10:48:49

	// 按照2006/01/02 15:04:05 进行格式化，这个日期是go 语言规定的
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 2021-01-03 10:48:49
}