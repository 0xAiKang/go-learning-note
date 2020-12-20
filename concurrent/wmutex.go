package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个读写锁
var rwMutex *sync.RWMutex

func say(name string)  {
	fmt.Println(name + "准备说话...")

	// 写加锁
	rwMutex.Lock()
	fmt.Println(name + "说话中")
	time.Sleep( 1 * time.Second)

	// 写解锁
	rwMutex.Unlock()
	fmt.Println(name + "说话结束...")
}

func eat(name string)  {
	fmt.Println(name + "准备吃饭")

	// 写锁定
	rwMutex.Lock()
	fmt.Println(name + "吃饭中")
	time.Sleep(1 * time.Second)

	// 写解锁
	rwMutex.Unlock()
	fmt.Println(name + "吃饭结束......")
}

func main()  {
	rwMutex = new(sync.RWMutex)

	go eat("小明")
	go say("小红")
	go eat("小红")
	time.Sleep(4 * time.Second)
}