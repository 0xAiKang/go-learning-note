package main

import (
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func otherTask()  {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}

func asycService() chan string{
	// 创建一个buffer channel
	resCh := make(chan string, 1)

	go func() {
		res := service()
		fmt.Println("returned result")
		resCh <- res
		fmt.Println("serviced exited")
	}()
	return resCh
}

func main()  {
	// 传统串行的方式：主程序 + 一些任务
	fmt.Println(service())
	otherTask()

	fmt.Println("------------------")
	// channel
	resCh := asycService()
	otherTask()
	fmt.Println(<-resCh)

	fmt.Println("------------------")
	select {
	case res:= <-asycService():
		fmt.Println(res)
		// 当耗时大于超时设定的时间时，就会提示超时
	case <-time.After(time.Millisecond * 100):
		fmt.Println("time out")
	}
}