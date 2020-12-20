package main

import (
	"fmt"
	"sync"
)

// 互斥锁
func main()  {
	var common int = 100
	// 声明一个互斥锁
	var mutex sync.Mutex
	for  {
		go func() {
			// 加锁
			mutex.Lock()
			if common >0 {
				common --
				fmt.Println(common)
			}
			// 解锁
			mutex.Unlock()
			// 加锁的好处是防止多个线程抢占
		}()
	}
}