package main

import (
	"fmt"
	"time"
)

// select 语句
func main()  {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for  {
		// select 语句使得一个 goroutine 在多个通讯操作上等待。
		select {
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("BOOM!")
				return
			default:
				fmt.Println("   .")
				time.Sleep(50 * time.Millisecond)
		}
	}
}