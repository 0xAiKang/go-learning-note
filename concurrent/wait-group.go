package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	fmt.Println(count())

	fmt.Println(countThreadSave())

	fmt.Println(countWaitGroup())
}

func count() int {
	counter := 0
	//  创建一个协程进行加数
	for i:=0; i< 5000; i++ {
		go func() {
			counter ++
		}()
	}
	// 可以发现每一次累加的结果都不是 5000
	return counter
}

func countThreadSave() int {
	var mut sync.Mutex
	counter := 0
	for i:=0; i< 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	// 这种方式并不好的原因是：并不知道需要等待的时间是多少
	time.Sleep(1 * time.Second)
	return counter
}

func countWaitGroup() int {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i:=0; i< 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}