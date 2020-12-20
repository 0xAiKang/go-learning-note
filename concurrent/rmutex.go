package main

import (
	"fmt"
	"sync"
	"time"
)

//
var rwMutex *sync.RWMutex

func say(name string)  {
	fmt.Println(name + "正准备说话")
	// 读锁加锁
	rwMutex.RLock()
	fmt.Println(name + "说话中...")
	time.Sleep(1 * time.Second)
	// 读锁解锁
	rwMutex.RUnlock()
	fmt.Println(name + "说话结束")
}

//  读写锁
func main()  {
	/*
	    加锁的目的是防止多个线程同时进行共享内存的修改，go  的读写锁遵循两个基本原则：
		1. 读的时候资源不是互斥的，可以多个协程一起读，也可以加多个读锁
		2. 写的时候只能有一个协程进行写，写锁只能有一个，其他协程既不能读也不能写

		func (rw *RWMutex) Lock() 写锁

		func (rw *RWMutex) Unlock() 写锁解锁

		func (rw *RWMutex) RLock() 读锁

		func (rw *RWMutex)RUnlock()　读锁解锁
	*/

	rwMutex = new(sync.RWMutex)
	/*传统执行方式
	say("boo")
	say("mac")*/

	// 协程执行方式
	go say("boo")
	go say("mac")
	time.Sleep(2 * time.Second)
}